package tests

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed public
var public embed.FS

type LogMiddleware struct{
	Handler http.Handler
}

type ErrorMiddleware struct{
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request){
	fmt.Println("Start Log")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("Finish Log")
}

func (middleware *ErrorMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request){
	defer func (){
		message := recover()
		if message != nil{
			fmt.Printf("Recovered Error: %s \n", message)
		}
	}()
	middleware.Handler.ServeHTTP(w, r)
}

func TestHttpRouter(t *testing.T){
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
		fmt.Fprint(w, "Hello World")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Hello World", string(bytes), "Result Must be Equal to Hello World")
}

func TestHttpRouterParams(t *testing.T){
	router := httprouter.New()
	router.GET("/product/:id", func(w http.ResponseWriter, r *http.Request, params httprouter.Params){
		fmt.Fprintf(w, "Product %s", params.ByName("id"))
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/product/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1", string(bytes), "Product 1")
}

func TestHttpRouterCatchAllParams(t *testing.T){
	router := httprouter.New()
	router.GET("/image/*image", func(w http.ResponseWriter, r *http.Request, params httprouter.Params){
		fmt.Fprintf(w, params.ByName("image"))
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/image/small/test.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "/small/test.png", string(bytes), "/small/test.png")
}

func TestServeFile(t *testing.T){
	router := httprouter.New()
	directory, _ := fs.Sub(public, "public")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/files/test.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Hello Text Embed", string(body))
}

func TestHttpRouterAll(t *testing.T){
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
		fmt.Fprint(w, "Middleware")
	})
	
	logMD := LogMiddleware{router}
	errMD := ErrorMiddleware{&logMD}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/", nil)
	recorder := httptest.NewRecorder()

	errMD.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Middleware", string(body))
}