package main

import (
	"fmt"
	"os"

	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/model"
)

var MainWindow *gui.MainWindow

func main() {
	fmt.Println("START")
	initConfiguration()
	connectToDB()
	defer closeDBConnection()
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

func connectToDB() {
	err := model.Connect()
	if err != nil {
		fmt.Println("Error connecting to database")
		os.Exit(1)
	}
}

func closeDBConnection() {
	err := model.Close()
	if err != nil {
		fmt.Println("Error trying to disconnect to database")
		os.Exit(1)
	}
}
