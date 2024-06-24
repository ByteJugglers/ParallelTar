package main

import (
	"log"
	"os"
	"strings"
)

var globalContinue map[string]bool = make(map[string]bool)

func init() {
	bytes, _ := os.ReadFile(".continue")
	dirs := strings.Split(string(bytes), "\n")
	for _, dir := range dirs {
		globalContinue[dir] = true
	}
}

func record(file, data string) {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatalf("create record file failed: %s\n", err)
	}
	defer f.Close()

	_, err = f.Write([]byte(data + "\n"))
	if err != nil {
		log.Printf("write record file failed: %s\n", err)
	}

	globalContinue[data] = true
}
