package auth_test

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

func TestLib_AddBook (t *testing.T) {

	requestBody, err := json.Marshal(&models.Book{
		Name:   "little prince10",
		Author: "lc",
		Genre:  "ch",
		Year:   10,
	})

	if err != nil {
		log.Error(err)
	}

	log.Println(string(requestBody))
	responseAddBook, err := http.Post( fmt.Sprintf("http://localhost:%d/addBook",1323),
		"application/json",
		bytes.NewBuffer(requestBody))
	if err != nil {
		log.Error(err)
	}
	defer responseAddBook.Body.Close()
	_, err = ioutil.ReadAll(responseAddBook.Body)
	//EndOfFile or error.
	if err != nil {
		log.Error(err)
	}
}
