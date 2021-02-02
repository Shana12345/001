package main

import "fmt"
import "encoding/json"
import "log"

type Cat struct {
	Name string
	Age int
	breed string
}

func main() {
	cat := Cat{Name: "Charlie", Age: 3, breed: "Shakes"}
	cat.Speak()
  b, err := json.Marshal(cat)

  log.Print(err)
  log.Print(string(b)) 
}

func (c Cat) Speak() {
	fmt.Println("MEOW")

}
