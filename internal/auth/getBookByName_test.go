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

func TestServer_GetBookByName(t *testing.T) {

	name := fmt.Sprint("little prince10")

	requestBody, err := json.Marshal(map[string]string {
		"Name": name,
	})

	if err != nil {
		log.Error(err)
	}

	log.Println(string(requestBody))
	responseGetBookByName, err := http.Post( fmt.Sprintf("http://localhost:%d/getByName",1323),
		"application/json",
		bytes.NewBuffer(requestBody))
	if err != nil {
		log.Error(err)
	}
	defer responseGetBookByName.Body.Close()
	bodyGetByName, err := ioutil.ReadAll(responseGetBookByName.Body)
	//EndOfFile or error.
	if err != nil {
		log.Error(err)
	}
	var res models.Book
	err = json.Unmarshal(bodyGetByName,&res)
	log.Println(res.Author)
	log.Println(res.Year)
}


