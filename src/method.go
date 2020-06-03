package src

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"strings"
	"time"
)

const (
	phoneMethod      = "1"
	emailMethod      = "2"
	phoneEmailMethod = "3"

	// Creds - for dev purpose placed here
	emailUsername = "AlifCapitalHoldingsTest@gmail.com"
	emailPass     = "AlifCapitalHoldings_2020"
	emailHost     = "smtp.gmail.com"
	emailPort     = ":587"

	phoneFrom  = "+16173907139"
	accountSid = "ACa082605f14e2a3f6717fbe77bbc4b8c2"
	authToken  = "42cd961afbd2bf8740479b2ba1501591"
	urlStr     = "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
)

var notifyMap = map[string]func(msg string, client *Buyer) (string, error){
	phoneMethod:      phone,
	emailMethod:      email,
	phoneEmailMethod: phoneEmail,
}

func phone(msg string, client *Buyer) (string, error) {
	// Set account keys & information
	accountSid := "ACa082605f14e2a3f6717fbe77bbc4b8c2"
	authToken := "42cd961afbd2bf8740479b2ba1501591"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Set up rand
	rand.Seed(time.Now().Unix())

	// Pack up the data for our message
	msgData := url.Values{}
	msgData.Set("To", client.phone)
	msgData.Set("From", phoneFrom)
	msgData.Set("Body", msg)
	msgDataReader := *strings.NewReader(msgData.Encode())

	// Create HTTP request client
	clie := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP POST request and return message SID
	resp, _ := clie.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err != nil {
			fmt.Println(data["sid"])
			log.Println("Phone method error", data["sid"])
		}
	} else {
		log.Println("Phone method error: status: ", resp.Status)
	}
	return "Sent to " + client.phone, nil
}

func email(msg string, client *Buyer) (string, error) {
	log.Println("email methond chosen")
	// Set up authentication information.
	auth := smtp.PlainAuth("", emailUsername, emailPass, emailHost)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{client.email}
	body := "To: " + client.email + "\r\n" +
		"Subject: Purchase Info!\r\n" +
		"\r\n" +
		msg + "\r\n"

	// Send
	err := smtp.SendMail(emailHost+emailPort, auth, emailUsername, to, []byte(body))
	if err != nil {
		log.Fatal("Email method error: ", err)
		return "", err
	}
	return "Email Sent", nil
}

func phoneEmail(msg string, client *Buyer) (string, error) {
	resPhone, err := phone(msg, client)
	if err != nil {
		log.Printf("%v", err)
	}
	resEmail, err := email(msg, client)
	if err != nil {
		log.Printf("%v", err)
	}
	return fmt.Sprintf("%v\n%v", resPhone, resEmail), nil
}

// Notify notifies buyer
func (b *Buyer) Notify(msg, notify string) (string, error) {
	method := notifyMap[notify]
	return method(msg, b)
}
