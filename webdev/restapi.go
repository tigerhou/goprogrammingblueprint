package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateOn    time.Time `json:"createdon"`
}

const contentType = "Content-Type"
const jsonType = "application/json"

var noteStore = make(map[string]Note)

var id int = 0

func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	note.CreateOn = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note
	noteJson, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(noteJson)
}

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}
	w.Header().Set(contentType, jsonType)

	noteJson, err := json.Marshal(notes)

	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(noteJson)
}

func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	k := vars["id"]
	var noteToUpd Note
	err = json.NewDecoder(r.Body).Decode(&noteToUpd)
	if err != nil {
		panic(err)
	}
	if note, ok := noteStore[k]; ok {
		noteToUpd.CreateOn = note.CreateOn
		delete(noteStore, k)
		noteStore[k] = noteToUpd
	} else {
		log.Printf("Could not find key of Note %s to delete", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if _, ok := noteStore[id]; ok {
		delete(noteStore, id)
	} else {
		log.Printf("Could not find key of Note %s to delete", id)
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	router.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	router.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	router.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
