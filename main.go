package main

import (
	"Glib/io"
	"fmt"
	"net/http"
	"os"
)
import "github.com/julienschmidt/httprouter"

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":9999", r)
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) ()  {
	wd,_   := os.Getwd()
	upload := io.NewFileUpload("file", wd)
	fmt.Println(upload.UploadHandler(w, r, p))
	//fmt.Fprint(w, "success")
}

//RegisterHandlers ...
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/upload", uploadHandler)
	return router
}



