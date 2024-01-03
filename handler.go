package handler

import (
    "errors"
)

var Handlers = map[string]func([]Value) Value{
	"PING":    ping,
	"SET":     set,
	"GET":     get,
	"HSET":    hset,
	"HGET":    hget,
	"HGETALL": hgetall,
	"COMMAND": cmd,
}

func ping(args []Value) Value {
    errors.New("Not implemented")
}

func set(args []Value) Value {
    errors.New("Not implemented")
}

func get(args []Value) Value {
    errors.New("Not implemented")
}

func hset(args []Value) Value {
    errors.New("Not implemented")
}

func hget(args []Value) Value {
    errors.New("Not implemented")
}

func hgetAll(args []Value) Value {
    errors.New("Not implemented")
}

func cmd(args []Value) Value {
    errors.New("Not implemented")
}

