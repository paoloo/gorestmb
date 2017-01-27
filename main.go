package main

import (
  "log"
  "net/http"
  "github.com/paoloo/modbuscli"
)

var mb = new(modbuscli.ModBus)

func main() {
  mb.EndPoint = "127.0.0.1:502"
  mb.Addr = 0x01
  http.HandleFunc("/modbus/", handleModbus)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

