package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/knid/ilet/internal/models"
	"github.com/knid/ilet/internal/utils"
)

func (h *HTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetUserFromRequest(r, h.DB)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, user)
}

func (h *HTTPHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username       string `json:"username"`
		Password       string `json:"password"`
		PasswordRepeat string `json:"password_repeat"`
	}

	body, err := io.ReadAll(r.Body)
	json.Unmarshal(body, &credentials)

	if credentials.Password != credentials.PasswordRepeat {
		utils.JsonErrorResponse(w, http.StatusBadRequest, errors.New("passwords does not match"))
		return
	}

	hasher := sha256.New()
	hasher.Write([]byte(credentials.Password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	user := models.User{
		Username: credentials.Username,
		Password: string(hashedPassword),
	}

	createdUser, err := h.DB.CreateUser(user)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	utils.JsonResponse(w, http.StatusCreated, createdUser)
}

func (h *HTTPHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	body, err := io.ReadAll(r.Body)
	json.Unmarshal(body, &credentials)

	hasher := sha256.New()
	hasher.Write([]byte(credentials.Password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	user, err := h.DB.GetUserByCredentials(credentials.Username, string(hashedPassword))
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if token, err := h.DB.GetTokenByUser(user); err == nil {
		utils.JsonResponse(w, http.StatusOK, token)
		return
	}

	token := models.Token{
		User:  user,
		Token: utils.GenerateRandomString(64),
	}

	if _, err := h.DB.CreateToken(token); err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	utils.JsonResponse(w, http.StatusCreated, token)
}

func (h *HTTPHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetUserFromRequest(r, h.DB)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	var credentials struct {
		username       string
		password       string
		passwordRepeat string
	}

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&credentials)

	if credentials.password != credentials.passwordRepeat {
		utils.JsonErrorResponse(w, http.StatusBadRequest, errors.New("passwords does not match"))
		return
	}

	hasher := sha256.New()
	hasher.Write([]byte(credentials.password))
	hashedPassword := hasher.Sum(nil)

	user.Username = credentials.username
	user.Password = string(hashedPassword)

	updatedUser, err := h.DB.UpdateUser(user)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, updatedUser)
}

func (h *HTTPHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetUserFromRequest(r, h.DB)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	h.DB.DeleteUser(user)
}
