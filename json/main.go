package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// json.Marshal()
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func ReadJSON(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading user.json", err)
		return nil, err
	}
	fmt.Println("Success reading user.json")
	return data, nil
}
