package handler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/velenyak/redis-server/internal/resp"
)

func HandleResp(input resp.Resp) (resp.Resp, error) {
	if input.ValueType != resp.RESP_ARRAY {
		return resp.Resp{}, errors.New("Expected array input")
	}
	if len(input.ArrayValue) == 0 {
		return resp.Resp{}, errors.New("Expected non-empty array input")
	}
	cmd := input.ArrayValue[0]
	args := input.ArrayValue[1:]

	fmt.Println("To handle command", cmd.StrValue)
	fmt.Println("To handle args", args)

	handler, ok := handlers[strings.ToUpper(cmd.StrValue)]
	if !ok {
		return resp.Resp{}, errors.New("Unknown command")
	}

	return handler(args), nil

}

var handlers = map[string]func([]resp.Resp) resp.Resp{
	"PING":    ping,
	"SET":     set,
	"GET":     get,
	"HSET":    hset,
	"HGET":    hget,
	"HGETALL": hgetall,
	"COMMAND": cmd,
}

func ping(args []resp.Resp) resp.Resp {
	return resp.Resp{
		ValueType: resp.RESP_STRING,
		StrValue:  "PONG",
	}
}

func set(args []resp.Resp) resp.Resp {
	return resp.Resp{
		ValueType: resp.RESP_ERROR,
		StrValue:  "Not implemented",
	}
}

func get(args []resp.Resp) resp.Resp {
	return resp.Resp{
		ValueType: resp.RESP_ERROR,
		StrValue:  "Not implemented",
	}
}

func hset(args []resp.Resp) resp.Resp {
	return resp.Resp{
		ValueType: resp.RESP_ERROR,
		StrValue:  "Not implemented",
	}
}

func hget(args []resp.Resp) resp.Resp {
	return resp.Resp{
		ValueType: resp.RESP_ERROR,
		StrValue:  "Not implemented",
	}
}

func hgetall(args []resp.Resp) resp.Resp {
	return resp.Resp{
		ValueType: resp.RESP_ERROR,
		StrValue:  "Not implemented",
	}
}

func cmd(args []resp.Resp) resp.Resp {
	return resp.Resp{
		ValueType: resp.RESP_STRING,
		StrValue:  "OK",
	}
}
