# _To go_
A simple CLI based To do tracker built with Go

# Instalation

Simply download the binary corresponding to your machine, move it to a directory of your preference and add its path to your PATH environment variable.

# Usage

Robust and easy to use, _To go_ supports the following option:

 - `to-go help`: Adds a task based of on the users input
 - `to-go add`: Adds a task based of on the users input"
 - `to-go remove`: Adds a task based of on the users input
 - `to-go edit`: Edits a task chosen by you
 - `to-go update`: Updates the status of a task chosen by you
 - `to-go finish`: Updates the status of a task chosen by you to 'done'
 - `to-go list`: Lists all the tasks, including ther titles, priorities, statuses and dates
 - `to-go describe`: Displays all the info pertaining to a task chosen by you

# Customization

_To go_ offers a few values for you to play around and change to whatever you think suits you better. In the directory you placed the binary, a config file, seen bellow, will be automatically created after the first execution.

```yaml
// to_go.cfg.yaml

useUnicode: true
storage: sqlite
colors:
    priority:
        high: red
        medium: yellow
        low: green
    status:
        pending: yellow
        doing: blue
        done: green
    attention: purple
    success: green
    warning: yellow
    error: red
    reset: "\e[0m"

```

| Option | Description | Values |
| ------ | ------------| --------------- |
| `useUnicode` | Whether or not to use unicode characters to represent the task status | `true` or `false`
| `storage` | Determines the way _To go_ stores data | Anything other than `sqlite` will defaults to using `YAML` files for persistency
| `colors` | Defines the colors to be used in the output | Change the value of its properties to one of [these colors](#colors) |

## Colors

There are 9 different colors to choose from:
 - black
 - red
 - green
 - yellow
 - blue
 - purple
 - cyan
 - white
 - default
