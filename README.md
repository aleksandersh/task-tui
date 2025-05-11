# Terminal user interface for Task

What is [Task and Taskfile](https://taskfile.dev/)?
> Task is a task runner / build tool that aims to be simpler and easier to use than, for example, [GNU Make](https://www.gnu.org/software/make/).

> Once [installed](https://taskfile.dev/installation), you just need to describe your build tasks using a simple [YAML](http://yaml.org/) schema in a file called Taskfile.yml

The tool is a client with a user interface for the [Task](https://taskfile.dev/) that makes it easier to run tasks from the terminal.

![tuiPack example](./examples/task_tui_screenshot.png "Example")

## Features

- List and execute tasks
- Filtering mode to speed up navigation between tasks
- Task summary page to show the description of a task
- Repeating last executed command

🟥 __The [labels](https://taskfile.dev/usage/#overriding-task-name) are not properly supported right now__ 🟥

If you are using [labels](https://taskfile.dev/usage/#overriding-task-name), you must also specify [namespace alias](https://taskfile.dev/usage/#namespace-aliases) or [task alias](https://taskfile.dev/usage/#task-aliases) for those tasks.

## Usage

#### Cli

Task-tui cli supports common Task arguments, see [task documentation](https://taskfile.dev/reference/cli) or `--help` cli argument.

###### Example usage

```bash
# run for a default taskfile
task-tui
# or specify the taskfile explicitly
task-tui -t ./Taskfile.yml
# ask for help
task-tui --help

# combine with other task arguments
task-tui -xsv -t ./examples
```

###### Repeat last executed command

```bash
task-tui -r
```

#### Hotkeys

##### General

`Enter` - to execute the selected task  
`/` - to enter the filtering mode (in this mode enter a substring to search for and return to the list by pressing `Enter`)  
`s` - to show the task summary  
`Esc` - to go back  
`Ctrl+C` - to exit  
`h` - to show the help page  

##### Task list navigation

`Key down` or `tab` - select a next task  
`Key up` - select a previous task  
`end` (or `fn` + `Key right` on MacOS) - select a last task  
`home` (or `fn` + `Key left` on MacOS) - select a first task  

## Installation

Make sure Task is [installed](https://taskfile.dev/installation/).

#### Homebrew tap

```bash
brew install aleksandersh/task-tui/task-tui
```

#### Go install

Requires Go 1.22

```bash
go install github.com/aleksandersh/task-tui@latest
```

## Best practice

#### Use terminal aliases

Improve your productivity by setting up terminal [aliases](https://www.gnu.org/software/bash/manual/html_node/Aliases.html) for frequently used taskfiles.

###### Bash example

```bash
$ echo $'alias task-tui-example=\'task-tui -x -t "$HOME/taskfile-tui/examples"\'' >> "$HOME/.bash_aliases"
$ source "$HOME/.bash_aliases"
$ task-tui-example
```
