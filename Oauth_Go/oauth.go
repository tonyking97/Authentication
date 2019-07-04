package main

import (
	//go get golang.org/x/oauth2 --- get oauth2 package using this command
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/linkedin"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	googleOauthConfig   *oauth2.Config
	linkedinOauthConfig *oauth2.Config
	githubOauthConfig   *oauth2.Config

	oauthStateString = "pseudo-random"
)

func init() {

	os.Setenv("GOOGLE_CLIENT_ID", "1010888033141-4mgp4i3qe046lik8bkl6l5cftblvh1hd.apps.googleusercontent.com")
	os.Setenv("GOOGLE_CLIENT_SECRET", "Bm849VhYtyzTpCcJGgutrhuy")

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	os.Setenv("LINKEDIN_CLIENT_ID", "81s4pgfbuyhs0b")
	os.Setenv("LINKEDIN_CLIENT_SECRET", "AtSOtSe0ZwccMjdK")

	linkedinOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/linkedin/callback",
		ClientID:     os.Getenv("LINKEDIN_CLIENT_ID"),
		ClientSecret: os.Getenv("LINKEDIN_CLIENT_SECRET"),
		Scopes:       []string{"r_emailaddress", "r_liteprofile", "w_member_social"},
		Endpoint:     linkedin.Endpoint,
	}

	os.Setenv("GITHUB_CLIENT_ID", "bd23751748b3a1034998")
	os.Setenv("GITHUB_CLIENT_SECRET", "e9337a7aefaa709ac31e8cdc86ae7adb6142941f")

	githubOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/github/callback",
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{"user"},
		Endpoint:     github.Endpoint,
	}

}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/google/login", handleGoogleLogin)
	http.HandleFunc("/linkedin/login", handleLinkedinLogin)
	http.HandleFunc("/github/login", handleGithubLogin)

	http.HandleFunc("/google/callback", handleGoogleCallback)
	http.HandleFunc("/linkedin/callback", handleLinkedinCallback)
	http.HandleFunc("/github/callback", handleGithubCallback)

	fmt.Println(http.ListenAndServe(":8080", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html>
<body>
	<a href="/google/login">Google Log In</a><br>
	<a href="/linkedin/login">Linkedin Log In</a><br>
	<a href="/github/login">Github Log In</a>
</body>
</html>`

	fmt.Fprintf(w, htmlIndex)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleLinkedinLogin(w http.ResponseWriter, r *http.Request) {
	url := linkedinOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGithubLogin(w http.ResponseWriter, r *http.Request) {
	url := githubOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("state") != oauthStateString {
		fmt.Println("State is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))

	if err != nil {
		fmt.Println("Could not get token:,%s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	//fmt.Println(token.AccessToken)

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	if err != nil {
		fmt.Println("Could not create get request:,%s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not parse response:,%s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, string(content))
	fmt.Println(string(content))
}

func handleLinkedinCallback(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("state") != oauthStateString {
		fmt.Println("State is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := linkedinOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		fmt.Println("Could not get token:,%s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.linkedin.com/v2/me", nil)
	req.Header.Add("Connection", `Keep-Alive`)
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	resp, err := client.Do(req)

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not parse response:,%s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, string(content))
	fmt.Println(string(content))
}

func handleGithubCallback(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("state") != oauthStateString {
		fmt.Println("State is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	token, err := githubOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))

	if err != nil {
		fmt.Println("Could not get token:,%s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://api.github.com/user?access_token=" + token.AccessToken)

	if err != nil {
		fmt.Println("Could not create get request:,%s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not parse response:,%s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, string(content))
	fmt.Println(string(content))
}
