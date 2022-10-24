package server

import (
	"fmt"
	"log"
	"net/http"

	"mytodo/golangtodo/internal"
	"mytodo/golangtodo/model"

	"github.com/gin-gonic/gin"
)

type Json map[string]any

type CreateRequest struct {
	ListName  string `json:"list_name"`
	Overwrite bool   `json:"overwrite"`
}

type TodoServer struct {
	*gin.Engine
	app *internal.App
}

func (s *TodoServer) GetTableNames() ([]string, error) {
	return s.app.GetAllListNames()
}

type Entry struct {
	Content string `json:"content"`
	Status  bool   `json:"checked"`
}

type GetTableResponse struct {
	Title   string  `json:"title"`
	Entries []Entry `json:"entries"`
}

func (s *TodoServer) GetTable(name string) (*GetTableResponse, error) {
	list, err := s.app.GetList(name)
	if err != nil {
		return nil, err
	}
	entries := make([]Entry, len(list.Items))
	for _, item := range list.Items {
		entries = append(entries, Entry{
			Content: item.Content,
			Status:  item.ItemStatus == model.COMPLETE,
		})
	}
	return &GetTableResponse{
		Title:   list.ListName,
		Entries: entries,
	}, nil
}

func NewTodoServer() *TodoServer {
	r := gin.Default()
	app, err := internal.NewApp()
	if err != nil {
		panic(err)
	}
	out := &TodoServer{
		Engine: r,
		app:    app,
	}
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ListName": "jomama",
			"Items": []string{
				"Hi threr",
				"bobby hill hahahahaha",
				"who has my pie",
			},
		})
	})
	r.OPTIONS("/get_table_names", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	})
	r.GET("/get_table_names", func(c *gin.Context) {
		log.Println("Getting table names!")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		names, err := out.GetTableNames()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, gin.H{
			"tableNames": names,
		})
	})

	r.OPTIONS("/get_table", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	})
	r.GET("/get_table", func(c *gin.Context) {
		log.Println("Getting table names!")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		name := c.Query("tableName")
		if name == "" {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no table name provided"))
		}
		response, err := out.GetTable(name)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, response)
	})
	return out
}
