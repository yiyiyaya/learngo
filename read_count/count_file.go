package main

import (
	"bufio"
	//"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//os.Create("hello.go")
	infos, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range infos {
		log.Println(file.Name())
		f, err := os.Open(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		r := bufio.NewReader(f)
		stat := make(map[rune]int)
		for {
			ch, _, err := r.ReadRune()
			//log.Println(ch)
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			stat[ch] = stat[ch] + 1
		}
		for r, i := range stat {
			log.Println(string(r), i)
		}
	}
}
