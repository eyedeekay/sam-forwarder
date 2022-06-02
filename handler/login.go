package samtunnelhandler

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func (m *TunnelHandlerMux) Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if creds.Username != m.user {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if creds.Password != m.password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	m.sessionToken, err = GenerateRandomString(32)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   m.sessionToken,
		Expires: time.Now().Add(10 * time.Minute),
	})
}

func (m *TunnelHandlerMux) CSS(w http.ResponseWriter, r *http.Request) {
	if m.CheckCookie(w, r) == false {
		return
	}
	w.Header().Add("Content-Type", "text/css")
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte(fmt.Sprintf("%s\n", m.cssString)))
	w.Write([]byte(fmt.Sprintf("%s\n", DefaultCSS())))
}

func (m *TunnelHandlerMux) JS(w http.ResponseWriter, r *http.Request) {
	if m.CheckCookie(w, r) == false {
		return
	}
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte(fmt.Sprintf("%s\n", m.jsString)))
	w.Write([]byte(fmt.Sprintf("%s\n", DefaultJS())))
}
