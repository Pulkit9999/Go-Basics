package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello !! good to go with Go")

	// Take input from user
	fmt.Println("Please enter a rating for Pizza:")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	fmt.Println("Thanks for rating!!...Your rating is", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hello", numRating+1)
	}
	/////////////////////////// ARRAYS in Go/////////////////////////////////////////////

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Peach"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	////////////////////////// SLICES in Go/////////////////////////////////////////////////////////

	var fruitLists = []string{"Apple", "Orange", "Banana"}
	fmt.Println(fruitLists)
	fruitLists = append(fruitLists, "Peach", "Guava")
	fmt.Println("new slice is:", fruitLists)

	//////////////////////////// MAPS  in  Go///////////////////////////////////////////////////////////

	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["rs"] = "React JS"
	languages["py"] = "Python"

	fmt.Println(languages)
	fmt.Println("Looping over map:languages")

	for key, value := range languages {
		fmt.Printf("For key %v , value is %v\n", key, value)

	}
	////////////////////////////////// STRUCTS in Go/////////////////////////////////////////////////

	pulkit := User{"Pulkit", "pullkit.p@gmail.com", 26}
	fmt.Println("Struct is:", pulkit)
	fmt.Printf("STRUCT is %+v", pulkit)

}

type User struct {
	Name  string
	Email string
	Age   int
}
