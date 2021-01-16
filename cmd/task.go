package main

import (
	"fmt"
	"goworkshop/task"
	"os"
	"strconv"
)

func main() {

	db := task.Database{
		FilePath: "database.json",
	}

	args := os.Args
	input := args[1]
	switch input {
	case "add":
		Add(db, args)
		break
	case "list":
		List(db)
		break
	default:
		break
	}
}

func Add(db task.Database, args []string) {
	fmt.Println("ADD")

	age, err := strconv.Atoi(args[3])
	var item task.Person
	if err == nil {
		item = task.Person{
			Name: args[2],
			Age:  age,
		}
		err = db.Add(item)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("add ", item)
	}

}

func List(db task.Database) {
	fmt.Println("List")
	fmt.Println(db.List())
}
