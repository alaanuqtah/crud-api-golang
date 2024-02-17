package main

import (
	"crud-api-golang/initializers"
	"crud-api-golang/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
