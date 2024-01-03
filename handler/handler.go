package handler

var Handlers = map[string]func([]string) string{
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
