package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Status      int
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Status:      1,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func (todos *Todos) changeStatus(index int, status int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	if status > 3 || status < 1 {
		return errors.New("invalid status index")
	}

	if status == 3 {
		completedTime := time.Now()
		t[index].CompletedAt = &completedTime
	} else {
		t[index].CompletedAt = nil
	}

	t[index].Status = status

	return nil
}

func (todos *Todos) update(index int, title string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Status", "Created At", "Completed At")

	for i, todo := range *todos {
		completed := "❌"
		completedAt := ""

		switch todo.Status {
		case 2:
			completed = "Doing"
		case 3:
			completed = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(i), todo.Title, completed, todo.CreatedAt.Format(time.RFC1123), completedAt)

	}

	table.Render()
}

func (todos *Todos) printByStatus(status int) {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Status", "Created At", "Completed At")

	for i, todo := range *todos {
		if todo.Status != status {
			continue
		}

		completed := "❌"
		completedAt := ""

		switch todo.Status {
		case 2:
			completed = "Doing"
		case 3:
			completed = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(i), todo.Title, completed, todo.CreatedAt.Format(time.RFC1123), completedAt)

	}

	table.Render()
}
