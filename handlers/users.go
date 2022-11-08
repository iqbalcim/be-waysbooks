package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	dto "waysbooks/dto/result"
	usersdto "waysbooks/dto/users"
	"waysbooks/models"
	"waysbooks/pkg/bcrypt"
	"waysbooks/repositories"

	"github.com/gorilla/mux"
)

type handler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

func (h *handler) FindUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUsers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, p := range users {
		users[i].Image = os.Getenv("PATH_FILE") + p.Image
	}


	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: users}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user,err := h.UserRepository.GetUser(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user.Image = os.Getenv("PATH_FILE") + user.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: convertResponse(user)}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	fmt.Println(filename)

	request := usersdto.UpdateUserRequest{
		Name: r.FormValue("fullName"),
		Email: r.FormValue("email"),
		Password: r.FormValue("password"),
		Phone: r.FormValue("phone"),
		Gender: r.FormValue("gender"),
		Image: filename,
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	user := models.User{}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Password != "" {
		user.Password = password
	}

	if request.Phone != "" {
		user.Phone = request.Phone
	}

	if request.Gender != "" {
		user.Gender = request.Gender
	}

	if request.Image != "" {
		user.Image = filename
	}

	data, err := h.UserRepository.UpdateUser(user, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)

}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.UserRepository.DeleteUser(user, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusText(http.StatusOK), Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponse(u models.User) usersdto.UserResponse{
	return usersdto.UserResponse{
		ID: u.ID,
		Name: u.Name,
		Email: u.Email,
		Gender: u.Gender,
		Phone: u.Phone,
		Image: u.Image,
	}
}