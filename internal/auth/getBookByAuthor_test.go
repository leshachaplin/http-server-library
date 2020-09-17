package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/leshachaplin/http-server-library/internal/models"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestServer_GetBookByAuthor(t *testing.T) {

	author := fmt.Sprint("leshiy1")

	requestBody, err := json.Marshal(map[string]string {
		"Author": author,
	})

	if err != nil {
		log.Error(err)
	}

	log.Println(string(requestBody))
	responseGetBookByAuthor, err := http.Post( fmt.Sprintf("http://localhost:%d/getByAuthor",1323),
		"application/json",
		bytes.NewBuffer(requestBody))
	if err != nil {
		log.Error(err)
	}
	defer responseGetBookByAuthor.Body.Close()
	bodyGetByAuthor, err := ioutil.ReadAll(responseGetBookByAuthor.Body)
	//EndOfFile or error.
	if err != nil {
		log.Error(err)
	}
	var res models.Book
	err = json.Unmarshal(bodyGetByAuthor,&res)
	log.Println(res.Name)
	log.Println(res.Year)
}

