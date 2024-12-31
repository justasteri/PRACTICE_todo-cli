package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

func DefineFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add a task. Format: title",
		},
		&cli.StringFlag{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "Update a task title. Format: index:new_title",
		},
		&cli.IntFlag{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "Delete a task. Format: index",
		},
		&cli.StringFlag{
			Name:    "change_status",
			Aliases: []string{"cs"},
			Usage:   "Change the status of a task. Status: 1-To do, 2-Doing, 3-Done. Format: index: status_index",
		},
		&cli.BoolFlag{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "List all the tasks",
		},
		&cli.IntFlag{
			Name:    "list_status",
			Aliases: []string{"ls-s"},
			Usage:   "List tasks by status. Status: 1-To do, 2-Doing, 3-Done",
		},
	}
}

func handleChangeStatus(todos *Todos, cmd_arg string) {
	parts := strings.Split(cmd_arg, ":")

	if len(parts) != 2 {
		fmt.Println("Invalid Format. Please use the format: index:status_index")
	}

	indexTask, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		fmt.Println("Index must be a valid integer")
	}

	indexStatus, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		fmt.Println("Status index must be a valid integer")
	}

	todos.changeStatus(indexTask, indexStatus)
}

func handleCommands(ctx *cli.Context, todos *Todos) error {
	title := ctx.String("add")
	update := ctx.String("update")
	indexDelete := ctx.Int("delete")
	change_status := ctx.String("change_status")
	list := ctx.Bool("list")
	list_status := ctx.Int("list_status")

	switch {
	case list:
		todos.print()
	case list_status != 0:
		todos.printByStatus(list_status)
	case title != "":
		todos.add(title)
	case update != "":
		parts := strings.Split(update, ":")

		if len(parts) != 2 {
			fmt.Println("Invalid Format. Please use the format: index:new_value")
		}

		index, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			fmt.Println("Index must be a valid integer")
		}

		newTitle := strings.TrimSpace(parts[1])
		todos.update(index, newTitle)
	case indexDelete != 0:
		todos.delete(indexDelete)
	case change_status != "":
		handleChangeStatus(todos, change_status)
	default:
		fmt.Println("Invalid command")
	}

	return nil
}

func SetupCli() *cli.App {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)

	return &cli.App{
		Name:  "Todo App",
		Usage: "Create to-do's in the terminal",
		Flags: DefineFlags(),
		Action: func(ctx *cli.Context) error {
			if err := handleCommands(ctx, &todos); err != nil {
				fmt.Println("Error: ", err)
				return err
			}

			storage.Save(todos)

			return nil
		},
	}
}
