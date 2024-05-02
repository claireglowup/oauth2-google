package google

import (
	"log"
	"net/http"
)

func (cfg *config) GoogleLogin(w http.ResponseWriter, r *http.Request) {

	url := cfg.GoogleLoginConfig.AuthCodeURL("randomstate")
	log.Printf("INFO PATH %v, METHOD %v, CODE %v", r.URL.Path, r.Method, http.StatusTemporaryRedirect)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
