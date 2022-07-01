package main

import (
	"fmt"

	"github.com/tschneiderAK/go_webservice/models"
)

func main() {
	u := models.User{
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)
}
