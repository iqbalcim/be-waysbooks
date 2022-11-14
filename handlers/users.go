package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	dto "waysbooks/dto/result"
	usersdto "waysbooks/dto/users"
	"waysbooks/models"
	"waysbooks/pkg/bcrypt"
	"waysbooks/repositories"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gorilla/mux"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) FindUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUsers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, p := range users {
		users[i].Image =  p.Image
	}


	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: users}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) GetUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user,err := h.UserRepository.GetUser(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: convertResponse(user)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) UpdateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
  	filepath := dataContex.(string)

	  var ctx = context.Background()
	  var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	  var API_KEY = os.Getenv("API_KEY")
	  var API_SECRET = os.Getenv("API_SECRET")

	  	// Add your Cloudinary credentials ...
cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

// Upload file to Cloudinary ...
resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "waysbooks"});

if err != nil {
	panic(err)
}

	request := usersdto.UpdateUserRequest{
		Name: r.FormValue("fullName"),
		Email: r.FormValue("email"),
		Password: r.FormValue("password"),
		Phone: r.FormValue("phone"),
		Gender: r.FormValue("gender"),
		Address: r.FormValue("address"),
		Image: resp.SecureURL,
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
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

	if request.Address != "" {
		user.Address = request.Address
	}

	if request.Image != "" {
		user.Image = resp.SecureURL
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

func (h *handlerUser) DeleteUser(w http.ResponseWriter, r *http.Request){
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
		Address: u.Address,
	}
}