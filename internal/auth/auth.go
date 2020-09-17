package auth

import (
	"context"
	"github.com/labstack/echo"
	"github.com/leshachaplin/grpc-server-library/protocol"
	"github.com/leshachaplin/http-server-library/internal/models"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Lib struct {
	client protocol.BookServiceClient
}

func NewHandler(client protocol.BookServiceClient) *Lib {
	return &Lib{client: client}
}

func (a *Lib) AddBook(c echo.Context) error {
	log.Info("AddBook")
	var book models.Book
	if err := c.Bind(&book); err != nil {
		log.Errorf("echo.Context Error AddBook %s", err)
		return err
	}
	_, err := a.client.AddBook(context.Background(), &protocol.AddBookRequest{
		Book: &protocol.Book{
			Name:   book.Name,
			Author: book.Author,
			Genre:  book.Genre,
			Year:   book.Year,
		},
	})
	if err != nil {
		log.Errorf("GRPC Error AddClaims %s", err)
		return err
	}
	return c.JSON(http.StatusOK, "")
}

func (a *Lib) DeleteBook(c echo.Context) error {
	log.Info("DeleteBook")
	var book models.Book
	if err := c.Bind(&book); err != nil {
		log.Errorf("echo.Context Error DeleteBook %s", err)
		return err
	}
	_, err := a.client.DeleteBook(context.Background(), &protocol.DeleteBookRequest{
		Name: book.Name,
	})
	if err != nil {
		log.Errorf("GRPC Error DeleteBook %s", err)
		return err
	}
	return c.JSON(http.StatusOK, "")
}

func (a *Lib) GetBookByName(c echo.Context) error {
	log.Info("GetBookByName")
	var book models.Book
	if err := c.Bind(&book); err != nil {
		log.Errorf("echo.Context Error GetBookByName %s", err)
		return err
	}

	res, err := a.client.GetBookByName(context.Background(), &protocol.GetBookByNameRequest{Name: book.Name})
	if err != nil {
		log.Errorf("GRPC Error GetBookByName %s", err)
		return echo.ErrBadRequest

	}
	return c.JSON(http.StatusOK, &models.Book{
		Name:   res.Book.Name,
		Author: res.Book.Author,
		Genre:  res.Book.Genre,
		Year:   res.Book.Year,
	})
}

func (a *Lib) GetBookByAuthor(c echo.Context) error {
	log.Info("GetBookByAuthor")
	var book models.Book
	if err := c.Bind(&book); err != nil {
		log.Errorf("echo.Context Error GetBookByAuthor %s", err)
		return err
	}

	res, err := a.client.GetBookByAuthor(context.Background(), &protocol.GetBooksByAuthorRequest{Author: book.Author})
	if err != nil {
		log.Errorf("GRPC Error GetBookByAuthor %s", err)
		return echo.ErrBadRequest

	}
	return c.JSON(http.StatusOK, &models.Book{
		Name:   res.Book.Name,
		Author: res.Book.Author,
		Genre:  res.Book.Genre,
		Year:   res.Book.Year,
	})
}

func (a *Lib) GeAlltBooks(c echo.Context) error {
	log.Info("GeAlltBooks")
	var book models.Book
	if err := c.Bind(&book); err != nil {
		log.Errorf("echo.Context Error GeAlltBooks %s", err)
		return err
	}

	res, err := a.client.GetAllBooks(context.Background(), &protocol.EmptyRequest{})
	if err != nil {
		log.Errorf("GRPC Error GeAlltBooks %s", err)
		return echo.ErrBadRequest
	}
	
	result := models.Books{}
	books := make([]models.Book, len(res.Books))
	
	for i := 0; i < len(res.Books); i++ {
		books[i].Author = res.Books[i].Author
		books[i].Genre = res.Books[i].Genre
		books[i].Name = res.Books[i].Name
		books[i].Year = res.Books[i].Year
	}
	result.B = books
	
	return c.JSON(http.StatusOK, books)
}