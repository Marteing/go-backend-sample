package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"go-backend-sample/dao"
	"go-backend-sample/utils"
	"go-backend-sample/web"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strconv"
	"time"
)

var (
	Version string
	BuildStmp string
	GitHash string

	port               = 8020
	statisticsDuration = 20 * time.Second
	logLevel           = "warning"
	db                 = 2
	dbConfigFile       = ""
)

func main() {
	app := cli.NewApp()
	app.Name = utils.AppName
	app.Usage = "todolist service launcher"

	timeStmp, err := strconv.Atoi(BuildStmp)
	if err != nil {
		timeStmp = 0
	}
	app.Version = Version + ", build on " + time.Unix(int64(timeStmp), 0).String() + ", git hash " + GitHash
	app.Authors = []cli.Author{{Name: "xma"}}
	app.Copyright = "Copyright " + strconv.Itoa(time.Now().Year())

	// command line flags
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Value:       port,
			Name:        "port, p",
			Usage:       "Set the listening port of the webserver",
			Destination: &port,
		},
		cli.DurationFlag{
			Value:       statisticsDuration,
			Name:        "statd, s",
			Usage:       "Set the statistics accumulation duration (ex : 1h, 2h30m, 30s, 300ms)",
			Destination: &statisticsDuration,
		},
		cli.StringFlag{
			Value:       logLevel,
			Name:        "logl, l",
			Usage:       "Set the output log level (debug, info, warning, error)",
			Destination: &logLevel,
		},
		cli.IntFlag{
			Value:       db,
			Name:        "db",
			Usage:       "Set the database connection parameters (0 - Redis | 1 - MongoDB | 2 - Mock)",
			Destination: &db,
		},
		cli.StringFlag{
			Value:       dbConfigFile,
			Name:        "dbConfig",
			Usage:       "Set the path of database connection parameters file",
			Destination: &dbConfigFile,
		},
	}

	app.Action = func(c *cli.Context) error {

		err := utils.InitLog(logLevel)
		if err != nil {
			logrus.Warn("error setting log level, using debug as default")
		}

		taskDAO, err := dao.GetDAO(dao.DBType(db), dbConfigFile)
		if err != nil {
			logrus.WithField("db", db).WithField("dbConfigFile", dbConfigFile).Error("unable to build the required DAO")
		}

		webServer := negroni.New()

		recovery := negroni.NewRecovery()
		recovery.PrintStack = false
		webServer.Use(recovery)

		// add middleware n.Use()

		taskCtrl := web.NewTaskController(taskDAO)

		router := web.NewRouter(taskCtrl)

		webServer.UseHandler(router)

		webServer.Run(":" + strconv.Itoa(port))

		return nil
	}

	err = app.Run(os.Args)
	if err != nil {
		logrus.Error("run error %q\n", err)
	}
}
