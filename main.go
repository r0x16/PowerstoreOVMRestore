package main

import (
	"fmt"

	"github.com/r0x16/PowerstoreOVMRestore/gui"
)

var MainWindow *gui.MainWindow

func main() {
	fmt.Println("START")
	MainWindow = gui.NewMainWindow("PowerstoreOVMRestore")
	MainLayout := gui.NewMainLayout()
	MainWindow.SetContent(MainLayout)
	MainWindow.Play()
}
