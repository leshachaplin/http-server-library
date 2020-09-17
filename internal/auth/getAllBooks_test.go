package auth

import (
	"encoding/json"
	"fmt"
	"github.com/leshachaplin/http-server-library/internal/models"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestServer_GetAllBooks(t *testing.T) {


	requestBody := make([]byte, 0)

	log.Println(requestBody)
	responseGetAllBooks, err := http.Get( fmt.Sprintf("http://localhost:%d/getAll",1323))
	if err != nil {
		log.Error(err)
	}
	defer responseGetAllBooks.Body.Close()
	bodyGetAll, err := ioutil.ReadAll(responseGetAllBooks.Body)
	//EndOfFile or error.
	if err != nil {
		log.Error(err)
	}
	var res models.Books
	err = json.Unmarshal(bodyGetAll,&res.B)
	log.Println(res.B[2].Name)
	log.Println(res.B[1].Author)

}


