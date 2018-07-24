package storage

import (
	"sync"
)

const stdlifetime int = 3

type Record struct {
	Id      string `json:"id"`
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	Address int    `json:"address"`
}

var storage = struct {
	sync.RWMutex
	records map[string]Record
}{records: make(map[string]Record)}

func IsKeyInStorage() bool {
	storage.RLock()
	defer storage.RUnlock()
	for k := range storage.records {
		if k == r.Key {
			return true
		}
	}
	return false
}

func IsValueInStorageNotNull(key string) bool {
	storage.RLock()
	defer storage.RUnlock()
	r := storage.records[key]
	if value != "" {
		return true
	}
	return false
}

func GetRecord(key string) (Record, bool) {
	storage.Lock()
	defer storage.Unlock()
	r := storage.records[key]
	if r.LifeTime != 0 {
		return storage.records[key], true
	}
	return Record{}, false
}

//func GetAllRecord() map[string]RecordStr {
//	storage.RLock()
//	defer storage.RUnlock()
//	return storage.records
//}

func SetRecord(id, name, email, address string) {
	storage.Lock()
	defer storage.Unlock()
	storage.records[id] = Record{id, name,email, address}
}

func DeleteStorageRecord(key string) {
	storage.Lock()
	defer storage.Unlock()
	delete(storage.records, key)
}
