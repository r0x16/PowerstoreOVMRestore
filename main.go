package main

import (
	"fmt"
	"os"

	"github.com/r0x16/PowerstoreOVMRestore/main/infraestructure/view"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/database"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/window"
)

var MainWindow *window.MainWindow

func main() {
	fmt.Println("START")
	initConfiguration()
	connectToDB()
	defer closeDBConnection()
	MainWindow = window.NewMainWindow("PowerstoreOVMRestore")
	MainLayout := view.NewMainLayout()
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

func connectToDB() {
	err := database.Connect()
	if err != nil {
		fmt.Println("Error connecting to database")
		os.Exit(1)
	}
}

func closeDBConnection() {
	err := database.Close()
	if err != nil {
		fmt.Println("Error trying to disconnect to database")
		os.Exit(1)
	}
}
