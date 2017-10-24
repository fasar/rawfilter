package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	rawDir = "./raw"
)

func listFiles(files []os.FileInfo, suffix string) []string {
	filesSize := len(files)
	res := make([]string, 0, filesSize)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), suffix) {
			res = append(res, f.Name())
		}
	}
	return res
}

func filterDstFile(src []string, dst []string) []string {
	res := make([]string, 0, len(dst))

	for _, f := range dst {
		if !contains(src, f) {
			res = append(res, f)
		}
	}
	return res
}

func contains(s []string, e string) bool {
	if len(e) <= 4 {
		return false
	}
	eName := e[:len(e)-4]
	for _, a := range s {
		if len(a) <= 4 {
			continue
		}
		aName := a[:len(a)-4]
		if aName == eName {
			return true
		}
	}
	return false
}

func main() {
	var rawFilesInfo []os.FileInfo
	var rawFiles []string

	os.MkdirAll(rawDir, os.ModePerm)
	rawFilesInfo, err := ioutil.ReadDir(rawDir)
	if err == nil {
		rawFiles = listFiles(rawFilesInfo, ".CR2")
	} else {
		fmt.Println("Error when reading raw dir : ")
		fmt.Println(err)
	}

	fmt.Print("RAW files : ")
	fmt.Println(rawFiles)

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	imgFiles := listFiles(files, ".JPG")
	fmt.Print("Images files : ")
	fmt.Print(imgFiles)
	fmt.Print("\n")

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".CR2") {
			rawFiles = append(rawFiles, f.Name())
			os.Rename(f.Name(), rawDir+"/"+f.Name())
		}
	}

	toRemove := filterDstFile(imgFiles, rawFiles)

	fmt.Println("Files to remove : ")
	fmt.Println(toRemove)

	for _, f := range toRemove {
		toSuppr := filepath.Join(rawDir, f)
		os.Remove(toSuppr)
	}
}
