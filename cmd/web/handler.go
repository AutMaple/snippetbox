package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }
  files := []string{
    "./ui/html/base.html",
    "./ui/html/pages/home.html"
  }
  ts, err := template.ParseFiles("./ui/html/pages/home.tmpl")
  if err != nil {
    log.Print(err.Error())
    http.Error(w, "Internal Server Error!", http.StatusInternalServerError)
    return
  }
  err = ts.Execute(w, nil)
  if err != nil {
    log.Printf(err.Error())
    http.Error(w, "Internal Server Error!", http.StatusInternalServerError)
  }
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    http.NotFound(w, r)
    return
  }
  fmt.Fprintf(w, "Display a specific snippet with ID %d ...", id)
  w.Write([]byte("Display a specific snippet ..."))
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    w.Header().Set("Allow", http.MethodPost)
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    return
  }
  w.Write([]byte("Create a new snippet..."))
}
