package main

import (
	"html/template"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
	"log"
)
type Option struct {
	Text string
	Arc  string
}

type Block struct {
	Anchor  string
	Title   string
	Story   []string
	Options []Option
}

func main() {

	raw, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	var blocks []Block
	err = json.Unmarshal(raw, &blocks)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}


	lookup := make(map[string]Block)
	for _, x := range blocks {
		lookup[x.Anchor] = x
	}

	tmpl := template.Must(template.ParseFiles("layout.html"))

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a := r.URL.Path
		if a == "/" {
			a = "/intro"
		}
		data := lookup[a[1:len(a)]]
		tmpl.Execute(w, data)
	})

	err = http.ListenAndServe(":9999", h)
	log.Fatal(err)
}

