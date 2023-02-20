package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
)

func generateData(stringLength int, fileName string) {
	if stringLength < 14 {
		stringLength = problemSize
	}
	n := rand.Intn(9)
	answerIndex := stringLength * n / (n + 1) //at least in the second half
	log.Println("Answer Index:", answerIndex+14)
	dataFile, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer dataFile.Close()
	w := bufio.NewWriter(dataFile)

	//alphabet := "abcdefghijklmnopqrstuvwxyz"
	nonUnique := "abcdefghijklm"
	ans := "abcdefghijklmnopq"
	var randomRune int
	var val byte
	for i := 0; i < (stringLength); i++ {
		//We can guarentee a solution, need at least one char after to return position
		if (answerIndex <= i) && (i <= answerIndex+15) {
			val = ans[i%14]
		} else {
			randomRune = rand.Intn(13)
			val = nonUnique[randomRune]
		}
		err := w.WriteByte(val)
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}
	w.Flush()
}

func generateBufferedData(stringLength int) ([]byte, int) {
	if stringLength < 14 {
		stringLength = problemSize
	}
	n := rand.Intn(9)
	answerIndex := stringLength * n / (n + 1) //at least in the second half
	buf := make([]byte, stringLength)

	nonUnique := "abcdefghijklm"
	ans := "abcdefghijklmnopq"
	var randomRune int
	var val byte
	for i := 0; i < (stringLength); i++ {
		if (answerIndex <= i) && (i <= answerIndex+14) {
			val = ans[i%14]
		} else {
			randomRune = rand.Intn(13)
			val = nonUnique[randomRune]
		}
		buf[i] = val
	}
	return buf, answerIndex + 14
}
