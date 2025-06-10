package main 

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("please prov ide a command line argument")
		return
	}
	argument := os.Args[1]
	 switch argument{
	 case "0":
		fmt.Println("zero")
	 case "1":
		fmt.Println("one!")
	 case "2", "3", "4":
		fmt.Println("2 or 3 or 4")
		fallthrough
	 default:
		fmt.Println("Value : ",argument)
	 }

	 value, err := strconv.Atoi(argument)
	 if err != nil {
		fmt.Println("Cannot convert to int ", argument)
		return 
	 }

	 switch {
	 case value == 0:
		fmt.Println("Zero!")
	 case value > 0:
		fmt.Println("Positive integer")
	 case value < 0:
		fmt.Println("Negative interger")
	 default:
		fmt.Println("This should no Happen ",value)
	 }
}