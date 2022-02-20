package handlers

import (
	"ecosia/pkg/utils"
	"log"
	"net/http"
	"os"
)

const version = "0.1.0"

type appStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {

	appInfo := appStatus{
		Status:      "Available",
		Environment: os.Getenv("ENVIRONMENT"),
		Version:     version,
	}

	err := utils.ConvertToJsonWithKey(w, http.StatusOK, appInfo, "info")
	if err != nil {
		log.Println(err)
	}

}
