package handler

import (
	"encoding/json"
	"net/http"
	"pagi/model"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

func (repo *MysqlDB) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var credential model.User

	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		http.Error(w, "error when getting body", http.StatusBadRequest)
		return
	}

	var storedUser model.User

	err = repo.DB.QueryRow("SELECT id, fullname, email, password, age, occupation, role FROM users WHERE email = ?", credential.Email).Scan(&storedUser.Id, &storedUser.Fullname, &storedUser.Email, &storedUser.Password, &storedUser.Age, &storedUser.Occupation, &storedUser.Role)
	if err != nil {
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(credential.Password))
	if err != nil {
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storedUser)
}
