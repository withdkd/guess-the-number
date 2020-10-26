package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(10) + 1
	for true {
		fmt.Print("Type: ")
		str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		str = strings.TrimSpace(str)
		guess, _ := strconv.Atoi(str)
		if target == guess {
			fmt.Println("Yes")
			break
		} else {
			fmt.Println("No")
		}
	}
	fmt.Print("you win")
}
