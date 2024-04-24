package auth

import (
	//"apiServerBook/data"
	"encoding/json"
	"fmt"
	"github.com/shuvo-14/api-server/db"
	"net/http"
	"time"
)

func LogIn(w http.ResponseWriter, r *http.Request) {
	var tmp db.Credential
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&tmp)
	if err != nil {
		http.Error(w, "Cannot Decode", http.StatusBadRequest)
		return
	}
	cred, okay := db.CredList[tmp.UserName]
	if !okay {
		http.Error(w, "User Name doesn't Exist", http.StatusBadRequest)
		return
	}
	if cred.Password != tmp.Password {
		http.Error(w, "Password didn't matched", http.StatusBadRequest)
		return
	}
	et := time.Now().Add(15 * time.Minute)
	_, tokenString, err := db.TokenAuth.Encode(map[string]interface{}{
		"aud": "Parvej Mia",
		"exp": et.Unix(),
	})
	fmt.Println(tokenString)
	if err != nil {
		http.Error(w, "Can not generate jwt", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   tokenString,
		Expires: et,
	})
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Successfully Logged In"))
	if err != nil {
		http.Error(w, "Can not Write data", http.StatusInternalServerError)
		return
	}

}
func LogOut(w http.ResponseWriter, _ *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Expires: time.Now(),
	})
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Successfully Logged Out"))
	if err != nil {
		http.Error(w, "Can not Write data", http.StatusInternalServerError)
		return
	}
}
