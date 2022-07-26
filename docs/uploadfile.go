package docs

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20)

	file, fileHeader, err := r.FormFile("file")

	fmt.Println(fileHeader.Filename)
	path := r.FormValue("path")

	if err != nil {
		fmt.Println("File upload error")
	}

	defer file.Close()
	dst, err := os.Create(fmt.Sprintf("./disk%s%d%s", path, time.Now().Day(), fileHeader.Filename))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Println("1Error Server")
		return
	}

	fmt.Println("Upload successful \t\t remote Address [", r.RemoteAddr, "]")

}
