package main

import (
	"fmt"
	"sync"
)

type dbConnection struct {
	url      string
	userName string
	password string
}

var mu = &sync.Mutex{}
var dbClient *dbConnection

func getInstance(url, user, pwd string) *dbConnection {
	if dbClient == nil {
		mu.Lock()
		defer mu.Unlock()
		if dbClient == nil {
			fmt.Println("dbClient created")
			dbClient = &dbConnection{
				url:      url,
				userName: user,
				password: pwd,
			}
		} else {
			fmt.Println("Client already exist")
		}
	} else {
		fmt.Println("Client already exist")
	}
	return dbClient
}

func main() {
	for i := 0; i < 10; i++ {
		go getInstance("url", "user1", "user123")
	}
}
