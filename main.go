package main

import (

	// "CHaserServer/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

func main(){
  server := app.New()
  main_window := server.NewWindow("CHaserServer | Main")
  main_window.Resize(fyne.NewSize(1280, 720))


	main_window.ShowAndRun() 
}
