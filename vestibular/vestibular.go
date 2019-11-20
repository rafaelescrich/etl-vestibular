package vestibular

import (
	"bufio"
	"bytes"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Metadata holds information about the file
// type Metadata struct {
// 	gorm.Model
// 	filename string `json:"filename"`
// 	size     int64  `json:"size"`
// }

// Content has the contents of a file
// type Content struct {
// 	gorm.Model
// 	cpf                string `json:"cpf"`
// 	private            int    `json:"private"`
// 	incompleto         int    `json:"incompleto"`
// 	dataUltimaCompra   string `json:"data_ultima_compra"`
// 	ticketMedio        string `json:"ticket_medio"`
// 	ticketUltimaCompra string `json:"ticket_ultima_compra"`
// 	lojaMaisFrequente  string `json:"loka_mais_frequente"`
// 	lojaUltimaCompra   string `json:"loka_ultima_compra"`
// 	metadataID         uint   `json:"metadata_id"`
// }

// CandidateInfo has all infos about the candidate
type CandidateInfo struct {
	gorm.Model
	IDEvento             int
	IDCandidato          int
	DataNascimento       time.Time
	IDLingua             int
	IDLocal              int
	MesSegundoGrau       int
	AnoSegundoGrau       int
	IDSexo               string
	PorExperiencia       string
	Bairro               string
	Cidade               string
	UnidadeFederativa    string
	GradeSocioeconomico  string
	ClassificacaoGeral   int
	Estabelecimento2Grau string
	AcertosTotal         float32
	IDRaca               int
	IDCategoria          int
	EnsinoPublico        int
}

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

	questions := []Question{}

	for _, str := range linesQuestionnaire {
		var q Question
		q = Question{}
		fmt.Println(str)
	}

	// // lines -1 to remove the header
	// contents := make([]Content, len(lines)-1)

	// for i, line := range lines {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	str := strings.Fields(line)

	// 	priv, err := strconv.Atoi(str[1])
	// 	if err != nil {
	// 		priv = 0
	// 	}

	// 	inc, err := strconv.Atoi(str[2])
	// 	if err != nil {
	// 		inc = 0
	// 	}

	// 	// i-1 because we skip the header of the file
	// 	contents[i-1] = Content{
	// 		cpf:                str[0],
	// 		private:            priv,
	// 		incompleto:         inc,
	// 		dataUltimaCompra:   str[3],
	// 		ticketMedio:        str[4],
	// 		ticketUltimaCompra: str[5],
	// 		lojaMaisFrequente:  str[6],
	// 		lojaUltimaCompra:   str[7],
	// 		metadataID:         id,
	// 	}
	// }

	// // begin a transaction to reduce the time to save in the database
	// tx := db.DB.Begin()

	// // do some database operations in the transaction (use 'tx' from this point, not 'db')
	// for _, content := range contents {
	// 	tx.Create(&content)
	// }

	// // Or commit the transaction
	// tx.Commit()

	return nil
}
