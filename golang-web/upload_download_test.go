package golang_web

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request){
	err := myTemplates.ExecuteTemplate(w, "upload", nil)
	if err != nil{
		panic(err)
	}
}

func Upload(w http.ResponseWriter, r *http.Request){
	file, fileHeader, err := r.FormFile("file")
	if err != nil{
		panic(err)
	}
	fileDestination, err := os.Create("./public/asset/" + fileHeader.Filename)
	if err != nil{
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil{
		panic(err)
	}
	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload.success", map[string]interface{}{
		"Name" : name,
		"File":"/static/" + fileHeader.Filename,
	})
}

func TestUploadDownloadForm(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/form", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./public/asset"))))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}