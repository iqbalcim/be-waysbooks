package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	booksdto "waysbooks/dto/books"
	dto "waysbooks/dto/result"
	"waysbooks/models"
	"waysbooks/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerBook struct {
	BookRepository repositories.BookRepository
}

func HandlerBook(BookRepository repositories.BookRepository) *handlerBook {
	return &handlerBook{BookRepository}
}

func (h *handlerBook) FindBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	books,err := h.BookRepository.FindBooks()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, p := range books {
		books[i].Thumbnail = os.Getenv("PATH_FILE") + p.Thumbnail
	}

	for i,p := range books {
		books[i].BookAttachment = os.Getenv("PATH_FILE") + p.BookAttachment
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: books}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerBook) GetBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	book, err := h.BookRepository.GetBook(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	book.Thumbnail = os.Getenv("PATH_FILE") + book.Thumbnail

	book.BookAttachment = os.Getenv("PATH_FILE") + book.BookAttachment

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: book}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBook) CreateBook (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
  	filename := dataContex.(string)

	dataPDF := r.Context().Value("dataPDF")
  	filePDF := dataPDF.(string)

	pages, _ := strconv.Atoi(r.FormValue("Pages"))
	price, _ := strconv.Atoi(r.FormValue("Price"))

	request := booksdto.BookRequest{
		Title:       			r.FormValue("Title"),
		PublicationDate:    	r.FormValue("PublicationDate"),
		Pages:    				pages,
		ISBN:    				r.FormValue("ISBN"),
		Author:    				r.FormValue("Author"),
		Price:    				price,
		Description:       		r.FormValue("Description"),
	}

	validation := validator.New()

	err := validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	book := models.Book{
		Title:       			request.Title,
		PublicationDate:    	request.PublicationDate,
		Pages:    				request.Pages,
		ISBN:    				request.ISBN,
		Author:   				request.Author,
		Price:    				request.Price,
		Description:       		request.Description,
		BookAttachment: 		filePDF,
		Thumbnail: 				filename,
	}
	

	book , _ = h.BookRepository.CreateBook(book)

	book, _ = h.BookRepository.GetBook(book.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: book}
	json.NewEncoder(w).Encode(response)

}