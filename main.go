package main

import (
	h "GoPI/Handlers"
	"fmt"
)

func main() {
	h.InitMux()
	e := h.NewEndpoint("/test", "this is a test", "GET")
	endpoints := h.NewEndpoints()
	endpoints.Bind(e)
	endpoints.BindListeners()

	fmt.Printf("Server is running....\n")

	h.Serve(8080)
}
