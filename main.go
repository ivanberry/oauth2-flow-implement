package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Response custom response struct
type Response struct {
	http.ResponseWriter
}

// Message custom message
type Message struct {
	ID string `json:"app_id"` // json:  tell the struct how to populate
	Key string `json:"app_key"`
	Scope string `json:"scope"`
	RedirectUri string `json:"redirect_uri"`
}

func (r *Response) text(code int, body string) {
	r.Header().Set("Content-Type", "text/plain")
	r.WriteHeader(code)
	io.WriteString(r, fmt.Sprintf("%s\n", body))
}

func main() {

	// ping-pong service
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("pong")
	})

	// get token endpoint
	http.HandleFunc("/oauth2/authorize", func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		var msg Message
		err := decoder.Decode(&msg)
		resp := Response{w}

		if err != nil {
			panic(err)
		}

		if r.Method != "POST" {
			resp.text(http.StatusNotFound, "Not Found")
		} else {

			// 是否已经登陆，未登陆返回登陆页面

			if msg.RedirectUri == ""  {
				resp.text(http.StatusUnprocessableEntity, "Param redirect_uri Missing")
				return
			}

			if msg.ID == "" {
				resp.text(http.StatusUnprocessableEntity, "Param app id Missing")
				return
			}

			// use some field to generate code
			// and redirect to the redirect_uri
			http.Redirect(w, r, msg.RedirectUri, 301)
		}

		// check all param
		// app_id
		// app_key
		// redirect_url
		// scope
		// state[CSRF]

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
