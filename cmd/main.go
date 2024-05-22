package main

import (
	"log"
	"net"
	"time"

	"github.com/codecrafters-io/redis-starter-go/pkg"
)

const dataPath = `store.dat`

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		return
	}
	defer listener.Close()
	store := pkg.NewStore()
	// clean expired keys
	store.StartCleaner(time.Second * 5)
	// load data when restart
	err = store.LoadFromFile(dataPath)
	if err != nil {
		log.Println(`Load from file error`)
	}
	// persist data
	go pkg.PersistStore(store, dataPath, time.Second*5)
	for {
		conn, err := listener.Accept()
		log.Println(`New connection: `, conn.RemoteAddr())
		if err != nil {
			continue
		}
		go pkg.HandleConnection(conn, store)
	}
}
