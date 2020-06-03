package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// GetInput gets user input
func GetInput() (string, *Buyer) {
	reader := bufio.NewReader(os.Stdin)
	phone := ReadInput("Your phone: +", reader, validatePhone)
	email := ReadInput("Your email: ", reader, validateEmail)
	notifyType := ReadInput("Which notification channel to send. \n1 (phone), 2 (email), 3 (both): ", reader, validateChanelType)
	c := &Buyer{
		phone: "+" + phone,
		email: email,
	}
	return notifyType, c

}

func validatePhone(text string) bool {
	if !strings.HasPrefix(text, "+") {
		text = text[1:]
	}
	if _, err := strconv.Atoi(text); err != nil {
		return false
	}
	return true
}

func validateEmail(text string) bool {
	var alpaNum = `[a-zA-Z0-9\.]{1,}`
	re := regexp.MustCompile(alpaNum + `@` + alpaNum + `\.[a-zA-Z]`)
	return re.MatchString(text)
}

func validateChanelType(text string) bool {
	if _, ok := notifyMap[text]; ok {
		return true
	}
	return false
}

// ReadInput reads input from user
func ReadInput(msg string, reader *bufio.Reader, validate func(text string) bool) string {
	for {
		fmt.Print(msg)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("%v", err)
			continue
		}
		text = strings.TrimSpace(text)
		if len(text) == 0 {
			continue
		}
		if !validate(text) {
			log.Printf("Ivalid %v", text)
			continue
		}
		return text
	}
}
