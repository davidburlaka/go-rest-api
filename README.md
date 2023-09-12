# Go REST API with Gin and PostgreSQL

This is a simple Go REST API built using the Gin framework and PostgreSQL to manage players and their scores. The API provides several endpoints to perform CRUD operations on player information and retrieve a leaderboard.

## Prerequisites 

1. Installed PostgreSQL
2. Installed Golang

## Local Installation and Setup
### 1. Install Dependencies

Use following command to install the required dependencies:
```go mod download```

### 2. Database Configuration

Create a PostgreSQL database called "gamedatabase". Add `players` table with columns `player_id`, `player_name` and `score`.

### 3. Run the API

To start the API, run the following command:
```go run main.go```

The API should now be accessible at `http://localhost:8080`.

## API Endpoints

The API contains the following endpoints:

### 1. Get Player Information
  Endpoint: ```/player```

  Method: GET
  
  Description: Retrieve information about all players.
  
### 2. Add Player
  Endpoint: ```/player```
  
  Method: POST
  
  Description: Add a new player with score information.
  
### 3. Delete Player
  Endpoint: ```/player/{id}```
  
  Method: DELETE
  
  Description: Delete a player by their ID.
  
### 4. Update Player Information
  Endpoint: ```/player/{id}```
  
  Method: POST
  
  Description: Update player information: name, score or both.
  
### 5. Get Leaderboard
  Endpoint: ```/leaderboard```
  
  Method: GET
  
  Description: Retrieve the player leaderboard sorted by score.
