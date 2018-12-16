package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	IndexJunkCode int
)

func OpenFile(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return file
}

func CreateFile(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	return file
}

func Inject(junkCodes []*Variable) {
	file := OpenFile("test.go")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	out := CreateFile("test_2.go")
	defer out.Close()
	buffer := bufio.NewWriter(out)

	var deepLevel int

	for scanner.Scan() {
		line := scanner.Text()

		if deepLevel == 1 {

			for i := 0; i < 2; i++ {
				if IndexJunkCode > len(junkCodes) - 1 {
					fmt.Println("no more junk :/")
					break
				}
				junkCode := junkCodes[IndexJunkCode]

				_, err := buffer.WriteString(fmt.Sprintf(
					"\n%s\nif %s == %d {\n%s = %d\n}\n",
					junkCode.String(),
					junkCode.Name, junkCode.Value+10,
					junkCode.Name, junkCode.Value+10,
				))
				if err != nil {
					panic(err)
				}

				IndexJunkCode++
			}

			deepLevel = 0
		}

		if strings.Index(line, "func") != -1 {
			deepLevel++
		}

		_, err := buffer.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}

	err := buffer.Flush()
	if err != nil {
		panic(err)
	}
}
