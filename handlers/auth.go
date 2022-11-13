package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	authdto "waysbooks/dto/auth"
	dto "waysbooks/dto/result"
	"waysbooks/models"
	"waysbooks/pkg/bcrypt"
	jwtToken "waysbooks/pkg/jwt"
	"waysbooks/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	

	request := new(authdto.RegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(&request);	
	err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	password,err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	role := "user"

	user := models.User {
		Name : request.Name,
		Email: request.Email,
		Password: password,
		Role: role,
		Image: os.Getenv("PATH_FILE")+ "default.png",
	}

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: data}
	
	json.NewEncoder(w).Encode(response)

}

func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user := models.User{
		Email: request.Email,
		Password: request.Password,
	}


	// Check email
	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check Password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)

	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Generate Token

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 Hours Expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
		return
	}

	loginResponse := authdto.LoginResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Token: token,
		Role: user.Role,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: "Success", Data: loginResponse}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerAuth) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Check User by Id
	user, err := h.AuthRepository.Getuser(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	CheckAuthResponse := authdto.CheckAuthResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: "Success", Data: CheckAuthResponse}
	json.NewEncoder(w).Encode(response)
}