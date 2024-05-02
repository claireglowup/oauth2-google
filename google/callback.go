package google

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type user struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func (cfg *config) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	if state != "randomstate" {
		http.Error(w, "state doesn't match", http.StatusInternalServerError)
		return

	}

	code := r.URL.Query().Get("code")

	token, err := cfg.GoogleLoginConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "state doesn't match", http.StatusInternalServerError)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		http.Error(w, "state doesn't match", http.StatusInternalServerError)
		return

	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "state doesn't match", http.StatusInternalServerError)
		return

	}
	var userJson user
	err = json.Unmarshal(userData, &userJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("INFO PATH %v, METHOD %v, CODE %v", r.URL.Path, r.Method, http.StatusOK)

	// Mengatur header respons untuk menunjukkan bahwa konten adalah JSON
	w.Header().Set("Content-Type", "application/json")

	// Mengirim respons JSON
	json.NewEncoder(w).Encode(userJson)

}
