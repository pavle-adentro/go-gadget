package handlers

import (
	"net/http"
	"os"
	"path"
	"text/template"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {

	clientId := os.Getenv("CLIENT_ID")
	redirectUri := os.Getenv("REDIRECT_URI")
	oauthUrl := "https://accounts.google.com/o/oauth2/v2/auth?client_id=" + clientId + "&response_type=token&scope=email profile+https://www.googleapis.com/auth/userinfo.profile&redirect_uri=" + redirectUri
	http.Redirect(w, r, oauthUrl, http.StatusSeeOther)
}

func AuthCallBackHandler(w http.ResponseWriter, r *http.Request) {

	fp := path.Join("templates", "success.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
