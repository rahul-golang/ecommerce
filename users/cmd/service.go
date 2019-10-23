package cmd

import (
	"flag"
	"fmt"

	"net/http"

	"github.com/rahul-golang/ecommerce/users/pkg/httphandler"

	"github.com/rahul-golang/ecommerce/users/pkg/service"

	"github.com/gorilla/mux"
	"github.com/rahul-golang/ecommerce/database"
	"github.com/rahul-golang/ecommerce/log"
	"github.com/rahul-golang/ecommerce/users/pkg/endpoints"
	"github.com/rahul-golang/ecommerce/users/pkg/repository"
)

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var fs = flag.NewFlagSet("users", flag.ExitOnError)
var debugAddr = fs.String("debug.addr", ":8080", "Debug and metrics listen address")
var httpAddr = fs.String("http-addr", ":8081", "HTTP listen address")
var grpcAddr = fs.String("grpc-addr", ":8082", "gRPC listen address")

func Run() {

	router := mux.NewRouter()

	//Injected dependancies
	dataStoreInterface := database.NewMySQLClientConn()
	userRepositoryInterface := repository.NewUserRepository(dataStoreInterface)
	usersServiceInterface := service.NewBasicUsersService(userRepositoryInterface)
	restHandler := httphandler.NewUserRestHandler(usersServiceInterface)
	router.Use(loggingMiddleware)
	endpoints.NewUserEndpoints(router, restHandler)
	fmt.Println("Server running on port : 8081")
	fmt.Println(http.ListenAndServe(":8081", router))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Do stuff here
		//	log.Println(r.RequestURI)
		req = req.WithContext(log.WithRqID(req.Context()))

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, req)
	})
}
