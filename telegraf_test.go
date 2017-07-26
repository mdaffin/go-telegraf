package telegraf

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMetricsToString(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name   string
		input  Metric
		output string
	}{
		{"Simple", Metric{"simple", nil, nil, now}, fmt.Sprintf("simple %d", now.UnixNano())},

		{"WithStringField", Metric{"measurement", nil, map[string]interface{}{"key": "value"}, now}, fmt.Sprintf("measurement key=\"value\" %d", now.UnixNano())},
		{"WithIntField", Metric{"measurement", nil, map[string]interface{}{"key": 42}, now}, fmt.Sprintf("measurement key=42i %d", now.UnixNano())},
		{"WithUIntField", Metric{"measurement", nil, map[string]interface{}{"key": uint(42)}, now}, fmt.Sprintf("measurement key=42i %d", now.UnixNano())},
		{"WithUInt8Field", Metric{"measurement", nil, map[string]interface{}{"key": uint8(41)}, now}, fmt.Sprintf("measurement key=41i %d", now.UnixNano())},
		{"WithUInt16Field", Metric{"measurement", nil, map[string]interface{}{"key": uint16(40)}, now}, fmt.Sprintf("measurement key=40i %d", now.UnixNano())},
		{"WithUInt32Field", Metric{"measurement", nil, map[string]interface{}{"key": uint32(39)}, now}, fmt.Sprintf("measurement key=39i %d", now.UnixNano())},
		{"WithUInt64Field", Metric{"measurement", nil, map[string]interface{}{"key": uint64(38)}, now}, fmt.Sprintf("measurement key=38i %d", now.UnixNano())},
		{"WithIntField", Metric{"measurement", nil, map[string]interface{}{"key": int(-42)}, now}, fmt.Sprintf("measurement key=-42i %d", now.UnixNano())},
		{"WithInt8Field", Metric{"measurement", nil, map[string]interface{}{"key": int8(-43)}, now}, fmt.Sprintf("measurement key=-43i %d", now.UnixNano())},
		{"WithInt16Field", Metric{"measurement", nil, map[string]interface{}{"key": int16(-44)}, now}, fmt.Sprintf("measurement key=-44i %d", now.UnixNano())},
		{"WithInt32Field", Metric{"measurement", nil, map[string]interface{}{"key": int32(-45)}, now}, fmt.Sprintf("measurement key=-45i %d", now.UnixNano())},
		{"WithInt64Field", Metric{"measurement", nil, map[string]interface{}{"key": int64(-46)}, now}, fmt.Sprintf("measurement key=-46i %d", now.UnixNano())},
		{"WithFloat32Field", Metric{"measurement", nil, map[string]interface{}{"key": float32(42.5)}, now}, fmt.Sprintf("measurement key=42.5 %d", now.UnixNano())},
		{"WithFloat64Field", Metric{"measurement", nil, map[string]interface{}{"key": float64(42.5)}, now}, fmt.Sprintf("measurement key=42.5 %d", now.UnixNano())},
		// Non determinstic - causes the test to fail some of the time
		//{"WithMultipleFields", Metric{"measurement", nil, map[string]interface{}{"key1": 5, "key2": "value", "key3": float64(23)}, now}, fmt.Sprintf("measurement key1=5i,key2=\"value\",key3=23 %d", now.UnixNano())},

		{"WithTag", Metric{"measurement", map[string]string{"key": "value"}, nil, now}, fmt.Sprintf("measurement,key=value %d", now.UnixNano())},
		// Non determinstic - causes the test to fail some of the time
		//{"WithMultipleTags", Metric{"measurement", map[string]string{"key1": "value1", "key2": "value2"}, nil, now}, fmt.Sprintf("measurement,key1=value1,key2=value2 %d", now.UnixNano())},
		{"WithEscapedTagName", Metric{"measurement", map[string]string{"tag,= \"-": "value"}, nil, now}, fmt.Sprintf("measurement,tag\\,\\=\\ \"-=value %d", now.UnixNano())},
		{"WithEscapedTagValue", Metric{"measurement", map[string]string{"key": "tag,= \"-"}, nil, now}, fmt.Sprintf("measurement,key=tag\\,\\=\\ \"- %d", now.UnixNano())},
		{"WithEscapedFieldName", Metric{"measurement", nil, map[string]interface{}{"field,= \"-": "value"}, now}, fmt.Sprintf("measurement field\\,\\=\\ \"-=\"value\" %d", now.UnixNano())},
		{"WithEscapedFieldValue", Metric{"measurement", nil, map[string]interface{}{"key": "field,= \"-"}, now}, fmt.Sprintf("measurement key=\"field,= \\\"-\" %d", now.UnixNano())},
		{"WithEscapedMeasurement", Metric{"meas,ure ment", nil, nil, now}, fmt.Sprintf("meas\\,ure\\ ment %d", now.UnixNano())},
		{"WithNoTimestamp", Metric{"measurement", nil, nil, time.Time{}}, fmt.Sprintf("measurement")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, test.input.String())
		})
	}
}

func BenchmarkMetricsToString(b *testing.B) {
	metric := NewMetric("weather")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		metric.String()
	}
}

func BenchmarkUDPWriteRaw(b *testing.B) {
	conn, err := net.Dial("udp", "127.0.0.1:8094")
	if err != nil {
		b.Fatal("failed to connect: %s", err)
	}
	defer conn.Close()
	text := "weather,location=us-midwest temperature=82\n"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if _, err := fmt.Fprintf(conn, text); err != nil {
			b.Fatal("failed to write: %s", err)
		}
	}
}
