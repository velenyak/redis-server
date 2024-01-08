package resp

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
)

type RespReader struct {
	reader *bufio.Reader
}

type RespType int

const (
	RESP_ERROR RespType = iota
	RESP_STRING
	RESP_INTEGER
	RESP_ARRAY
	RESP_BULK_STRING
)

type Resp struct {
	ValueType  RespType
	StrValue   string
	intValue   int
	ArrayValue []Resp
}

const (
	ARRAY       = '*'
	BULK_STRING = '$'
)

func New(reader *bufio.Reader) *RespReader {
	return &RespReader{reader}
}

func (r *RespReader) Read() (Resp, error) {
	prefix, err := r.reader.ReadByte()
	if err != nil {
		return Resp{}, err
	}
	fmt.Println("To read prefix", string(prefix))
	switch prefix {
	case ARRAY:
		return r.readArray()
	case BULK_STRING:
		return r.readBulkString()
	default:
		fmt.Println("TODO: implement other types", prefix)
		return Resp{}, errors.New("TODO: implement other types")
	}
}

func (r *RespReader) readLine() (string, error) {
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

func (r *RespReader) readInt() (int, error) {
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

func (r *RespReader) readArray() (Resp, error) {
	resp := Resp{
		ValueType: RESP_ARRAY,
	}
	len, err := r.readInt()
	if err != nil {
		return Resp{}, err
	}
	fmt.Println("To read array len", len)
	r.reader.ReadLine()
	resp.ArrayValue = make([]Resp, len)
	for i := 0; i < len; i++ {
		line, err := r.Read()
		if err != nil {
			return Resp{}, err
		}
		resp.ArrayValue[i] = line
	}
	return resp, nil
}

func (r *RespReader) readBulkString() (Resp, error) {
	len, err := r.readInt()
	if err != nil {
		return Resp{}, err
	}
	fmt.Println("To read bulk string len", len)
	r.reader.ReadLine()
	bulk := make([]byte, len)
	r.reader.Read(bulk)
	_, err = r.readLine()
	if err != nil {
		return Resp{}, err
	}
	bulkStr := string(bulk)
	fmt.Println("Bulk string", bulkStr)
	return Resp{
		ValueType: RESP_BULK_STRING,
		StrValue:  bulkStr,
	}, nil
}

func Write(value string) (string, error) {
	return "+" + value + "\r\n", nil
}
