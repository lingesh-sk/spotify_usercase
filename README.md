
# Spotify User Case API

This is a simple API that allows users to interact with Spotify track information. Users can perform actions such as retrieving track details by ISRC code, searching tracks by artist name, creating new tracks, and updating existing track information.

## Getting Started

### Prerequisites

Before running the application, make sure you have the following installed:

- Go 
- PostgreSQL DB

### Installation

1. Clone the repository:
   ```bash
	git clone https://github.com/lingesh-sk/spotify_usercase.git
   ```
2. Install dependencies:
   ```bash
	go mod tidy
   ```
4. Set up the PostgreSQL database:

      Create a new PostgreSQL database named spotifyusercasedb.
      Update the database connection details in the main.go file.
   ```bash
   db, err := gorm.Open("postgres", "postgres://your_username:your_password@localhost:5432/spotifyusercasedb? 
   sslmode=disable")
   ```

### Configuration

Go to the Spotify Developer Dashboard (https://developer.spotify.com/dashboard/applications).
Replace Spotify application credentials in main.go with your own:
```bash
{
	ClientID:     "your_spotify_client_id",
	ClientSecret: "your_spotify_client_secret",
}
```

5. Run the application:
   
    To generate the swagger docs 
   ```bash
	swag init
   ```
    To start the server
   ```bash
	go run main.go
   ```
   
### API Endpoints

 ```GET /track/:isrc```
 Retrieve track details by ISRC code.

 ```GET /track/artist/:artistName```
 Search tracks by artist name.

 ```POST /track```
 Create a new track.

 ```PUT /track/:isrc```
 Update an existing track by ISRC code.


### Swagger Documentation

#### Swagger documentation is available at http://localhost:8080/swagger/index.html.


## Output Screenshots

### Swagger homepage

![swagger_homepage_SS](https://github.com/lingesh-sk/spotify_usercase/assets/119925929/892357ef-c288-421b-bbd8-13d30cc6a0ae)

### Retrieve track details by ISRC code 

 ![GetByISRC_SS](https://github.com/lingesh-sk/spotify_usercase/assets/119925929/8ffbaa09-b64d-4fe2-9ad7-2453cab804c4)

### Search tracks by artist name.
 ![GetByArtistName_SS](https://github.com/lingesh-sk/spotify_usercase/assets/119925929/b7417b45-0474-4d83-b42d-6ab7e1780e66)

### Create a new track.

![PostbyISRC_SS](https://github.com/lingesh-sk/spotify_usercase/assets/119925929/03332e49-cb7a-40be-a489-919796620650)

### Update an existing track by ISRC code.

![PutbyISRC_SS](https://github.com/lingesh-sk/spotify_usercase/assets/119925929/2e628036-d165-4f73-bf71-8d278c01796c)

### Data stored in the PostgreSQL database

![DB_details_SS](https://github.com/lingesh-sk/spotify_usercase/assets/119925929/ff8c2c4d-7bd7-47dc-b136-5c378ed5e9f2)
