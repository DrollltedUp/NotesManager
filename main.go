package main

import (
	"firstProject/crud"
	"firstProject/storage"
	"fmt"
)

func main() {
	var notes crud.Notes
	storage := storage.NewStorage[crud.Notes]("Notes.json")
	if err := storage.Load(&notes); err != nil {
		fmt.Println("Error loading notes:", err)
	}

	notes.CreateNote("Buy Milk", "Today we buy milk")
	notes.CreateNote("Buy Bread", "Today we buy Bread")

	if err := notes.ToggleStatus(0); err != nil {
		fmt.Println("Error toggling status:", err)
	}

	if err := notes.DeleteNote(0); err != nil {
		fmt.Println("Error deleting note:", err)
	}

	notes.Print()

	if err := storage.Save(notes); err != nil {
		fmt.Println("Error saving notes:", err)
	}
}

// data := [][]string{
// 	{"1", "Buy Milk", "Today we buy milk", "TODO", "2023-11-01 10:00", "No"},
// 	{"2", "Buy Bread", "Today we buy bread", "In Progress", "2023-11-02 11:00", "No"},
// }

// storage := storage.NewStorage[crud.Notes]("Notes.json")
// storage.Load(&)

// table := tablewriter.NewWriter(os.Stdout)
// table.Header([]string{"#", "Title", "Description", "Status", "Created As", "Completed"})

// for _, v := range data {
// 	table.Append(v)
// }
// table.Render()
