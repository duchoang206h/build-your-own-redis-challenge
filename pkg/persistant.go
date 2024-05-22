package pkg

import (
	"log"
	"time"
)

func PersistStore(s *Store, fileName string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		log.Println("Persist store")
		err := s.SaveToFile(fileName)
		if err != nil {
			log.Println("Error saving store to file:", err)
		}
	}
}
