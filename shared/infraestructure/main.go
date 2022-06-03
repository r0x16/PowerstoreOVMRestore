package infraestructure

import (
	"fmt"
	"log"
	"os"

	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/database"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui"
)

func NewMain() {
	initConfiguration()
	connectToDB()
	defer closeDBConnection()

	gui.NewMainWindow("PowerstoreOVMRestore")
	gui.MainWindowContainer.SetMainContent()
	gui.MainWindowContainer.PlayMainApplication()
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
