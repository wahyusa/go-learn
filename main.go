package main

import (
	"go-learn/database"
)

func main() {
	db := database.ConnectionStart()

	wtf := &controllers.store()
}
