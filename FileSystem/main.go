package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

// todo: create file if not exists
// 	todo: show err message: `key does not exist`
// todo: implement del command

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please provide the arguments")
		return
	}

	switch args[0] {
	case "set":
		if len(args) < 3 || args[1] == "" {
			fmt.Println("please provide correct key value pair")
			return
		}
		setFunction(args[1], args[2])

	case "get":
		if len(args) < 2 || args[1] == "" {
			fmt.Println("Please provide the key value")
			return
		}
		getFunction(args[1])
	case "del":
		if len(args) < 2 || args[1] == "" {
			fmt.Println("Please provide the key value")
			return
		}
		delFunction(args[1])
	default:
		fmt.Println("Please enter correct argument")
		return
	}

}

func setFunction(key string, value string) {
	f, err := os.OpenFile("myfile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(key + "@" + value + "\n"); err != nil {
		panic(err)
	}

}

func getFunction(key string) {
	file, err := os.OpenFile("myfile.txt", os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("failed reading data from file, file doesn't exist")
	}

	scanner := bufio.NewScanner(file)
	flag := false
	for scanner.Scan() {
		res := strings.Split(scanner.Text(), "@")
		if res[0] == key {
			flag = true
			fmt.Println(res[1])
			return
		}

	}
	if !flag {
		fmt.Println("Key doesn't exist")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func delFunction(key string) {
	file, err := os.OpenFile("myfile.txt", os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("failed reading data from file, file doesn't exist")
	}
	var bs []byte
	buf := bytes.NewBuffer(bs)
	flag := false
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		res := strings.Split(scanner.Text(), "@")
		if res[0] != key {
			_, err := buf.Write(scanner.Bytes())
			if err != nil {
				log.Fatal(err)
			}
			_, err = buf.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}
		if res[0] == key {
			fmt.Println("delete successful")
			flag = true
		}
	}
	if !flag {
		fmt.Println("Key doesn't exist")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("myfile.txt", buf.Bytes(), 0666)
	if err != nil {
		log.Fatal(err)
	}

}
