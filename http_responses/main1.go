package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type Profile struct {
	Name    string
	Hobbies []string
}

type XMLProfile struct {
	Name    string
	Hobbies []string `xml:"Hobbies>Hobby"`
}

func main() {
	http.HandleFunc("/default", foo)
	http.HandleFunc("/plain", foo1)
	http.HandleFunc("/json", foo2)
	http.HandleFunc("/xml", foo3)
	http.HandleFunc("/file", foo4)
	http.HandleFunc("/html", foo5)
	http.HandleFunc("/htmlbuffer", foo6)
	http.HandleFunc("/templates", foo7)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}

func foo1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func foo2(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"lucas", []string{"cooking", "eating"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func foo3(w http.ResponseWriter, r *http.Request) {
	profile := XMLProfile{"lucas", []string{"cooking"}}

	x, err := xml.MarshalIndent(profile, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}

func foo4(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("images", "triple_pat.png")
	http.ServeFile(w, r, fp)
}

func foo5(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"lucas", []string{"cooking"}}

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func foo6(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"lucas", []string{"cooking"}}

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	templateString := buf.String()
	fmt.Println(templateString)
}

func foo7(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"lucas", []string{"cooking"}}

	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", "index1.html")

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
