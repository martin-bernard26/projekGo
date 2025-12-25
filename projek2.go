package main

import (
	"html/template"
	"net/http"
	"fmt"
)

// Data struktur untuk pesan
type Pesan struct {
	Nama string
	Isi  string
}

// Data yang akan dikirim ke HTML
type PageData struct {
	Judul  string
	Daftar []Pesan
}

var semuaPesan []Pesan

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Ambil data dari form
		nama := r.FormValue("nama")
		isi := r.FormValue("pesan")
		
		// Simpan ke slice
		semuaPesan = append(semuaPesan, Pesan{Nama: nama, Isi: isi})
	}

	// Load file HTML
	tmpl, _ := template.ParseFiles("index.html")
	
	data := PageData{
		Judul:  "Buku Tamu Komunitas Go",
		Daftar: semuaPesan,
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server aktif di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}