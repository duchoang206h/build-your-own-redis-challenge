package pkg

import (
	"encoding/gob"
	"errors"
	"log"
	"os"
	"sync"
	"time"
)

type (
	Store struct {
		data map[string]Value
		mu   sync.RWMutex
	}
	Value struct {
		Data       interface{}
		Expiration int64
	}
)

func NewStore() *Store {
	return &Store{
		data: make(map[string]Value),
	}
}

func (s *Store) Get(key string) (interface{}, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	if !ok {
		return nil, errors.New(`key not found`)
	}
	return val.Data, nil
}

func (s *Store) Set(key string, val interface{}, expInSecond int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	exp := time.Now().Add(time.Second * time.Duration(expInSecond)).UnixNano()
	s.data[key] = Value{
		Data:       val,
		Expiration: exp,
	}
}

func (s *Store) Del(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.data[key]
	if !ok {
		return errors.New(`key not found`)
	}
	delete(s.data, key)
	return nil
}

func (s *Store) StartCleaner(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)
			s.cleanExpiredKeys()
		}
	}()
}

// Clean expired key-val in interval time
func (s *Store) cleanExpiredKeys() {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().UnixNano()
	for key, val := range s.data {
		if val.Expiration > 0 && val.Expiration < now {
			log.Printf(`key %s expired`, key)
			delete(s.data, key)
		}
	}
}

func (s *Store) SaveToFile(filename string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(s.data)
}

func (s *Store) LoadFromFile(filename string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(&s.data)
}
