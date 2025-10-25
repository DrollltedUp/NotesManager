package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	data := [][]string{
		{"1", "Buy Milk", "Today we buy milk", "TODO", "2023-11-01 10:00", "No"},
		{"2", "Buy Bread", "Today we buy bread", "In Progress", "2023-11-02 11:00", "No"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"#", "Title", "Description", "Status", "Created As", "Completed"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
