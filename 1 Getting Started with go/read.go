package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname, lname string
}

// This func reads a file and returns a string containing the content of the file
func FileToString() string {
	var (
		FileName, str_dat string
	)
	fmt.Println("Enter a file name: ")
	fmt.Scan(&FileName)
	dat, err := os.ReadFile(FileName)
	if err != nil {
		log.Fatal(err)
	}
	str_dat = string(dat)

	return str_dat
}

// This func takes a string input and returns a slice of structs of the type Name
func SrtToStruct(str_input string) []Name {
	names := strings.Split(str_input, "\n")
	x := len(names) - 1
	fmt.Printf("A total of %d names will be added to the Struct: \n", x)
	var struct_of_Names []Name
	for i := 0; i < x; i++ {
		FNandLN := strings.Split(names[i], " ")
		// adjust the len to the first 20 chars
		switch {
		case len(FNandLN[0]) > 20 && len(FNandLN[1]) > 20:
			strct := Name{FNandLN[0][:20], FNandLN[1][:20]}
			struct_of_Names = append(struct_of_Names, strct)
		case len(FNandLN[0]) > 20:
			strct := Name{FNandLN[0][:20], FNandLN[1]}
			struct_of_Names = append(struct_of_Names, strct)
		case len(FNandLN[1]) > 20:
			strct := Name{FNandLN[0], FNandLN[1][:20]}
			struct_of_Names = append(struct_of_Names, strct)
		default:
			strct := Name{FNandLN[0], FNandLN[1]}
			struct_of_Names = append(struct_of_Names, strct)
		}

	}
	return struct_of_Names
}

func main() {
	final_struct := SrtToStruct(FileToString())
	fmt.Println("The names obtained are: ")
	for i := 0; i < len(final_struct); i++ {
		name_item := final_struct[i]
		fmt.Println(name_item.fname, name_item.lname)
	}
}
