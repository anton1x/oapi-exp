package main

import (
	openapi "github.com/anton1x/petstore/go"
	"github.com/gorilla/mux"
	"gorillagrpc/internal/middleware"
	"gorillagrpc/internal/repo"
	"gorillagrpc/internal/service"
	"io"
	"log"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
)

func main() {
	log.Printf("Server started")
	log.Println("=^.^=")

	PetApiService := service.NewPetApi(repo.NewPetRepoInmem())
	PetApiController := openapi.NewPetApiController(PetApiService)

	StoreApiService := openapi.NewStoreApiService()
	StoreApiController := openapi.NewStoreApiController(StoreApiService)

	UserApiService := openapi.NewUserApiService()
	UserApiController := openapi.NewUserApiController(UserApiService)

	api := openapi.NewRouter(PetApiController, StoreApiController, UserApiController)
	root := mux.NewRouter()

	root.HandleFunc("/debug/pprof/", pprof.Index)
	root.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	root.HandleFunc("/debug/pprof/profile", pprof.Profile)
	root.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	root.HandleFunc("/debug/pprof/trace", pprof.Trace)

	root.PathPrefix("/health").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.WriteHeader(200)
		io.WriteString(w, "im ok")
	}))

	root.NewRoute().Handler(api)

	api.Use(middleware.OnlyAuthorised)

	//router.GetRoute("AddPet").Handler(
	//	middleware.OnlyAuthorised(router.GetRoute("AddPet").GetHandler()),
	//)
	//
	//_ = middleware.InjectMiddlewareToNamedRoute(
	//	router, "AddPet", "DeletePet",
	//)(
	//	middleware.BlaBla,
	//	middleware.OnlyAuthorised,
	//)

	log.Fatal(http.ListenAndServe(":8089", root))

}
