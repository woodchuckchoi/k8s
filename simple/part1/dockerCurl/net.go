package main

import (
	"fmt"
	"net/http"
	"bufio"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
    	text, _ := reader.ReadString('\n')
    
		text = strings.Replace(text, "\n", "", -1)
		
		resp, err := http.Get(text)
		if err != nil {
			fmt.Println("Something didn't go as expected...")
			continue
		}
		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}
