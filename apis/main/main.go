package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/rept0id/birthday-planetary-backend/apis/main/model"

	db_stars_connection "github.com/rept0id/birthday-planetary-backend/apis/main/dbs/stars/connection"

	"github.com/gin-gonic/gin"

	routes "github.com/rept0id/birthday-planetary-backend/apis/main/routes/stars"
)

func main() {
	/*** * DBs * ***/

	db_stars_conn, err := db_stars_connection.Connection()
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db_stars_conn.Close()

	dbs := model.DBs{
		DBStarsConn: db_stars_conn,
	}

	/*** * Router * ***/

	// Instance

	router := gin.Default()

	router.Any("/", func(c *gin.Context) {
		c.String(200, "Welcome to Birthday Planetary Main API !")
	})

	// Routes

	routes.Routes(dbs, router)

	// Port

	port := os.Getenv("API_MAIN_PORT")
	if port == "" {
		port = "8080"
	}

	// Run

	log.Printf("Listening on port %s", port)

	router.Run(":" + port)
}
