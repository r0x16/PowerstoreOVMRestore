package infraestructure

import (
	"fmt"
	"log"
	"os"

	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/database"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui"
)

var MainWindow *gui.MainWindow
var MainLayout *gui.MainLayout

func NewMain() {
	initConfiguration()
	connectToDB()
	defer closeDBConnection()

	MainWindow = gui.NewMainWindow("PowerstoreOVMRestore")
	MainLayout = gui.NewMainLayout()
	MainWindow.SetContent(MainLayout)
	MainWindow.Play()
}

func initConfiguration() {
	err := config.Initialize()
	if err != nil {
		log.Fatal("Error loading configuration file, application cannot start", err)
	}
}

func connectToDB() {
	err := database.Connect()
	if err != nil {
		log.Fatal("Error connecting to database")
	}
}

func closeDBConnection() {
	err := database.Close()
	if err != nil {
		fmt.Println("Error trying to disconnect to database")
		os.Exit(1)
	}
}
