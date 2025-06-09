package main

import (
	"fmt"
	"search-person/utils"
)

func main(){
	users := []string{
		"Lenanne Graham",
		"Ervin Howell",
		"Clementine Bauch",
		"Chelsey Dietrich",
		"Mrs. Dennis Schulist",
		"Kurtis Weissnat",
		"Nicholas Runolfsdottir",
		"Glena Reichert",
		"Clementina DuBuque",
	}

	result := utils.SearchPerson(users ,"Lenanne")

	fmt.Println(result)
}

