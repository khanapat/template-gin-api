package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"template-gin-api/config"
	"template-gin-api/internal/database"
	"template-gin-api/internal/logz"
	"time"

	_ "time/tzdata"

	"github.com/gin-gonic/gin"
)

func init() {
	runtime.GOMAXPROCS(1)
	initTimezone()
}

func initTimezone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Printf("error loading location 'Asia/Bangkok': %v\n", err)
	}
	time.Local = ict
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Printf("failed to load config: %v\n", err)
	}

	logger, err := logz.NewLogConfig(cfg)
	if err != nil {
		log.Printf("failed to init log: %v\n", err)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Printf("failed to sync log: %v\n", err)
		}
	}()

	postgresDB, err := database.NewPostgresConn(cfg)
	if err != nil {
		logger.Error(err.Error())
	}
	defer postgresDB.Close()

	ctx, cancel := context.WithTimeout(context.Background(), cfg.App.Timeout)
	defer cancel()

	if err := postgresDB.Ping(ctx); err != nil {
		logger.Error(err.Error())
	}

	router := gin.Default()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.App.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-c
	logger.Info("Shutdown Server...")

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Server exiting")
	os.Exit(0)
}

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "ping",
// 		})
// 	})

// 	r.GET("/albums", getAlbums)
// 	r.GET("/albums/:id", getAlbumByID)
// 	r.POST("/albums", postAlbums)

// 	r.Run(":8080")
// }

// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// func getAlbumByID(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, a := range albums {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }

// func postAlbums(c *gin.Context) {
// 	var newAlbum album
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }
