package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Name string
     Domain string
}

func main() {
     var reply Item
     var db []Item

     client, err := rpc.DialHTTP("tcp", "localhost:4040")
     if err != nil {
          log.Fatal("RPC Connection Error:", err)
     }
     a := Item{"ijas", "GO"}
	b := Item{"niyas", "ML"}
	c := Item{"rishal", "Flutter"}
     d := Item{"jk", "MERN"}

    err =  client.Call("API.AddItem", a, &reply)
    if err != nil {
          log.Println("Error adding item:", err)
    }
     client.Call("API.AddItem", b, &reply)
     client.Call("API.AddItem", c, &reply)
     client.Call("API.AddItem", d, &reply)

     client.Call("API.GetDB", "", &db)
     fmt.Println(db)

     search := "jk"
     client.Call("API.GetByName", search, &reply)
     fmt.Println("Get by name: ", search, reply)

     err = client.Call("API.UpdateItem", Item{"rishal", "Python"}, &reply)
     if err != nil {
          log.Println("Error adding item:", err)
    }
     fmt.Println("Updated: ", reply)

     err = client.Call("API.DeleteItem", Item{"ijas", ""}, &reply)
     if err != nil {
          log.Println("Error adding item:", err)
    }
     fmt.Println("Deleted", reply)

     client.Call("API.GetDB", "", &db)
     fmt.Println(db)
}