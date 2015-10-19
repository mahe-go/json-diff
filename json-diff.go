package main

import(
	"os"
	"io/ioutil"
	"encoding/json"
	"log"
	"fmt"
	diff "github.com/mahe-go/json-diff/diff"
)

func main() {
	firstFile, err := os.Open(os.Args[1])
	defer firstFile.Close()
	if err != nil {
		panic("Cannot open " + os.Args[1])
	}
	bytes1, err := ioutil.ReadAll(firstFile)
	if err != nil {
		panic("Cannot read " + os.Args[1])
	}
	
	secondFile, err := os.Open(os.Args[2])
	defer secondFile.Close()
	if err != nil {
		panic("Cannot open " + os.Args[2])
	}
	bytes2, err := ioutil.ReadAll(secondFile)
	if err != nil {
		panic("Cannot read " + os.Args[2])
	}
	
	var first interface{}
	err = json.Unmarshal(bytes1, &first)
	if err != nil {
		log.Fatalf("Cannot parse %s: %s", os.Args[1], err)
	}
	var second interface{}
	err = json.Unmarshal(bytes2, &second)
	if err != nil {
		log.Fatalf("Cannot parse %s: %s", os.Args[1], err)
	}
	
	diffresult, err := json.Marshal(diff.Diff(first.(map[string]interface{}), second.(map[string]interface{})))
	if err != nil {
		log.Fatalf("Cannot marshal result: %s", os.Args[1])
	}
	fmt.Printf("%s\n", diffresult)	
}
