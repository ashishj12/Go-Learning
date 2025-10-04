// manage long running process

//classic example of a web server that when hit kicks off a potentially long-running process to fetch some data for it to return in the response.

// We will exercise a scenario where a user cancels the request before the data can be retrieved and we'll make sure the process is told to give up.

package main

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}
