package handler

import (
	"errors"

	"github.com/velenyak/redis-server/internal/resp"
)

func handleResp(input resp.Resp) (resp.Resp, error) {
	if input.ValueType != resp.RESP_ARRAY {
		return resp.Resp{}, errors.New("Expected array input")
	}
	return resp.Resp{}, errors.New("TODO: implement handler")
}

var handlers = map[string]func([]string) string{
	"PING":    ping,
	"SET":     set,
	"GET":     get,
	"HSET":    hset,
	"HGET":    hget,
	"HGETALL": hgetall,
	"COMMAND": cmd,
}

func ping(args []string) string {
	return "NOT IMPLEMENTED"
}

func set(args []string) string {
	return "NOT IMPLEMENTED"
}

func get(args []string) string {
	return "NOT IMPLEMENTED"
}

func hset(args []string) string {
	return "NOT IMPLEMENTED"
}

func hget(args []string) string {
	return "NOT IMPLEMENTED"
}

func hgetall(args []string) string {
	return "NOT IMPLEMENTED"
}

func cmd(args []string) string {
	return "NOT IMPLEMENTED"
}
