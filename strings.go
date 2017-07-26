package telegraf

import (
	"fmt"
	"strconv"
	"strings"
)

var measurementEscaper = strings.NewReplacer(`,`, `\,`, ` `, `\ `)
var keyEscaper = strings.NewReplacer(`,`, `\,`, ` `, `\ `, `=`, `\=`)
var tagValueEscaper = keyEscaper

//var fieldValueEscaper = strings.NewReplacer(`"`, `\"`)

func (m Metric) String() string {
	line := measurementEscaper.Replace(m.measurement)
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
