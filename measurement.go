package telegraf

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	// To escape the various parts of the influxdb line protocal
	measurementEscaper = strings.NewReplacer(`,`, `\,`, ` `, `\ `)
	keyEscaper         = strings.NewReplacer(`,`, `\,`, ` `, `\ `, `=`, `\=`)
	tagValueEscaper    = keyEscaper
)

// Measurement that can be sent to influxdb or telegraf. The measurement
// consists of three parts, the name of the metric, a set of fields and their
// values, optional tags and a timestamp.
type Measurement struct {
	name      string
	tagSet    map[string]string
	fieldSet  map[string]interface{}
	timestamp time.Time
}

// MeasureFloat64 creates a new measurement with the given float64 field.
func MeasureFloat64(name string, field string, value float64) Measurement {
	return newMeasurement(name, field, value)
}

// MeasureFloat32 creates a new measurement with the given float32 field.
func MeasureFloat32(name string, field string, value float32) Measurement {
	return newMeasurement(name, field, value)
}

// MeasureInt creates a new measurement with the given int field.
func MeasureInt(name string, field string, value int) Measurement {
	return newMeasurement(name, field, value)
}

// MeasureInt8 creates a new measurement with the given int8 field.
func MeasureInt8(name string, field string, value int8) Measurement {
	return newMeasurement(name, field, value)
}

// MeasureInt16 creates a new measurement with the given int16 field.
func MeasureInt16(name string, field string, value int16) Measurement {
	return newMeasurement(name, field, value)
}

// MeasureInt32 creates a new measurement with the given int32 field.
func MeasureInt32(name string, field string, value int32) Measurement {
	return newMeasurement(name, field, value)
}

// MeasureInt64 creates a new measurement with the given int64 field.
func MeasureInt64(name string, field string, value int64) Measurement {
	return newMeasurement(name, field, value)
}

// MeasureUInt creates a new measurement with the given uint field.
func MeasureUInt(name string, field string, value uint) Measurement {
	return newMeasurement(name, field, value)
}

// MeasureUInt8 creates a new measurement with the given uint8 field.
func MeasureUInt8(name string, field string, value uint8) Measurement {
	return newMeasurement(name, field, value)
}

// MeasureUInt16 creates a new measurement with the given uint16 field.
func MeasureUInt16(name string, field string, value uint16) Measurement {
	return newMeasurement(name, field, value)
}

// MeasureUInt32 creates a new measurement with the given uint32 field.
func MeasureUInt32(name string, field string, value uint32) Measurement {
	return newMeasurement(name, field, value)
}

// MeasureString creates a new measurement with the given string field.
func MeasureString(name string, field string, value string) Measurement {
	return newMeasurement(name, field, value)
}

// MeasureBool creates a new measurement with the given bool field.
func MeasureBool(name string, field string, value bool) Measurement {
	return newMeasurement(name, field, value)
}

// MeasureUInt64 creates a new measurement with the given uint64 field.
func MeasureUInt64(name string, field string, value uint64) Measurement {
	return newMeasurement(name, field, value)
}

func newMeasurement(name string, field string, value interface{}) Measurement {
	return Measurement{
		name:      name,
		tagSet:    map[string]string{},
		fieldSet:  map[string]interface{}{field: value},
		timestamp: time.Now(),
	}
}

func (m Measurement) SetTime(time time.Time) Measurement {
	m.timestamp = time
	return m
}

func (m Measurement) AddTag(name string, value string) Measurement {
	m.tagSet[name] = value
	return m
}

// AddNanosecondsSince t as the field called name stored as an int64.
func (m Measurement) AddNanosecondsSince(name string, t time.Time) Measurement {
	m.fieldSet[name] = time.Since(t).Nanoseconds()
	return m
}

// AddMillisecondsSince t as the field called name stored as a int64.
func (m Measurement) AddMillisecondsSince(name string, t time.Time) Measurement {
	m.fieldSet[name] = time.Since(t).Nanoseconds() / 1e6
	return m
}

