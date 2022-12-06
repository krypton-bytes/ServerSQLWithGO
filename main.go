package main

import "servidor/server"

func main() {
	serve := server.New(":8000")

	err := serve.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
