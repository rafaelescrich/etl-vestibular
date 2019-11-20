package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/rafaelescrich/etl-vestibular/config"
	"github.com/rafaelescrich/etl-vestibular/db"
	"github.com/rafaelescrich/etl-vestibular/vestibular"
)

func migrate() {
	db.DB.AutoMigrate(&vestibular.CandidateInfo{}, &vestibular.Questionnaire{})
}

func main() {

	err := config.Load()

	if err != nil {
		log.Fatal("Error while initializing config: ", err)
	}

	err = db.Connect()
	if err != nil {
		log.Fatal("Could not connect to database: ", err)
	}

	// Add tables to db if they dont exist
	migrate()

	//populaDimensoes()

	f := flag.String("file", "grade_socioeconomico.csv", "file path to read from")
	flag.Parse()
	data, err := ioutil.ReadFile(*f)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	err = vestibular.Save(data)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("File load successfully")
	}

}
