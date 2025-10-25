package command

import (
	"firstProject/crud"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Delete int
	Rename string
	Toggle int
	List   bool
}

func NewCommandFlugs() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add new Notes with specifice title")
	flag.StringVar(&cf.Rename, "edit", "", "Edit a notes by index & specificy a new title")
	flag.IntVar(&cf.Delete, "del", -1, "Specify a notes by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a notes by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "List All notes")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(notes *crud.Notes) {
	switch {
	case cf.List:
		notes.Print()
	case cf.Add != "":
		notes.CreateNote(cf.Add, "No Description")
	case cf.Rename != "":
		parts := strings.SplitN(cf.Rename, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: invalid format for Edit. Please Use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error : invalid index for edit")
			os.Exit(1)
		}
		notes.RenameTitle(index, parts[1])

	case cf.Toggle != -1:
		notes.ToggleStatus(cf.Toggle)

	default:
		fmt.Println("Invalid Command")
	}
}
