package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func MakePasswordInFile(pass string) (bool, error) {
	file, err := os.Create("private_key.txt")
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	defer file.Close()

	file.Write([]byte(pass))
	return true, nil
}

func ReadPasswordInFile() (string, error) {
	file, err := os.Open("private_key.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	fmt.Println(wr.String())
	//TODO fix error
	//TODO fix response
	fmt.Println("### ReadFile ###")
	return wr.String(), nil

}

func WritePasswordInFile(path, password string) (string, error) {
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write([]byte(password))

	return "", nil

}
