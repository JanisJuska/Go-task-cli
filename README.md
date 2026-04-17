# Go Task CLI

A simple command-line task manager written in Go.

Manage your todos directly from the terminal — add tasks, list them, and mark them as done.

---

## Installation

1. Clone the repository:

```bash
git clone https://github.com/JanisJuska/Go-task-cli.git
cd Go-task-cli
```

2. Build the project:

```bash
go build -o task-cli
```

3. (Optional) Move the binary to your PATH:

```bash
mv task-cli /usr/local/bin/
```

Now you can run it from anywhere using:

```bash
task-cli
```

---

## Usage

### Manual page

```bash
task-cli help
```

---

### Add a task

```bash
task-cli add {your task here}
```

---

### List all tasks

```bash
task-cli list
```

Example output:

```
ID   | Title                          | Done
----------------------------------------------
1    | Buy milk                       | ❌
2    | Finish project                 | ✔️
```

---

### Mark a task as done

```bash
task-cli done {id}
```

---

### Remove a task from the list

```bash
task-cli delete {id}
```

---

## Notes

- Tasks are stored in a local `todos.json` file
- The file is created automatically if it doesn’t exist

---

## Features

- Manual page
- Add new tasks
- List all tasks
- Mark tasks as done
- Simple JSON-based storage
- Delete tasks

---

## License

MIT
