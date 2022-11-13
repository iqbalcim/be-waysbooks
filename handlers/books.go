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
		books[i].Thumbnail =  p.Thumbnail
	}

	for i,p := range books {
		books[i].BookAttachment =  p.BookAttachment
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


	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: book}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBook) CreateBook (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
  	filename := dataContex.(string)

	filename = os.Getenv("PATH_FILE") + filename

	dataPDF := r.Context().Value("dataPDF")
  	filePDF := dataPDF.(string)

	filePDF = os.Getenv("PATH_FILE") + filePDF

	pages, _ := strconv.Atoi(r.FormValue("pages"))
	price, _ := strconv.Atoi(r.FormValue("price"))

	request := booksdto.BookRequest{
		Title:       			r.FormValue("title"),
		PublicationDate:    	r.FormValue("publication_date"),
		Pages:    				pages,
		ISBN:    				r.FormValue("ISBN"),
		Author:    				r.FormValue("author"),
		Price:    				price,
		Description:       		r.FormValue("description"),
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

func (h *handlerBook) UpdateBook (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	dataContex := r.Context().Value("dataFile")
  	filename := dataContex.(string)

	filename = os.Getenv("PATH_FILE") + filename

	dataPDF := r.Context().Value("dataPDF")
  	filePDF := dataPDF.(string)

	filePDF = os.Getenv("PATH_FILE") + filePDF

	pages, _ := strconv.Atoi(r.FormValue("pages"))
	price, _ := strconv.Atoi(r.FormValue("price"))

	request := booksdto.BookRequest{
		Title:       			r.FormValue("title"),
		PublicationDate:    	r.FormValue("publication_date"),
		Pages:    				pages,
		ISBN:    				r.FormValue("ISBN"),
		Author:    				r.FormValue("author"),
		Price:    				price,
		Description:       		r.FormValue("description"),
	}

	validation := validator.New()

	err := validation.Struct(request)
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	book := models.Book{}

	if request.Title != "" {
		book.Title = request.Title
	}

	if request.PublicationDate != "" {
		book.PublicationDate = request.PublicationDate
	}

	if request.Pages != 0 {
		book.Pages = request.Pages
	}

	if request.ISBN != "" {
		book.ISBN = request.ISBN
	}

	if request.Author != "" {
		book.Author = request.Author
	}

	if request.Price != 0 {
		book.Price = request.Price
	}

	if request.Description != "" {
		book.Description = request.Description
	}

	if filename != "" {
		book.Thumbnail = filename
	}

	if filePDF != "" {
		book.BookAttachment = filePDF
	}

	book , _ = h.BookRepository.UpdateBook(book, id)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: book}
	json.NewEncoder(w).Encode(response)
	
}

func (h *handlerBook) DeleteBook (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	book, err := h.BookRepository.GetBook(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.BookRepository.DeleteBook(book,id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}

