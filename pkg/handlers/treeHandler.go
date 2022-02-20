package handlers

import (
	"ecosia/pkg/utils"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type treeResponse struct {
	MyFavouriteTree string `json:"myFavouriteTree"`
}

func TreeHandlerWithParams(w http.ResponseWriter, r *http.Request) {

	response := treeResponse{MyFavouriteTree: chi.URLParam(r, "name")}
	err := utils.ConvertToJson(w, http.StatusOK, response)
	if err != nil {
		log.Println(err)
	}
}

func TreeHandler(w http.ResponseWriter, r *http.Request) {

	response := treeResponse{MyFavouriteTree: "Baobab"}

	err := utils.ConvertToJson(w, http.StatusOK, response)
	if err != nil {
		log.Println(err)
	}

}
