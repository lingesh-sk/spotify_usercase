package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lingesh-sk/spotify_usercase/dao"
	"github.com/lingesh-sk/spotify_usercase/docs"
	"github.com/lingesh-sk/spotify_usercase/model"
	"github.com/lingesh-sk/spotify_usercase/routes"
	"github.com/lingesh-sk/spotify_usercase/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

var SpotifyCredentials = struct {
	ClientID     string
	ClientSecret string
}{
	ClientID:     "eb5c76245429470098e63a4abf7ccc4e",
	ClientSecret: "2ce225e50c824469af729e464d63c0f3",
}

// @title Spotify API usercase
// @version 1.0
// @description A Golang application which interfaces to the Spotify API using GORM and the Gin web framework..
// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	//DB connection setup
	db, err := gorm.Open("postgres", "postgres://postgres:12345678@localhost:5432/spotifyusercasedb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&model.Track{})

	client := authenticateSpotify()
	if client == nil {
		log.Fatal("Failed to authenticate with Spotify API")
	}

	dbAccessor := dao.NewDatabaseAccessor(db)
	spotifyService := service.NewSpotifyService(client)
	trackService := service.NewTrackService(dbAccessor, spotifyService)

	// Swagger documentation setup
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	setupRoutes(router, trackService)

	router.Run(":8080")
}

func authenticateSpotify() *spotify.Client {
	config := &clientcredentials.Config{
		ClientID:     SpotifyCredentials.ClientID,
		ClientSecret: SpotifyCredentials.ClientSecret,
		TokenURL:     spotify.TokenURL,
	}

	token, err := config.Token(context.Background())
	if err != nil {
		log.Printf("Failed to get Spotify token: %v", err)
		return nil
	}

	client := spotify.Authenticator{}.NewClient(token)
	return &client
}

// It initializes all routes
func setupRoutes(router *gin.Engine, trackService *service.TrackService) {

	routes := routes.NewRoutes(trackService)

	routes.RegisterRoutes(router)
}
