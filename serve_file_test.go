package golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/ok.html")
	} else {
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Handler: http.HandlerFunc(ServeFile),
		Addr:    ":8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourcesOk string

//go:embed resources/ok.html
var resourcesNotFound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourcesOk)
	} else {
		fmt.Fprint(writer, resourcesNotFound)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Handler: http.HandlerFunc(ServeFileEmbed),
		Addr:    ":8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
