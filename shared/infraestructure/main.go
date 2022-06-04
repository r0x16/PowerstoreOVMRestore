package infraestructure

import (
	"fmt"
	"log"
	"os"

	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/database"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui"
	sites "github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure"
)

func NewMain() {
	initConfiguration()
	connectToDB()
	defer closeDBConnection()

	gui.NewMainWindow(config.Global.APP_NAME)
	gui.MainWindowContainer.SetMainContent()

	sites.NewSitesModule()

	gui.MainWindowContainer.PlayMainApplication()
}

// Initialize Global configuration and i18n support
func initConfiguration() {
	if err := config.Initialize(); err != nil {
		log.Fatal("Error loading configuration file, application cannot start", err)
	}

	if err := config.InitializeLang(); err != nil {
		log.Fatal("Error loading language configuration file, application cannot start", err)
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
