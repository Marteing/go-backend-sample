package main

import (
	"flag"
	"fmt"
	"net/http"
	"go-backend-sample/dao"
	"go-backend-sample/web"
	"strconv"
)

var (
	port               = 8020
	logLevel           = "warning"
	db                 = 1
	dbConfigFile       = ""
)

// Main
func main() {
	// Get arguments
	flag.IntVar(&port, "p", port, "webserver port")
	flag.IntVar(&db, "db", db, "database (0 - Redis | 1 - Mock)")
	flag.StringVar(&dbConfigFile, "dbConf", dbConfigFile, "config file path")
	flag.StringVar(&logLevel, "log", logLevel, "log level")

	// Parse arguments
	flag.Parse()

	// Get DAO Redis
	taskDAO, err := dao.GetDAO(dao.DBType(db))
	if err != nil {
		fmt.Println(err)
	}

	// New controller
	taskController := web.NewTaskController(taskDAO)

	// New router
	router := web.NewRouter(taskController)
	
	fmt.Println("Starting web server on port : " + strconv.Itoa(port))
	http.ListenAndServe(":"+strconv.Itoa(port), router)
}