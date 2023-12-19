
# Spotify User Case API

This is a simple API that allows users to interact with Spotify track information. Users can perform actions such as retrieving track details by ISRC code, searching tracks by artist name, creating new tracks, and updating existing track information.

## Getting Started

### Prerequisites

Before running the application, make sure you have the following installed:

- Go (Golang)
- PostgreSQL

### Installation

1. Clone the repository:

   
   git clone https://github.com/lingesh-sk/spotify_usercase.git
   cd spotify_usercase
   Install dependencies:


2. Install dependencies:
      go mod download

3. Set up the PostgreSQL database:

      Create a new PostgreSQL database named spotifyusercasedb.
      Update the database connection details in the main.go file.

4. Run the application:
      go run main.go


### API Endpoints
    - GET /track/:isrc

      Retrieve track details by ISRC code.

      ![GetByISRC_SS](https://github.com/lingesh-sk/spotify_usercase/assets/119925929/8ffbaa09-b64d-4fe2-9ad7-2453cab804c4)

    - GET /track/artist/:artistName

      Search tracks by artist name.

      ![GetByArtistName_SS](https://github.com/lingesh-sk/spotify_usercase/assets/119925929/b7417b45-0474-4d83-b42d-6ab7e1780e66)

    - POST /track

      Create a new track.

      ![PostbyISRC_SS](https://github.com/lingesh-sk/spotify_usercase/assets/119925929/03332e49-cb7a-40be-a489-919796620650)


    - PUT /track/:isrc

      Update an existing track by ISRC code.

      ![PutbyISRC_SS](https://github.com/lingesh-sk/spotify_usercase/assets/119925929/2e628036-d165-4f73-bf71-8d278c01796c)


  
Data which is stored in Postgres Database
![DB_details_SS](https://github.com/lingesh-sk/spotify_usercase/assets/119925929/ff8c2c4d-7bd7-47dc-b136-5c378ed5e9f2)
### Swagger Documentation

Swagger documentation is available at http://localhost:8080/swagger/index.html.


### Configuration
Database and Spotify API credentials can be configured in the main.go file.


var SpotifyCredentials = struct {
	ClientID     string
	ClientSecret string
}{
	ClientID:     "your_spotify_client_id",
	ClientSecret: "your_spotify_client_secret",
}

// ...

db, err := gorm.Open("postgres", "postgres://your_username:your_password@localhost:5432/spotifyusercasedb?sslmode=disable")