// AddSecondsSince t as the field called name stored as a float64.
func (m Measurement) AddSecondsSince(name string, t time.Time) Measurement {
	m.fieldSet[name] = time.Since(t).Seconds()
	return m
}

// AddMinutesSince t as the field called name stored as a float64.
func (m Measurement) AddMinutesSince(name string, t time.Time) Measurement {
	m.fieldSet[name] = time.Since(t).Minutes()
	return m
}

// AddHoursSince t as the field called name stored as a float64.
func (m Measurement) AddHoursSince(name string, t time.Time) Measurement {
	m.fieldSet[name] = time.Since(t).Hours()
	return m
}

// AddBool field called name.
func (m Measurement) AddBool(name string, value bool) Measurement {
	m.fieldSet[name] = value
	return m
}

// AddInt field called name.
func (m Measurement) AddInt(name string, value int) Measurement {
	m.fieldSet[name] = value
	return m
}

// AddInt8 field called name.
func (m Measurement) AddInt8(name string, value int8) Measurement {
	m.fieldSet[name] = value
	return m
}

// AddInt16 field called name.
func (m Measurement) AddInt16(name string, value int16) Measurement {
	m.fieldSet[name] = value
	return m
}

// AddInt32 field called name.
func (m Measurement) AddInt32(name string, value int32) Measurement {
	m.fieldSet[name] = value
	return m
}

// AddInt64 field called name.
func (m Measurement) AddInt64(name string, value int64) Measurement {
	m.fieldSet[name] = value
	return m
}

// AddUInt field called name.
func (m Measurement) AddUInt(name string, value uint) Measurement {
	m.fieldSet[name] = value
	return m
}

// AddUInt8 field called name.
func (m Measurement) AddUInt8(name string, value uint8) Measurement {
	m.fieldSet[name] = value
	return m
}

// AddUInt16 field called name.
func (m Measurement) AddUInt16(name string, value uint16) Measurement {
	m.fieldSet[name] = value
	return m
}

// AddUInt32 field called name.
func (m Measurement) AddUInt32(name string, value uint32) Measurement {
	m.fieldSet[name] = value
	return m
}

// AddUInt64 field called name.
func (m Measurement) AddUInt64(name string, value uint64) Measurement {
	m.fieldSet[name] = value
	return m
}

// AddFloat32 field called name.
func (m Measurement) AddFloat32(name string, value float32) Measurement {
	m.fieldSet[name] = value
	return m
}

// AddFloat64 field called name.
func (m Measurement) AddFloat64(name string, value float64) Measurement {
	m.fieldSet[name] = value
	return m
}

// AddString field called name.
func (m Measurement) AddString(name string, value string) Measurement {
	m.fieldSet[name] = value
	return m
}

func (m Measurement) String() string {
	line := measurementEscaper.Replace(m.name)
	for tag, value := range m.tagSet {
		line += "," + keyEscaper.Replace(tag) + "=" + tagValueEscaper.Replace(value)
	}

	first := true
	for field, value := range m.fieldSet {
		if first {
			line += " "
			first = false
		} else {
			line += ","
		}
		line += keyEscaper.Replace(field) + "="

		switch v := value.(type) {
		case int, uint, uint8, uint16, uint32, uint64, int8, int16, int32, int64:
			line += fmt.Sprintf("%di", v)
		case float32:
			line += strconv.FormatFloat(float64(v), 'f', -1, 32)
		case float64:
			line += strconv.FormatFloat(v, 'f', -1, 64)
		case string:
			line += fmt.Sprintf("%q", v)
		case bool:
			line += fmt.Sprintf("%t", v)
		default:
			line += fmt.Sprintf("%v", v)
		}
	}

	if !m.timestamp.IsZero() {
		line += fmt.Sprintf(" %d", m.timestamp.UnixNano())
	}
	return line
}
