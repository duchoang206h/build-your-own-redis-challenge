package pkg

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func HandleConnection(conn net.Conn, store *Store) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		input = strings.TrimSpace(input)
		commands := strings.SplitN(input, " ", 4)
		log.Println(`commands: `, commands, len(commands))
		res := HandleCommand(commands, store)
		conn.Write(res)
	}
}
