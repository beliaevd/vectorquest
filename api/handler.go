package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {

	conn, err := ApiUpload()
	if err != nil {
		fmt.Println("Conn error")
	}
	fmt.Println("Hi")

	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		conn.Write([]byte("-1"))
		fmt.Println("No File")
		return
	}

	conn.Write([]byte("send " + fileHeader.Filename))
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	_, err = conn.Write(fileBytes)
	if err != nil {
		fmt.Println("Lose connection")
	}
	defer file.Close()

	defer conn.Close()
}
