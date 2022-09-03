package clihelper

import "github.com/leo-alvarenga/to-go/shared/styles"

var logo = [6]string{
	" _____         _         _____              _",
	"|_   _|       | |       |_   _|            | |",
	"  | | ___   __| | ___     | |_ __ __ _  ___| | _____ _ __",
	"  | |/ _ \\ / _` |/ _ \\    | | '__/ _` |/ __| |/ / _ \\ '__|",
	"  | | (_) | (_| | (_) |   | | | | (_| | (__|   <  __/ |",
	"  \\_/\\___/ \\__,_|\\___/    \\_/_|  \\__,_|\\___|_|\\_\\___|_|",
}

func ShowLogo(style *styles.OutputStyle) {
	for _, line := range logo {
		styles.ShowWithStyle(line, style)
	}
}
