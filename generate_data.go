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
	dataFile, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer dataFile.Close()
	w := bufio.NewWriter(dataFile)

	runes := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < (stringLength); i++ {
		randomRune := rand.Intn(26)
		err := w.WriteByte(runes[randomRune])
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}
	w.Flush()
	//We can guarentee a solution, need at least one char after to return position
	_, err = w.WriteString("abcdefghijklmnzzz")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	w.Flush()
}
