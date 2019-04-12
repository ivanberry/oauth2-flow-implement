package main

import (
	"fmt"
	"log"
	"net/http"
)

// Response custom response struct
type Response struct {
	http.ResponseWriter
}

func main() {

	// get token endpoint
	http.HandleFunc("/oauth2/token", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			fmt.Fprint(w, "Error Method verber")
		} else {
			resp := Response{w}
			resp.Text(http.StatusOK, fmt.Fprint("Hello, %s", "get token"))
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
