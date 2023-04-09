package main

import "fmt"

func main() {
	score := 99

	switch score {
	case 80:
		fmt.Println("exec case 80")
	case 85:
		fmt.Println("exec case 85")
	case 90,99:
		fmt.Println("exec case 90")
	default:
		fmt.Println("exec case default ...")
	}

	// grade := 100

	switch grade:= 100;  {
	case grade >= 90:
		fmt.Println("exec case 90")
	case grade >= 80 && grade <= 90:
		fmt.Println("exec case 80")
	default:
		fmt.Println("")

	}

}
