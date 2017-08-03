package telegraf

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	// Escapes the various parts of the influxdb line protocal
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

// NewMeasurement creates a blank measurement without any fields, you must add
// a field before trying to send it to telegraf using on of the Add* methods.
// This is useful when you want to add tags before you have a field measurement
// avaiable.
func NewMeasurement(name string) Measurement {
	return Measurement{
		name:      name,
		tagSet:    map[string]string{},
		fieldSet:  map[string]interface{}{},
		timestamp: time.Now(),
	}
}

// Name of the measurement.
func (m Measurement) Name() string {
	return m.name
}

// SetTime of the measurement. The default is time.Now(), this can be used to
// override the default. Set it to a zero time to unset the time, which will
// cause telegraf or influxdb to set the time when they recieve the measurement
// instead.
func (m Measurement) SetTime(time time.Time) Measurement {
	m.timestamp = time
	return m
}

// AddTag to the measurement. Tags are global for all fields in this
// measurement - if you want them to have differenent tags you must create a
// second measurement with the alternate tags.
func (m Measurement) AddTag(name string, value string) Measurement {
	if value != "" {
		m.tagSet[name] = value
	}
	return m
}

// AddTags to the measurement. Tags are global for all fields in this
// measurement - if you want them to have differenent tags you must create a
// second measurement with the alternate tags.
func (m Measurement) AddTags(tags map[string]string) Measurement {
	for name, value := range tags {
		m.AddTag(name, value)
	}
	return m
}

// MeasureNanosecondsSince creates a new measurement with the given uint64 field.
func MeasureNanosecondsSince(name string, field string, t time.Time) Measurement {
	return NewMeasurement(name).AddNanosecondsSince(field, t)
}

// AddNanosecondsSince t as the field called name stored as an uint64.
func (m Measurement) AddNanosecondsSince(name string, t time.Time) Measurement {
	m.fieldSet[name] = time.Since(t).Nanoseconds()
	return m
}

// MeasureMillisecondsSince creates a new measurement with the given float64 field.
func MeasureMillisecondsSince(name string, field string, t time.Time) Measurement {
	return NewMeasurement(name).AddMillisecondsSince(field, t)
}

// AddMillisecondsSince t as the field called name stored as a float64.
func (m Measurement) AddMillisecondsSince(name string, t time.Time) Measurement {
	m.fieldSet[name] = float64(time.Since(t).Nanoseconds()) / float64((int64(time.Millisecond) / int64(time.Nanosecond)))
	return m
}

// MeasureSecondsSince creates a new measurement with the given float64 field.
func MeasureSecondsSince(name string, field string, t time.Time) Measurement {
	return NewMeasurement(name).AddSecondsSince(field, t)
}

// AddSecondsSince t as the field called name stored as a float64.
func (m Measurement) AddSecondsSince(name string, t time.Time) Measurement {
	m.fieldSet[name] = time.Since(t).Seconds()
	return m
}

// MeasureMinutesSince creates a new measurement with the given float64 field.
func MeasureMinutesSince(name string, field string, t time.Time) Measurement {
	return NewMeasurement(name).AddMinutesSince(field, t)
}

// AddMinutesSince t as the field called name stored as a float64.
func (m Measurement) AddMinutesSince(name string, t time.Time) Measurement {
	m.fieldSet[name] = time.Since(t).Minutes()
	return m
}

// MeasureHoursSince creates a new measurement with the given float64 field.
func MeasureHoursSince(name string, field string, t time.Time) Measurement {
	return NewMeasurement(name).AddHoursSince(field, t)
}

// AddHoursSince t as the field called name stored as a float64.
func (m Measurement) AddHoursSince(name string, t time.Time) Measurement {
	m.fieldSet[name] = time.Since(t).Hours()
	return m
}

// MeasureBool creates a new measurement with the given bool field.
func MeasureBool(name string, field string, value bool) Measurement {
	return NewMeasurement(name).AddBool(field, value)
}

// AddBool field called name.
func (m Measurement) AddBool(name string, value bool) Measurement {
	m.fieldSet[name] = value
	return m
}

// MeasureInt creates a new measurement with the given int field.
func MeasureInt(name string, field string, value int) Measurement {
	return NewMeasurement(name).AddInt(field, value)
}

