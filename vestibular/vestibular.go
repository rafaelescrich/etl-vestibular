package vestibular

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/rafaelescrich/etl-vestibular/db"
)

// Question has all socioeconomic question to the candidate
type Question struct {
	gorm.Model
	IDEvento       int
	NumeroQuestao  int
	PosicaoInicial int
	Tamanho        int
	LiteralQuestao string
}

func scanLines(data []byte) (lines []string) {

	r := bytes.NewReader(data)

	scanner := bufio.NewScanner(r)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return
}

// Save file to database
func Save(fileQuestionnaire []byte) (err error) {

	linesQuestionnaire := scanLines(fileQuestionnaire)
	if err != nil {
		return
	}

	questions := make([]Question, len(linesQuestionnaire)-1)

	for i, str := range linesQuestionnaire {
		if i == 0 {
			continue
		}
		s := strings.Split(str, ",")
		fmt.Println(s)

		idEv, err := strconv.Atoi(s[0])
		if err != nil {
			return err
		}

		numQ, err := strconv.Atoi(s[1])
		if err != nil {
			return err
		}
		posini, err := strconv.Atoi(s[2])
		if err != nil {
			return err
		}
		tam, err := strconv.Atoi(s[3])
		if err != nil {
			return err
		}

		questions[i-1] = Question{
			IDEvento:       idEv,
			NumeroQuestao:  numQ,
			PosicaoInicial: posini,
			Tamanho:        tam,
			LiteralQuestao: s[4],
		}

		// begin a transaction to reduce the time to save in the database
		tx := db.DB.Begin()
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		for _, q := range questions {
			tx.Create(&q)
		}
		// Or commit the transaction
		tx.Commit()

	}

	fmt.Print(questions)

	return nil
}
