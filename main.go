package main

import (
  "io/ioutil"
  "fmt"
  "github.com/davidlares/go-protobuf/src/simple"
  "github.com/golang/protobuf/proto"
  "log"
)

func main() {
  fmt.Println("Hello, Protobuf Go")
  sm:= doSimple()
  // write
  fmt.Println("Writing file")
  writeToFile("simple.bin", sm)
  // read
  sm2 := &simplepb.SimpleMessage{}
  fmt.Println("Reading file")
  readFromFile("simple.bin", sm2)
  fmt.Println(sm2)
}

// returning a pointer
func doSimple() *simplepb.SimpleMessage {
  // instance of SimpleMessage
  sm := simplepb.SimpleMessage {
    Id: 12345,
    IsSimple: true,
    Name: "Go Simple Message",
    SampleList: []int32{1,2,3,4},
  }
  // printing object
  fmt.Println(sm)
  // renaming attrs
  sm.Name = "GO Renamed Simple Message"
  fmt.Println("ID: ", sm.GetId())

  return &sm
}

func writeToFile(fname string, pb proto.Message) error {
  out, err := proto.Marshal(pb)
  if (err != nil) {
    log.Fatalln("Can't serialize to bytes", err)
    return err
  }
  // writing file
  if err:= ioutil.WriteFile(fname, out, 0644); err != nil {
    log.Fatalln("Can't write to file")
  }
  fmt.Println("Data written to file")
  return nil
}

func readFromFile(fname string, pb proto.Message) error {
  in, err := ioutil.ReadFile(fname)
  if(err != nil) {
    log.Fatalln("Error reading: ", err)
    return err
  }
  // byte array to protobuf
  err2 := proto.Unmarshal(in, pb)
  if (err2 != nil) {
    log.Fatalln("Cannot transform bytes to protobuf: ", err2)
    return err2
  }
  return nil
}
