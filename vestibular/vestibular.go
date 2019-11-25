package vestibular

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rafaelescrich/etl-vestibular/db"
)

// GradeSocioeconomico holds foreign keys to question and code
type GradeSocioeconomico struct {
	gorm.Model
	Question Question
	Code     Code
}

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
	AcertosTotal         float64
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

// Code has all codes to all questions
type Code struct {
	gorm.Model
	IDEvento        int
	NumeroQuestao   int
	IDCodigo        int
	DescricaoCodigo string
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

// SaveQuestions file questions to database
func SaveQuestions(fileQuestionnaire []byte) (err error) {

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

	}

	for _, q := range questions {
		db.DB.Create(&q)
	}

	fmt.Println("Finished creating all questions")

	return nil
}

// SaveCodes file codes to database
func SaveCodes(fileCode []byte) (err error) {

	linesCode := scanLines(fileCode)
	if err != nil {
		return
	}

	codes := make([]Code, len(linesCode)-1)

	for i, str := range linesCode {
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
		idcod, err := strconv.Atoi(s[2])
		if err != nil {
			return err
		}

		codes[i-1] = Code{
			IDEvento:        idEv,
			NumeroQuestao:   numQ,
			IDCodigo:        idcod,
			DescricaoCodigo: s[3],
		}
	}

	for _, c := range codes {
		db.DB.Create(&c)
	}

	fmt.Println("Finished creating all codes")

	return nil
}

// SaveCandidatesInfo file candidate to database
func SaveCandidatesInfo(fileCandidates []byte) (err error) {

	timeLayout := "Jan 02 2006 03:04PM"

	linesCandidates := scanLines(fileCandidates)
	if err != nil {
		return
	}

	candidatesInfo := make([]Code, len(linesCandidates)-1)

	for i, str := range linesCandidates {
		if i == 0 {
			continue
		}
		s := strings.Split(str, ",")
		fmt.Println(s)

		idEv, err := strconv.Atoi(s[0])
		if err != nil {
			return err
		}

		idcand, err := strconv.Atoi(s[1])
		if err != nil {
			return err
		}

		datnasc, err := time.Parse(timeLayout, s[2])
		if err != nil {
			return err
		}

		idling, err := strconv.Atoi(s[3])
		if err != nil {
			return err
		}

		idloc, err := strconv.Atoi(s[4])
		if err != nil {
			return err
		}

		mes2grau, err := strconv.Atoi(s[5])
		if err != nil {
			return err
		}

		ano2grau, err := strconv.Atoi(s[6])
		if err != nil {
			return err
		}

		idsex := s[7]

		porexp := s[8]
		// Remove all candidates that are doing the vestibular to gain experience
		if s[8] == "S" {
			continue
		}

		bairro := s[9]

		cidade := s[10]

		unidfed := s[11]

		//gradsocioecon

		class, err := strconv.Atoi(s[13])
		if err != nil {
			return err
		}

		estab := s[14]

		f, err := strconv.ParseFloat(s[14], 64)
		if err != nil {
			return err
		}

		idrac, err := strconv.Atoi(s[15])
		if err != nil {
			return err
		}

		idcat, err := strconv.Atoi(s[16])
		if err != nil {
			return err
		}

		ens, err := strconv.Atoi(s[17])
		if err != nil {
			return err
		}

		candidatesInfo[i-1] = CandidateInfo{
			IDEvento:             idEv,
			IDCandidato:          idcand,
			DataNascimento:       datnasc,
			IDLingua:             idling,
			IDLocal:              idloc,
			MesSegundoGrau:       mes2grau,
			AnoSegundoGrau:       ano2grau,
			IDSexo:               idsex,
			PorExperiencia:       porexp,
			Bairro:               bairro,
			Cidade:               cidade,
			UnidadeFederativa:    unidfed,
			GradeSocioeconomico:  "mudar pra outra estrutura",
			ClassificacaoGeral:   class,
			Estabelecimento2Grau: estab,
			AcertosTotal:         f,
			IDRaca:               idrac,
			IDCategoria:          idcat,
			EnsinoPublico:        ens,
		}
	}

	for _, cand := range candidatesInfo {
		db.DB.Create(&cand)
	}

	fmt.Println("Finished creating all codes")

	return nil
}

func parseAnswer(answer string) (grad GradeSocioeconomico) {
	return
}
