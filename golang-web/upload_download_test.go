package golang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
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

func DownloadFile(w http.ResponseWriter, r *http.Request){
	fileName := r.URL.Query().Get("file")
	if fileName == ""{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "BAD REQUEST")
		return 
	}

	w.Header().Add("Content-Disposition", `attachment; filename="`+fileName+`"`)
	http.ServeFile(w, r, "./public/asset/" + fileName)
}

func TestUploadDownloadForm(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/form", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./public/asset"))))
	mux.HandleFunc("/download", DownloadFile)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}

//go:embed public/asset/000000.png
var uploadFileTest []byte

func TestUploadFile(t *testing.T){
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Dummy Upload User")
	file, _ := writer.CreateFormFile("file", "dummy_upload.png")
	file.Write(uploadFileTest)
	writer.Close()

	url := BASE_URL + "/upload"
	request := httptest.NewRequest(http.MethodPost, url, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)
	bodyRes, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyRes))
}