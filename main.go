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
	db.DB.AutoMigrate(&vestibular.Question{}, &vestibular.Code{}, &vestibular.CandidateInfo{})
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

	f := flag.String("file", "grade_socioeconomico.csv", "file path to read from")
	flag.Parse()
	data, err := ioutil.ReadFile(*f)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	err = vestibular.SaveQuestions(data)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("File load successfully")
	}

	f2 := flag.String("file2", "codigo_questionario.csv", "file path to read from")
	flag.Parse()
	dataCode, err := ioutil.ReadFile(*f2)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	err = vestibular.SaveCodes(dataCode)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("File load successfully")
	}

	f3 := flag.String("file3", "candidato.csv", "file path to read from")
	flag.Parse()
	dataCandidato, err := ioutil.ReadFile(*f3)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	err = vestibular.SaveCandidatesInfo(dataCandidato)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("File load successfully")
	}

}
