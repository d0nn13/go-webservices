package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-martini/martini"
)

type response struct {
	Message string `json:"message"`
	Data    []byte `json:"data"`
}

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Post("/generate", func(req *http.Request, res http.ResponseWriter) (int, string) {
		log.Print("generating...")
		data, err := Generate(req.Body)
		res.Header().Add("Content-Type", "application/json")
		if err != nil {
			log.Print(err)
			r, _ := json.Marshal(response{Message: err.Error()})
			return 500, string(r[:])
		}
		r, _ := json.Marshal(response{Message: "ok", Data: data})
		return 200, string(r[:])
	})
	m.Run()
}
