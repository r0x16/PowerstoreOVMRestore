package main

import (
	"fmt"
	"os"

	"github.com/r0x16/PowerstoreOVMRestore/config"
	"github.com/r0x16/PowerstoreOVMRestore/gui"
)

var MainWindow *gui.MainWindow

func main() {
	fmt.Println("START")
	initConfiguration()
	MainWindow = gui.NewMainWindow("PowerstoreOVMRestore")
	MainLayout := gui.NewMainLayout()
	MainWindow.SetContent(MainLayout)
	MainWindow.Play()
}

func initConfiguration() {
	err := config.Initialize()
	if err != nil {
		fmt.Println("Error loading configuration file, application cannot start", err)
		os.Exit(1)
	}
}
