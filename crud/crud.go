package crud

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Note struct {
	Title       string
	Description string
	Status      Status
	CreatedAt   time.Time
	ChangeTime  time.Time
}

type Status string
type Notes []Note

const (
	Todo       Status = "TODO"
	InProgress Status = "In Progress"
	Finished   Status = "Finished"
)

// TODO: Create
func (notes *Notes) CreateNote(title, description string) error {
	note := Note{
		Title:       title,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
		ChangeTime:  time.Now(),
	}

	*notes = append(*notes, note)
	return nil
}

// TODO: Delete
func (notes *Notes) isValidate(index int) error {
	if index < 0 || index >= len(*notes) {
		fmt.Println("Invalid Index")
	}

	return nil
}

// TOOD: Update
func (notes *Notes) DeleteNote(index int) error {
	n := *notes

	if err := n.isValidate(index); err != nil {
		return err
	}

	*notes = append(n[:index], n[index+1:]...)
	return nil
}

func (notes *Notes) ToggleStatus(index int) error {
	n := *notes

	if err := n.isValidate(index); err != nil {
		return err
	}

	currentStatus := n[index].Status

	switch currentStatus {
	case Todo:
		n[index].Status = Todo
	case InProgress:
		n[index].Status = InProgress
	case Finished:
		n[index].Status = Finished

	}

	return nil
}

func (notes *Notes) RenameTitle(index int, title string) error {
	n := *notes

	if err := n.isValidate(index); err != nil {
		return err
	}

	n[index].Title = title

	return nil
}

func (notes *Notes) print() {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"#", "Title", "Description", "Status", "Created As", "Completed"})
	for i, note := range *notes {
		createdAt := note.CreatedAt.Format("2006-01-02 15:04")
		completed := ""
		// Логика для определения завершенности можно изменить при необходимости
		if note.ChangeTime.After(note.CreatedAt) {
			completed = "Yes"
		} else {
			completed = "No"
		}
		table.Append([]string{
			fmt.Sprintf("%d", i+1),
			note.Title,
			note.Description,
			string(note.Status),
			createdAt,
			completed,
		})
	}
	table.Render()
}
