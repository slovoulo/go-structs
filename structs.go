package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"time"
)

var reader = bufio.NewReader(os.Stdin)

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
	Price       float64
}

//2)Create concrete instances of this data type in 2 different ways:
// -Directly inside main
// -Inside of main, by using a "Creation helper function"
// Output created DS in cmd

func NewVideoGame(gameId string, gametitle string, gamedesc string, gamePrice float64) *VideoGame {

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

func getUserInput()*VideoGame{
	//Prompt user for struct values
	fmt.Println("Enter game details ")
	fmt.Println("-----------------------")
	fmt.Println("Enter new game ID: ")
	
	gameId, _ := reader.ReadString('\n') //read all user's input values until enter(\n) is pressed
	fmt.Print("Enter new game Title: ")
	gameTitle, _ := reader.ReadString('\n') 
	fmt.Print("Enter new game Decription: ")
	gameDesc, _ := reader.ReadString('\n') 
	fmt.Print("Enter new game Price: ")
	gamePrice, _ := reader.ReadString('\n') 

	//save user input in variables
	gameId = strings.TrimSpace(gameId)
	gamePrice = strings.TrimSpace(gamePrice)
	// gameTitle=gameTitle
	// gameDesc=gameDesc

	//Convert price string to float
	urgamePrice,_ :=strconv.ParseFloat(gamePrice,64)  //64 represents the float type 

	newUserGame:=NewVideoGame(gameId,gameTitle,gameDesc,urgamePrice)

	fmt.Println(newUserGame)

	return newUserGame
	
}



//4)Bonus: Add a connected "Store" function that writes that data into a file.
//	 The file name should be the unique ID, the function should be called at the end of main

//create a method pointing to VideoGame struct
func(writeGame *VideoGame)Store(){
	//use the os package to create a file with videogame ID as its name
	file,_ :=os.Create(writeGame.ID)
	//Create a formatted string using the Sprintf function
	content := fmt.Sprintf("ID: %v\nTitle: %v\nDescription %v\nPrice: %v\n", 
writeGame.ID,
writeGame.Title,
writeGame.Description,
writeGame.Price)

//Pass the formatted content to file
file.WriteString(content)

//Save the file by
file.Close()


}

func main() {
	

	//Exercise 2 solution
	gow := VideoGame{
		ID:          "1",
		Title:       "God of war",
		Description: "My vengeance ends now",
		Price:       32.89,
	}
	fmt.Print(gow)

	fmt.Println(NewVideoGame("2", "The last of us", "Survive zombie virus", 29.99))
	//getUserInput()
	getUserInput().Store()

}
