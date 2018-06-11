package main

import (
        "fmt"
        "goblog/accountservice/service"
        "goblog/accountservice/dbclient"
)

var appName = "accountservice"

func main() {
        fmt.Printf("Starting %v\n", appName)
        initializeBoltClient()
        service.StartWebServer("6767")
}

func initializeBoltClient() {
        service.DBClient = &dbclient.BoltClient{}
        service.DBClient.OpenBoltDb()
        service.DBClient.Seed()
}
