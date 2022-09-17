# _To go_
A simple CLI based To do tracker built with Go

# Instalation

Simply download the binary corresponding to your machine, move it to a directory of your preference and add its path to your PATH environment variable.

# Usage

Robust and easy to use, _To go_ supports the following option:

 - `to-go help`: Shows helpful information on the options available
 - `to-go add`: Adds a task based of on the users input
 - `to-go remove`: Removes a task selected by the user
 - `to-go edit`: Edits a task selected by the user
 - `to-go update`: Updates the status of a task selected by the user
 - `to-go finish`: Updates the status of a task selected by the user to 'done'
 - `to-go list`: Lists all the tasks, displaying their titles, priorities, statuses and dates
 - `to-go describe`: Displays all the info pertaining to a task chosen by you
 - `to-go dashboard`: Spins up a lightweight `http` server to a port of your choice ([see more](#customization)), with a interactive UI being served on the `/`;

# Customization

_To go_ offers a few values for you to play around and change to whatever you think suits you better. In the directory you placed the binary, a config file, seen bellow, will be automatically created after the first execution.

```yaml
// to_go.cfg.yaml

useUnicode: true
OpenBrowser: true
storage: sqlite
dashboardPort: "8080"
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
| `OpenBrowser` | Whether or not your system's default internet browser should be opened on the Dashboard UI when the `dashboard` option is ran | `true` or `false` |
| `storage` | Determines the way _To go_ stores data | Anything other than `sqlite` will defaults to using `YAML` files for persistency |
| `dashboardPort` | To which port the Dashboard's `http` server should listen | A string value with any number; Defaults to "8080" if not provided or otherwise invalid |
| `colors` | Defines the colors to be used in the output | The value of its properties should be one of [these colors](#colors)* |

(*) Changing the `reset` value does not apply any changes on behavior, and the field will return to its original value after executing _To go_ again.

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
