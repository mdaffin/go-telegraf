package telegraf

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMeasurementsToString(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name   string
		input  Measurement
		output string
	}{
		{"Simple", Measurement{"simple", nil, nil, now}, fmt.Sprintf("simple %d", now.UnixNano())},

		{"WithStringField", Measurement{"measurement", nil, map[string]interface{}{"key": "value"}, now}, fmt.Sprintf("measurement key=\"value\" %d", now.UnixNano())},
		{"WithIntField", Measurement{"measurement", nil, map[string]interface{}{"key": 42}, now}, fmt.Sprintf("measurement key=42i %d", now.UnixNano())},
		{"WithUIntField", Measurement{"measurement", nil, map[string]interface{}{"key": uint(42)}, now}, fmt.Sprintf("measurement key=42i %d", now.UnixNano())},
		{"WithUInt8Field", Measurement{"measurement", nil, map[string]interface{}{"key": uint8(41)}, now}, fmt.Sprintf("measurement key=41i %d", now.UnixNano())},
		{"WithUInt16Field", Measurement{"measurement", nil, map[string]interface{}{"key": uint16(40)}, now}, fmt.Sprintf("measurement key=40i %d", now.UnixNano())},
		{"WithUInt32Field", Measurement{"measurement", nil, map[string]interface{}{"key": uint32(39)}, now}, fmt.Sprintf("measurement key=39i %d", now.UnixNano())},
		{"WithUInt64Field", Measurement{"measurement", nil, map[string]interface{}{"key": uint64(38)}, now}, fmt.Sprintf("measurement key=38i %d", now.UnixNano())},
		{"WithIntField", Measurement{"measurement", nil, map[string]interface{}{"key": int(-42)}, now}, fmt.Sprintf("measurement key=-42i %d", now.UnixNano())},
		{"WithInt8Field", Measurement{"measurement", nil, map[string]interface{}{"key": int8(-43)}, now}, fmt.Sprintf("measurement key=-43i %d", now.UnixNano())},
		{"WithInt16Field", Measurement{"measurement", nil, map[string]interface{}{"key": int16(-44)}, now}, fmt.Sprintf("measurement key=-44i %d", now.UnixNano())},
		{"WithInt32Field", Measurement{"measurement", nil, map[string]interface{}{"key": int32(-45)}, now}, fmt.Sprintf("measurement key=-45i %d", now.UnixNano())},
		{"WithInt64Field", Measurement{"measurement", nil, map[string]interface{}{"key": int64(-46)}, now}, fmt.Sprintf("measurement key=-46i %d", now.UnixNano())},
		{"WithFloat32Field", Measurement{"measurement", nil, map[string]interface{}{"key": float32(42.5)}, now}, fmt.Sprintf("measurement key=42.5 %d", now.UnixNano())},
		{"WithFloat64Field", Measurement{"measurement", nil, map[string]interface{}{"key": float64(42.5)}, now}, fmt.Sprintf("measurement key=42.5 %d", now.UnixNano())},
		// Non determinstic - causes the test to fail some of the time
		//{"WithMultipleFields", Measurement{"measurement", nil, map[string]interface{}{"key1": 5, "key2": "value", "key3": float64(23)}, now}, fmt.Sprintf("measurement key1=5i,key2=\"value\",key3=23 %d", now.UnixNano())},

		{"WithTag", Measurement{"measurement", map[string]string{"key": "value"}, nil, now}, fmt.Sprintf("measurement,key=value %d", now.UnixNano())},
		// Non determinstic - causes the test to fail some of the time
		//{"WithMultipleTags", Measurement{"measurement", map[string]string{"key1": "value1", "key2": "value2"}, nil, now}, fmt.Sprintf("measurement,key1=value1,key2=value2 %d", now.UnixNano())},
		{"WithEscapedTagName", Measurement{"measurement", map[string]string{"tag,= \"-": "value"}, nil, now}, fmt.Sprintf("measurement,tag\\,\\=\\ \"-=value %d", now.UnixNano())},
		{"WithEscapedTagValue", Measurement{"measurement", map[string]string{"key": "tag,= \"-"}, nil, now}, fmt.Sprintf("measurement,key=tag\\,\\=\\ \"- %d", now.UnixNano())},
		{"WithEscapedFieldName", Measurement{"measurement", nil, map[string]interface{}{"field,= \"-": "value"}, now}, fmt.Sprintf("measurement field\\,\\=\\ \"-=\"value\" %d", now.UnixNano())},
		{"WithEscapedFieldValue", Measurement{"measurement", nil, map[string]interface{}{"key": "field,= \"-"}, now}, fmt.Sprintf("measurement key=\"field,= \\\"-\" %d", now.UnixNano())},
		{"WithEscapedMeasurement", Measurement{"meas,ure ment", nil, nil, now}, fmt.Sprintf("meas\\,ure\\ ment %d", now.UnixNano())},
		{"WithNoTimestamp", Measurement{"measurement", nil, nil, time.Time{}}, fmt.Sprintf("measurement")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, test.input.String())
		})
	}
}

func BenchmarkMeasurementsToString(b *testing.B) {
	weather := MeasureFloat64("weather", "temp", 22.7)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		weather.String()
	}
}
