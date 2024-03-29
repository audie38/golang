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

type LogMiddleware struct{
	Handler http.Handler
}

type ErrorMiddleware struct{
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request){
	fmt.Println("Before execute handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After execute handler")
}

func (middleware *ErrorMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request){
	defer func(){
		err := recover()
		fmt.Println("Recover: ", err)
		if err != nil{
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer,"Error: %s", err)
		}
	}()

	middleware.Handler.ServeHTTP(writer, request)
}

func TestUploadDownloadForm(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/",	func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Log Middleware example")
	})
	mux.HandleFunc("/panic",	func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Error Middleware example ")
		panic("Error Middleware")
	})
	mux.HandleFunc("/form", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./public/asset"))))
	mux.HandleFunc("/download", DownloadFile)

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorMiddlware := &ErrorMiddleware{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: errorMiddlware,
	}

	err := server.ListenAndServe()
	if err != nil{
		t.Fatal(err)
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