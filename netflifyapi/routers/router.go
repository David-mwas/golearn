package routers

import (
	"github.com/David-mwas/golearn/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/movies", controllers.GetAllMovies).Methods("GET")

	router.HandleFunc("/api/v1/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/v1/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/v1/movie/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/v1/movies", controllers.DeleteAllMovie).Methods("DELETE")

	return router
}
