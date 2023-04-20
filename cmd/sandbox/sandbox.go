package main

import (
	"fmt"
	"github.com/gingersamurai/gonban/internal/interfaces/storage"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	ps, err := storage.NewPostgresTaskStorage("host=localhost user=postgres password=15092003 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = ps.DeleteById(9)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(get)
	//id := ps.Add(entity.Task{
	//	Id:          0,
	//	Status:      "DONE",
	//	Name:        "bought beer",
	//	Description: "in pivarius",
	//	Performer:   "alina",
	//	Deadline:    time.Now(),
	//})
	//
	//fmt.Println("id = ", id)
}
