package pkg

import (
	"fmt"
	"strconv"
)

var (
	missingCommandMsg  = "ERR command missing\n"
	nilMsg             = "nil\n"
	okMsg              = "Ok\n"
	keyNotfoundMsg     = "key not found\n"
	commandNotfoundMsg = "command not found\n"
	pongMsg            = "PONG\r\n"
)

func HandleCommand(commands []string, store *Store) []byte {
	var res []byte
	switch commands[0] {
	case "PING":
		res = []byte(pongMsg)
	case "GET":
		if len(commands) < 2 {
			res = []byte(missingCommandMsg)
			break
		}
		key := commands[1]
		val, err := store.Get(key)
		if err != nil {
			res = []byte(nilMsg)
			break
		}
		res = []byte(convertToString(val))
	case "SET":
		if len(commands) < 4 {
			res = []byte(missingCommandMsg)
			break
		}
		// SET key val exp
		key := commands[1]
		val := commands[2]
		exp, _ := strconv.Atoi(commands[3])
		store.Set(key, val, int64(exp))
		res = []byte(okMsg)
	case "DEL":
		if len(commands) < 2 {
			res = []byte(missingCommandMsg)
			break
		}
		// DEL key
		key := commands[1]
		err := store.Del(key)
		if err != nil {
			res = []byte(keyNotfoundMsg)
			break
		}
		res = []byte(okMsg)
	default:
		res = []byte(commandNotfoundMsg)
	}
	return res
}

func convertToString(val interface{}) string {
	return fmt.Sprintf("%v\n", val)
}
