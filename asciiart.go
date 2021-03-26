package asciiart

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ConvertToAscii(arg string, fontStyle string) string {
	if fontStyle == "" {
		fontStyle = "standard"
	}
	fontStyle += ".txt"
	res := ""
	file, err := os.Open(fontStyle)
	if err != nil {
		fmt.Println("This file is not available")
		fmt.Println(err.Error())
		return res
	}
	//check for a validity of files with fonts
	content, err := ioutil.ReadFile(fontStyle)
	lines := getLines(content)
	l := len(lines)
	if l%9 != 0 {
		log.Fatal("Invalid template file:\nlength should be a multiple of 9")
		return res
	}

	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	ascii := getASCII(file)
	buf := make([]string, 8)
	str := strings.Split(arg, "\\n")
	for _, a := range str {
		for _, n := range a {
			if n < 32 || n > 126 {
				fmt.Println("Error: Message is not valid")
				return res
			}
			buf = addCh(buf, ascii[n])
		}
		for i := range buf {
			res += buf[i] + "\n"
		}

		buf = make([]string, 8)
	}
	//fmt.Println(res)
	return res
}

func addCh(buf, new []string) []string {
	for i := range buf {
		buf[i] = buf[i] + new[i]
	}
	return buf
}

func getASCII(file *os.File) map[rune][]string {
	ascii := make(map[rune][]string, 95)
	scanner := bufio.NewScanner(file)
	buf := make([]string, 8)
	charLine := 0
	asciiChar := 32
	for scanner.Scan() {
		if scanner.Text() == "" {
			charLine = 0
			buf = nil
			continue
		} else {
			buf = append(buf, scanner.Text())
			if asciiChar == 127 { //break the loop when, 96 char is read, as there are only 95 chars
				asciiChar = 0
				break
			}
			if charLine == 7 {
				ascii[rune(asciiChar)] = buf
				buf = nil
				charLine = 0
				asciiChar++
				continue
			}
			charLine++
		}
	}
	return ascii
}

func getLines(content []byte) []string {
	lines := []string{}
	currLine := ""
	for i := 0; i < len(content); i++ {
		currLine += string(content[i])
		if content[i] == '\n' {
			lines = append(lines, currLine)
			currLine = ""
		}
	}
	return lines
}
