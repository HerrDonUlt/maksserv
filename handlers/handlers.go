package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"practicegit/storage"
)

func handlerShowRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	record, err := storage.GetRecord(vars["id"])
	if err != nil {
		logAndEncode(w, "RecordStr " + vars["key"] + " is not valid")
		return
	}
	json.NewEncoder(w).Encode(record)
	log.Println("RecordStr " + vars["key"] + " showed")
}

//func handlerShowAllRecords(w http.ResponseWriter, r *http.Request) {
//	for _, s := range records {
//		record, err := storage.GetRecord()
//		json.NewEncoder(w).Encode(s)
//
//	}
//	log.Println("showed all records")
//}

func HandlerSetValueForRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]
	if storage.IsKeyInStorage(key) {
		storage.SetRecord(key, value)
		storage.AddLifetimeForRecord(key)
		logAndEncode(w, "RecordStr change value: " + value + " with key: " + key)
		return
	}
	storage.SetRecord(key, value)
	logAndEncode(w, "New record set with key: " + key)
}

func HandlerReturnValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	if storage.IsKeyInStorage(key) {
		if storage.IsValueInStorageNotNull(key) {
			record, err := storage.GetRecord(key)
			if err != nil {
				logAndEncode(w, "RecordStr " + vars["key"] + " is not valid")
				return
			}
			value := storage.GetRecordValue(record)
			storage.AddLifetimeForRecord(key)
			json.NewEncoder(w).Encode(value)
			log.Println("Value: " + value + "showed with key: " + key)
			return
		}
		logAndEncode(w, "Value is not exist")
		return
	}
	logAndEncode(w, "Key is not exist")
	return
}

//re
func HandlerDeleteRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	if storage.IsKeyInStorage(key) {
		if storage.IsValueInStorageNotNull(key) {
			storage.DeleteStorageRecord(key)
			logAndEncode(w, "RecordStr deleted with key: " + key)
			return
		}
		logAndEncode(w, "Value is not exist")
		return
	}
	logAndEncode(w, "Key is not exist")
	return
}

func logAndEncode(w http.ResponseWriter, s string) {
	json.NewEncoder(w).Encode(s)
	log.Println(s)
}

func InitHandlersAndStartServe() {
	r := mux.NewRouter()
	//r.HandleFunc("/all", handlerShowAllRecords)
	r.HandleFunc("/{key}", handlerShowRecord)
	r.HandleFunc("/setValue/{key}/{value}", HandlerSetValueForRecord)
	r.HandleFunc("/changeValue/{key}/{value}", HandlerReturnValue)
	r.HandleFunc("/delete/{key}", HandlerDeleteRecord)

	log.Fatal(http.ListenAndServe(":8080", r))
}
