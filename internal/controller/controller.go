package controller

import (
	"github.com/gorilla/mux"
	"github.com/viliusvv/cloud-lab/internal/controller"
	"fmt"
	"net/http"
	"os"
	"crypto/sha256"
)

func Sha256Handler(w http.ResponseWriter, r *http.Request) {
	// to get the username, use the following:
	username := mux.Vars(r)["username"];

	// to calculate sha256 hash of a string, use internal/hashes package and function GetHash
	sha256hash := sha256.Sum256([]byte(username));
	// to send the response, use the following:
	 fmt.Fprint(w, sha256hash);

	w.WriteHeader(http.StatusOK)
}

func GithubUsernameHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, os.Getenv("GITHUB_USERNAME"))
}