// AddInt field called name.
func (m Measurement) AddInt(name string, value int) Measurement {
	m.fieldSet[name] = value
	return m
}

// MeasureInt8 creates a new measurement with the given int8 field.
func MeasureInt8(name string, field string, value int8) Measurement {
	return NewMeasurement(name).AddInt8(field, value)
}

// AddInt8 field called name.
func (m Measurement) AddInt8(name string, value int8) Measurement {
	m.fieldSet[name] = value
	return m
}

// MeasureInt16 creates a new measurement with the given int16 field.
func MeasureInt16(name string, field string, value int16) Measurement {
	return NewMeasurement(name).AddInt16(field, value)
}

// AddInt16 field called name.
func (m Measurement) AddInt16(name string, value int16) Measurement {
	m.fieldSet[name] = value
	return m
}

// MeasureInt32 creates a new measurement with the given int32 field.
func MeasureInt32(name string, field string, value int32) Measurement {
	return NewMeasurement(name).AddInt32(field, value)
}

// AddInt32 field called name.
func (m Measurement) AddInt32(name string, value int32) Measurement {
	m.fieldSet[name] = value
	return m
}

// MeasureInt64 creates a new measurement with the given int64 field.
func MeasureInt64(name string, field string, value int64) Measurement {
	return NewMeasurement(name).AddInt64(field, value)
}

// AddInt64 field called name.
func (m Measurement) AddInt64(name string, value int64) Measurement {
	m.fieldSet[name] = value
	return m
}

// MeasureUInt creates a new measurement with the given uint field.
func MeasureUInt(name string, field string, value uint) Measurement {
	return NewMeasurement(name).AddUInt(field, value)
}

// AddUInt field called name.
func (m Measurement) AddUInt(name string, value uint) Measurement {
	m.fieldSet[name] = value
	return m
}

// MeasureUInt8 creates a new measurement with the given uint8 field.
func MeasureUInt8(name string, field string, value uint8) Measurement {
	return NewMeasurement(name).AddUInt8(field, value)
}

// AddUInt8 field called name.
func (m Measurement) AddUInt8(name string, value uint8) Measurement {
	m.fieldSet[name] = value
	return m
}

// MeasureUInt16 creates a new measurement with the given uint16 field.
func MeasureUInt16(name string, field string, value uint16) Measurement {
	return NewMeasurement(name).AddUInt16(field, value)
}

// AddUInt16 field called name.
func (m Measurement) AddUInt16(name string, value uint16) Measurement {
	m.fieldSet[name] = value
	return m
}

// MeasureUInt32 creates a new measurement with the given uint32 field.
func MeasureUInt32(name string, field string, value uint32) Measurement {
	return NewMeasurement(name).AddUInt32(field, value)
}

// AddUInt32 field called name.
func (m Measurement) AddUInt32(name string, value uint32) Measurement {
	m.fieldSet[name] = value
	return m
}

// MeasureUInt64 creates a new measurement with the given uint64 field.
func MeasureUInt64(name string, field string, value uint64) Measurement {
	return NewMeasurement(name).AddUInt64(field, value)
}

// AddUInt64 field called name.
func (m Measurement) AddUInt64(name string, value uint64) Measurement {
	m.fieldSet[name] = value
	return m
}

// MeasureFloat32 creates a new measurement with the given float32 field.
func MeasureFloat32(name string, field string, value float32) Measurement {
	return NewMeasurement(name).AddFloat32(field, value)
}

// AddFloat32 field called name.
func (m Measurement) AddFloat32(name string, value float32) Measurement {
	m.fieldSet[name] = value
	return m
}

// MeasureFloat64 creates a new measurement with the given float64 field.
func MeasureFloat64(name string, field string, value float64) Measurement {
	return NewMeasurement(name).AddFloat64(field, value)
}

// AddFloat64 field called name.
func (m Measurement) AddFloat64(name string, value float64) Measurement {
	m.fieldSet[name] = value
	return m
}

// MeasureString creates a new measurement with the given string field.
func MeasureString(name string, field string, value string) Measurement {
	return NewMeasurement(name).AddString(field, value)
}

// AddString field called name.
func (m Measurement) AddString(name string, value string) Measurement {
	m.fieldSet[name] = value
	return m
}

// ToLineProtocal converts the metric to the influxdb line protocal.
func (m Measurement) ToLineProtocal() string {
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
