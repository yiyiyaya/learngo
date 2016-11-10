package main

import (
	"bufio"
	//"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	//"time"
)

func main() {
	//os.Create("hello.go")
	infos, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	messageChan := make(chan map[rune]int, 10)
	for _, file := range infos {
		log.Println(file.Name())
		//file := file

		go func(fi os.FileInfo) {
			log.Println(fi.Name())
			f, err := os.Open(fi.Name())
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
			//for r, i := range stat {
			//log.Println(string(r), i)
			//}
			messageChan <- stat
		}(file)
	}
	log.Println(len(infos))
	allstat := make(map[rune]int)
	for i := 0; i < len(infos); i++ {
		stat := <-messageChan
		for a, b := range stat {
			allstat[a] = allstat[a] + b

		}

	}
	for a, b := range allstat {
		log.Println(string(a), b)
	}

}
