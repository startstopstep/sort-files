package main

import (
	"github.com/startstopstep/sort-files/sort"
	"log"
)

func main() {
	log.Println("Script launched.")
	files, err := sort.GetFilesInDir()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Script detected %d files in current directory.", len(files))
	num, err := sort.MoveFilesToDirs(files)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Script moved %d files to directories.", num)
}
