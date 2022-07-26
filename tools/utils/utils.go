package utils

import (
	"Vectorquest/db"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func CryptSession(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetAdminCat(id int) []string {
	type Result struct {
		Name string
	}

	c := []Result{}

	err := db.GetDB().Raw("SELECT c.name FROM admin_user_categories c JOIN admin_category_links l ON c.id = l.category_id WHERE l.user_id = ?", id).Find(&c).Error
	if err != nil {
		fmt.Println(err)
	}

	var resp []string

	for i := 0; i < len(c); i++ {
		resp = append(resp, c[i].Name)
	}

	return resp
}
