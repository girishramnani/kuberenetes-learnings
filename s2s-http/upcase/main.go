package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type ErrResponse struct {
	Error string
}

type SuccessResponse struct {
	Key     string
	Upcased string
}

type Request struct {
	Key string
}

func j(strct interface{}) []byte {
	data, _ := json.Marshal(strct)
	return data
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(j(ErrResponse{err.Error()}))
				return
			}
			var req Request
			log.Println(json.Unmarshal(data, &req))

			w.WriteHeader(http.StatusOK)
			w.Write(j(SuccessResponse{
				Key:     req.Key,
				Upcased: strings.ToUpper(req.Key),
			}))
			return
		} else {
			w.Write(j(ErrResponse{Error: "Only POST method allowed"}))
		}

	})

	fmt.Printf("Listening on port 8080\n")
	http.ListenAndServe(":8080", nil)

}
