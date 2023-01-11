package main

import (
	"fmt"
	//"time"
)

//Define a struct
// type structName struct{
//     Data int
//     AnotherVar string
//      Createddate time.Time
// }

//var structVariable structName

// Struct practise
// 1) Create a new type/ data structure for storing product data (videogame price ) The DS should contain these fields
// ID
// Title
// Short description
// Price without currency
type VideoGame struct {
	ID          string
	Title       string
	Description string
	Price       float32
}

//2)Create concrete instances of this data type in 2 different ways:
// -Directly inside main
// -Inside of main, by using a "Creation helper function"
// Output created DS in cmd

func NewVideoGame(gameId string, gametitle string, gamedesc string, gamePrice float32) *VideoGame {

	newGame := VideoGame{
		ID:          gameId,
		Title:       gametitle,
		Description: gamedesc,
		Price:       gamePrice,
	}
	return &newGame

}

//3)Change the program to fetch user input values for the different data fields and create only one concrete instance
//	with that data. Output that instance data

//4)Bonus: Add a connected "Store" function that writes that data into a file.
//	 The file name should be the unique ID, the function should be called at the end of main

func main() {
	// //create an instance of the struct
	// structVariable=structName{
	// 	Data: 1,
	// 	AnotherVar:"var another",
	// 	Createddate:time.Now(),
	// }
	// fmt.Println(structVariable)

	//Exercise 2 solution
	gow := VideoGame{
		ID:          "1",
		Title:       "God of war",
		Description: "My vengeance ends now",
		Price:       32.89,
	}
	fmt.Print(gow)

	fmt.Println(NewVideoGame("2", "The last of us", "Survive zombie virus", 29.99))

}
