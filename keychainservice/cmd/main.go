package main

import (
	"fmt"
	"keychainservice/controllers"
	"keychainservice/services"
	"log"
	"net/http"
	"reflect"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Key chain service is live")
}

func main() {

	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	http.HandleFunc(
		"/users",
		controllers.NewHandlerBuilder().
			AddValidation().
			AddAuthentication().
			Build(),
	)

	// http.HandleFunc("/users", userController.CreateUser)
	http.HandleFunc("/users", controllers.NewHTTPController(
		http.MethodPost,
		userController.CreateUser,
		reflect.TypeOf(controllers.UserStructBody{}),
		nil,
	))

	// Registering the handler for the root URL "/"
	http.HandleFunc("/", helloHandler)

	// Starting the server on port 8080
	fmt.Println("Server is starting on port 8080...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
