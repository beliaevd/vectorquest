package docs

import (
	"fmt"
	"net/http"
	"os"
)

func MakeDir(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("dir")

	os.Mkdir(fmt.Sprintf("./disk/%s", name), 0777)
	fmt.Println("Ok")

}
