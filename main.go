package main

import (
	"database/sql"
	"encoding/json"
	"golang-inter/model"
	"log"
	"net/http"
)

var username, password, host, namaDB, defaultDB string
var db *sql.DB
var err error

func init() {
	username = "root"
	password = ""
	host = "localhost"
	namaDB = "gunadarma"
	defaultDB = "mysql"
}

func main() {
	db, err = model.Connect(username, password, host, namaDB)
	if err != nil {
		return
	}
	defer db.Close()

	http.HandleFunc("/mahasiswa", mahasiswa)
	http.HandleFunc("/mahasiswas", mahasiswas)
	log.Println("localhost : 8000")
	http.ListenAndServe(":8000", nil)
}

func mahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	npm := r.URL.Query()["npm"][0]
	data, err := model.Get(db, npm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
}

func mahasiswas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := model.GetAllMahasiswa(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
}
