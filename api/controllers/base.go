package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgress database driver

	"github.com/Progete-Dev/go-whatsapp-rest/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error
	var DBURL string
	if Dbdriver == "mysql" {
		DBURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		fmt.Printf("%s\n", DbPassword)
	} else if Dbdriver == "postgres" {
		DBURL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s")
	} else {
		fmt.Printf("Cannot connect to %s database\n", Dbdriver)
	}
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}
	server.DB.Debug().AutoMigrate(&models.User{})      //database migration
	server.DB.Debug().AutoMigrate(&models.WpSession{}) //database migration
	server.DB.Debug().AutoMigrate(&models.UserToken{}) //database migration
	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {

	fs := http.FileServer(http.Dir("./storage"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
