package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mali3days/memberclub/pkg/mocks"
)

func GetAllMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Members)
}
