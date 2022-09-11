package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Name   string
	Domain string
}

type API int

var database []Item

func (a *API) GetByName(name string, reply *Item) error {
	var getItem Item
	for _, val := range database {
		if val.Name == name {
			getItem = val
			break
		}
	}
     *reply = getItem
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
     *reply = item
	return nil
}

func (a *API) UpdateItem(edit Item, reply *Item) error {
     var changedItem Item
	for idx, val := range database {
		if val.Name == edit.Name {
			database[idx] = edit
               changedItem = database[idx]
		}
	}
     *reply = changedItem
	return nil
}

func (a *API) DeleteItem(deleteItem Item, reply *Item) error{
     var del Item
	for idx, val := range database {
		if val.Name == deleteItem.Name {
               del = database[idx]
			database = append(database[:idx], database[idx+1:]...)
			break
		}
	}
     *reply = del
     return nil
}

func (a *API) GetDB(name string, reply *[]Item) error {
    *reply = database
	return nil
}

func main() {

     var api = new(API)
     err := rpc.Register(api)
     if err != nil {
          log.Fatalln("Error registering API:", err)
     }
     rpc.HandleHTTP()

     listener, err := net.Listen("tcp", ":4040")
     if err != nil {
          log.Fatal("Error listening on port 4040", err)
     }
     log.Println("Serving RPC on port 4040")
     err = http.Serve(listener, nil)
     if err != nil {
          log.Fatal("Error serving", err)
     }

	// fmt.Println("Initial database: ", database)
	// a := Item{"ijas", "GO"}
	// b := Item{"niyas", "ML"}
	// c := Item{"rishal", "Flutter"}
     // d := Item{"jk", "MERN"}

     // AddItem(a)
     // AddItem(b)
     // AddItem(c)
     // AddItem(d)

     // fmt.Println("Database: ", GetDB())

     // fmt.Println("Searched item: ", GetByName("jk"))

     // fmt.Println("Updated: ", UpdateItem("jk", Item{"JK", "MERN"}))

     // fmt.Println("Deleted: ", DeleteItem("ijas"))

     // fmt.Println("Database: ", GetDB())
}
