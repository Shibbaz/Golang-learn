package main

import(
	"net/http"
	"fmt"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store.Cancel()
		fmt.Fprint(w, store.Fetch())
	}
}

type Store interface {
	Fetch() string
	Cancel()
}

