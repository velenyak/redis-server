package handler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/velenyak/redis-server/internal/resp"
	"github.com/velenyak/redis-server/internal/storage"
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
	if len(args) != 2 {
		return resp.Resp{
			ValueType: resp.RESP_ERROR,
			StrValue:  "ERR wrong number of arguments for 'set' command",
		}
	}

	storage.Set(args[0].StrValue, args[1].StrValue)

	return resp.Resp{
		ValueType: resp.RESP_STRING,
		StrValue:  "OK",
	}
}

func get(args []resp.Resp) resp.Resp {
	if len(args) != 1 {
		return resp.Resp{
			ValueType: resp.RESP_ERROR,
			StrValue:  "ERR wrong number of arguments for 'get' command",
		}
	}

	value, _ := storage.Get(args[0].StrValue)

	return resp.Resp{
		ValueType: resp.RESP_STRING,
		StrValue:  value,
	}
}

func hset(args []resp.Resp) resp.Resp {
	if len(args) < 3 {
		return resp.Resp{
			ValueType: resp.RESP_ERROR,
			StrValue:  "ERR wrong number of arguments for 'hset' command",
		}
	}

	hashName := args[0].StrValue
	for i := 1; i < len(args); i += 2 {
		if args[i].ValueType != resp.RESP_BULK_STRING || args[i+1].ValueType != resp.RESP_BULK_STRING {
			return resp.Resp{
				ValueType: resp.RESP_ERROR,
				StrValue:  "ERR wrong number of arguments for 'hset' command",
			}
		}
		storage.HSet(hashName, args[i].StrValue, args[i+1].StrValue)
	}

	return resp.Resp{
		ValueType: resp.RESP_STRING,
		StrValue:  "OK",
	}
}

func hget(args []resp.Resp) resp.Resp {
	if len(args) != 2 {
		return resp.Resp{
			ValueType: resp.RESP_ERROR,
			StrValue:  "ERR wrong number of arguments for 'hget' command",
		}
	}

	value, _ := storage.HGet(args[0].StrValue, args[1].StrValue)

	return resp.Resp{
		ValueType: resp.RESP_STRING,
		StrValue:  value,
	}
}

func hgetall(args []resp.Resp) resp.Resp {
	if len(args) != 1 {
		return resp.Resp{
			ValueType: resp.RESP_ERROR,
			StrValue:  "ERR wrong number of arguments for 'hgetall' command",
		}
	}

	value, _ := storage.HGetAll(args[0].StrValue)

	values := []resp.Resp{}
	for k, v := range value {
		values = append(values, resp.Resp{
			ValueType: resp.RESP_BULK_STRING,
			StrValue:  k,
		})
		values = append(values, resp.Resp{
			ValueType: resp.RESP_BULK_STRING,
			StrValue:  v,
		})
	}

	return resp.Resp{
		ValueType:  resp.RESP_ARRAY,
		ArrayValue: values,
	}
}

func cmd(args []resp.Resp) resp.Resp {
	return resp.Resp{
		ValueType: resp.RESP_STRING,
		StrValue:  "OK",
	}
}
