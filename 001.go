package main

import "fmt"
import "encoding/json"
import "log"
import "io/ioutil"

type Cat struct {
	Name string `json:"adjustName"`
  Animal string `json:"Breed"`
	Age int `json:"age,omitempty"`
  OwnerPayment string `json:"-"`
}

func main() {
	cat := Cat{Name: "Charlie", Animal: "Tabby", OwnerPayment: "it's totally a secret"}
	cat.Speak()
  b, error := json.Marshal(cat)

  log.Print(error)
  log.Print(string(b))


  RawJson, error2 := ioutil.ReadFile("MyJson.json")
  log.Print(error2)

  var C2 Cat
  error3 := json.Unmarshal(RawJson, &C2)
  log.Print(error3)
  log.Printf("%+v", &C2)
}

func (c Cat) Speak() {
	fmt.Println("MEOW")
}
