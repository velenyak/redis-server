package resp

import (
	"errors"
	"fmt"
	"strings"
)

func Write(value Resp) (string, error) {
	fmt.Println("To write", value.ValueType, value.StrValue, value.ArrayValue)
	switch value.ValueType {
	case RESP_STRING, RESP_ERROR, RESP_INTEGER:
		return string(value.ValueType) + value.StrValue + "\r\n", nil
	case RESP_BULK_STRING:
		return writeBulkString(value)
	case RESP_ARRAY:
		return writeArray(value)
	default:
		return "", errors.New("Unknown value type")
	}
}

func writeBulkString(value Resp) (string, error) {
	prefix := string(RESP_BULK_STRING)
	length := len(value.StrValue)

	return fmt.Sprintf("%s%d\r\n%s\r\n", prefix, length, value.StrValue), nil
}

func writeArray(value Resp) (string, error) {
	prefix := string(RESP_ARRAY)
	length := len(value.ArrayValue)

	var lines []string
	for _, v := range value.ArrayValue {
		newLine, err := Write(v)
		if err != nil {
			return "", err
		}
		lines = append(lines, newLine)
	}

	return fmt.Sprintf("%s%d\r\n%s", prefix, length, strings.Join(lines, "")), nil
}
