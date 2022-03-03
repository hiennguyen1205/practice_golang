package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Test struct {
	Question string
	Answer   string
}

func main() {
	//open file
	f := OpenFile("problems.csv")

	//close file at the end program
	defer CloseFile(f)

	//read csv values using csv.Reader
	listTest := ReadFileCSV(f)

	//print question console, input answer and calculate your score
	yourScore, totalTest := CalculateYourScore(listTest)
	fmt.Printf("Your score is %d/%d", yourScore, totalTest)
}

func OpenFile(nameFile string) *os.File {
	f, err := os.Open(nameFile)
	if err != nil {
		log.Print("Error when open file csv: ", err)
		return nil
	}
	return f
}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatalf("Error when close file, err: %v", err.Error())
	}
}

func ReadFileCSV(file *os.File) []Test {
	csvReader := csv.NewReader(file)
	listTest := make([]Test, 0)
	for {
		res, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error when read file, ERROR:%v", err.Error())
		}
		if len(res) > 1 {
			listTest = append(listTest, Test{
				Question: res[0],
				Answer:   res[1],
			})
		}
	}
	return listTest
}

func CalculateYourScore(listTest []Test) (yourScore int64, totalTest int64) {
	readInput := bufio.NewReader(os.Stdin)
	yourScore = 0
	totalTest = int64(len(listTest))
	for _, test := range listTest {
		fmt.Print(test.Question, ": ")
		yourAnswer, err := readInput.ReadString('\n')
		//do chay tren window nen can format lai yourAnswer (do dinh dang dong o window khac linux)
		//yourAnswer = 'input\r\n`
		yourAnswer = strings.Replace(yourAnswer, "\r\n", "", -1)
		if err != nil {
			log.Fatalf("Error when scan input from keyboard, err: %v", err.Error())
			return 0, 0
		}
		if yourAnswer == test.Answer {
			yourScore++
		}
	}
	return yourScore, totalTest
}
