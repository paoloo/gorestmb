package main

import (
  "encoding/json"
  "log"
  "net/http"
  "io"
  "io/ioutil"
  "strings"
  "strconv"
)

type modbusReg struct {
  Addres int   `json:"address"`
  Values []int `json:"values"`
}

func JSONify(data []int, w http.ResponseWriter ) {
  bJSON, err := json.Marshal(data); if err != nil { log.Println("error:", err) }
  w.Header().Set("Content-type", "application/json")
  w.Write(bJSON)
}

func handleModbus(w http.ResponseWriter, r *http.Request) {
  parms := strings.Split( r.URL.Path[len("/modbus/"):] , "/" )

  if r.Method == "GET" {
    if len(parms) > 1 {
      _address, err := strconv.Atoi(parms[0]); if err != nil { w.WriteHeader(422) }
      _size, errr := strconv.Atoi(parms[1]); if errr != nil { w.WriteHeader(422) }
      res, _ := mb.ReadHoldingRegister(_address, _size)
      JSONify(res, w)
    } else {
      w.WriteHeader(422)
    }

  } else if r.Method == "POST" {
    _body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)); if err != nil { log.Println(err) }
    var _data modbusReg
    if err := json.Unmarshal(_body, &_data); err != nil {
      w.WriteHeader(422)
    } else {
      res,_ := mb.WriteRegisters(_data.Addres , _data.Values)
      JSONify(res, w)
    }

  } else if r.Method == "PUT" {
    _body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)); if err != nil { log.Println(err) }
    var _data modbusReg
    if err := json.Unmarshal(_body, &_data); err != nil {
      w.WriteHeader(422)
    } else {
      _address, err := strconv.Atoi(parms[0]); if err != nil { w.WriteHeader(422) }
      res,_ := mb.WriteRegisters(_address,_data.Values)
      JSONify(res, w)
    }

  } else if r.Method == "DELETE" {
    _address, err := strconv.Atoi(parms[0]); if err != nil { w.WriteHeader(422) }
    res,_ := mb.WriteRegister(_address,0x00)
    JSONify(res, w)

  } else {
    w.WriteHeader(501) /* not implemented */
  }

}
