package main

import (
  "os"
  "io"
  "log"
	// "CHaserServer/game"
	// "CHaserServer/network"
  "CHaserServer/utils"
)

func main(){
  file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  logger := utils.NewCustomLogger(io.MultiWriter(file, os.Stdout))

  logger.Info("This is an information message")

  logger.Warning("This is a warning message")

  logger.Important("This is an important message")

  logger.Error(err)


}
