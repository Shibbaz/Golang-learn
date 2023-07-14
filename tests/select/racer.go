package main

import(
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	aDuraction := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuraction < bDuration {
		return a
	}
	return b
}