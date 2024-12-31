# Todo App CLI

A simple command-line todo application written in Go.

> Practice project: https://roadmap.sh/projects/task-tracker/solutions?u=64d3087faa497d7fa5190888

## Features

- Add new tasks
- Change status of a task
- Edit existing tasks
- Delete tasks
- List all tasks
- List tasks by status

## Installation

1. Make sure you have Go installed on your system.
2. Clone this repository: `https://github.com/justasteri/todo-cli`
3. Navigate to the project directory: `cd todo-cli`
4. Build the application: `go build`

## Usage

`todo-cli [command] [options]`

### Commands:

- `add` 
    - Add a new task 
- `list` 
    - List all tasks toggle 
- `change_status` 
    - Change the status of a task
- `update` 
    - Update an existing task 
- `delete` 
    - Delete a task

Options:

* --help, -h 
    - Show help information
* --add string
    - Task description to add
* --list 
    - List all tasks
* --list_status int
    - Status Indexes:
        1. To Do
        2. Doing
        3. Done
    - List tasks by status
* --change_status string
    - Format: `Index:Status_Index`
        - (e.g., 1:2) 
    - Status Indexes:
        1. To Do
        2. Doing
        3. Done
    - Change the status of a task
* --update string 
    - Format: `Index:Status_Index`
        - (e.g., 1:Updated task) 
    - Update the title of a task 
* --delete int 
    - Index of the task to delete 

### Examples

- **Add a new task:**

`todo-cli --add "Buy groceries"`

- **List all tasks:**

`todo-cli --list`

- **Change the status of the task 1 to Doing:**

`todo-cli --change_status 1:2`

- **Update the second task:**

`todo-cli --update "2:Finish project report"`

- **Delete the third task:**

`todo-cli --delete 3`

## Data Storage

The application stores todo items in a JSON file named `todos.json` in the current directory.

## Resources
https://youtu.be/g16Zf0KQEWI?si=7LwKtliAWMerRmkw

## Contributing

Contributions are welcome! Please feel free to open issues or submit pull requests.


## License

This project is licensed under the MIT License. 
