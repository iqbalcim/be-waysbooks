package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func UploadPdf(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("attachment")

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode("Error Retrieving the File")
			return
		}
		defer file.Close()

		const MAX_UPLOAD_SIXE = 200 << 20

		r.ParseMultipartForm(MAX_UPLOAD_SIXE)
		if r.ContentLength > MAX_UPLOAD_SIXE {
			w.WriteHeader(http.StatusBadRequest)
			response := Result{Code: http.StatusBadRequest, Message: "Max size in 20mb"}
			json.NewEncoder(w).Encode(response)
			return
		}
		tempFile, err := ioutil.TempFile("uploads", "book-*.pdf")
		if err != nil {
			fmt.Println(err)
			fmt.Println("path upload error")
			json.NewEncoder(w).Encode(err)
			return
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		tempFile.Write(fileBytes)

		data := tempFile.Name()
		filepdf := data[8:]

		ctx := context.WithValue(r.Context(), "dataPDF", filepdf)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}