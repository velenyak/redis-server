package resp

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Resp struct {
	reader *bufio.Reader
}

const (
	ARRAY       = '*'
	BULK_STRING = '$'
)

func New(reader *bufio.Reader) *Resp {
	return &Resp{reader}
}

func (r *Resp) Read() (string, error) {
	prefix, err := r.reader.ReadByte()
	if err != nil {
		return "", err
	}
	fmt.Println("To read prefix", string(prefix))
	switch prefix {
	case ARRAY:
		return r.readArray()
	case BULK_STRING:
		return r.readBulkString()
	default:
		fmt.Println("TODO: implement other types", prefix)
		return "", errors.New("TODO: implement other types")
	}
}

func (r *Resp) readLine() (string, error) {
	line, isPrefix, err := r.reader.ReadLine()
	fmt.Println("To read line", string(line), isPrefix, err)
	if isPrefix {
		for isPrefix {
			var nextLine []byte
			nextLine, isPrefix, err = r.reader.ReadLine()
			if err != nil {
				return "", err
			}
			line = append(line, nextLine...)
		}
	}
	if err != nil {
		return "", err
	}
	return string(line), nil
}

func (r *Resp) readInt() (int, error) {
	numStr, err := r.reader.ReadByte()
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(string(numStr))
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (r *Resp) readArray() (string, error) {
	len, err := r.readInt()
	if err != nil {
		return "", err
	}
	fmt.Println("To read array len", len)
	r.reader.ReadLine()
	var fullArray []string
	for i := 0; i < len; i++ {
		line, err := r.Read()
		if err != nil {
			return strings.Join(fullArray, ", "), err
		}
		fullArray = append(fullArray, line)
	}
	return strings.Join(fullArray, ", "), nil
}

func (r *Resp) readBulkString() (string, error) {
	len, err := r.readInt()
	if err != nil {
		return "", err
	}
	fmt.Println("To read bulk string len", len)
	r.reader.ReadLine()
	bulk := make([]byte, len)
	r.reader.Read(bulk)
	_, err = r.readLine()
	if err != nil {
		return "", err
	}
	bulkStr := string(bulk)
	fmt.Println("Bulk string", bulkStr)
	return bulkStr, nil
}

func Write(value string) (string, error) {
	return "+" + value + "\r\n", nil
}
