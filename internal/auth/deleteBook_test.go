package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestServer_DeleteBook(t *testing.T) {

	name := fmt.Sprint("little prince1")

	requestBody, err := json.Marshal(map[string]string {
		"Name": name,
	})

	if err != nil {
		log.Error(err)
	}

	log.Println(string(requestBody))
	responseDeleteBook, err := http.Post( fmt.Sprintf("http://localhost:%d/deleteBook",1323),
		"application/json",
		bytes.NewBuffer(requestBody))
	if err != nil {
		log.Error(err)
	}
	defer responseDeleteBook.Body.Close()
	_, err = ioutil.ReadAll(responseDeleteBook.Body)
	//EndOfFile or error.
	if err != nil {
		log.Error(err)
	}
}

