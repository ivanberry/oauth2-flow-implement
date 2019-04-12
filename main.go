package main

import (
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
	appID string
	scope string
	state string
}

func (r *Response) text(code int, body string) {
	r.Header().Set("Content-Type", "text/plain")
	r.WriteHeader(code)
	io.WriteString(r, fmt.Sprintf("%s\n", body))
}

func main() {

	// get token endpoint
	http.HandleFunc("/oauth2/authorize", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.FormValue("app_id"))
		resp := Response{w}

		// populate msg
		msg := &Message{
			appID: r.FormValue("app_id"),
			scope: r.FormValue("scope"),
			state: r.FormValue("state"),
		}

		fmt.Println(msg)

		if r.Method != "POST" {
			resp.text(http.StatusNotFound, "Not Found")
		} else {
			resp.text(http.StatusOK, "Hello Token")
		}

		// check all paramters
		// app_id
		// app_key
		// redirect_url
		// scope
		// state[CSRF]

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
