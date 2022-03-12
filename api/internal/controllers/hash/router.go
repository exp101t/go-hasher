package hash

import (
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/exp101t/go-hasher/api/internal/models"
)

func getHasherHandler(hasher func([]byte) []byte) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.HashRequest
		var resp models.HashResponse

		err := json.NewDecoder(r.Body).Decode(&req)

		if err == nil {
			rawHash := hasher([]byte(req.Data))
			hexHash := hex.EncodeToString(rawHash[:])

			resp = models.HashResponse{Hash: hexHash}
		} else {
			resp = models.HashResponse{Error: err.Error()}
		}

		json.NewEncoder(w).Encode(resp)
	}
}

func NewHasherRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/sha256",
		getHasherHandler(hashWithSha256)).Methods("POST")
	router.HandleFunc("/streebog256",
		getHasherHandler(hashWithStreebog256)).Methods("POST")
	router.HandleFunc("/streebog512",
		getHasherHandler(hashWithStreebog512)).Methods("POST")

	return router
}
