package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Reminder struct {
	Topic string
	Time  time.Time
}

var reminders []Reminder

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error While Loading the .env file")
		return
	}
	for {
		fmt.Println("\n Reminder App")
		fmt.Println("1. Add Reminder")
		fmt.Println("2. View Reminders")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			addReminder()
		case 2:
			viewReminders()
		case 3:
			fmt.Println("Exiting..................................")
			return
		default:
			fmt.Println("Invalid Option")
		}
	}
}
func addReminder() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the topic : ")
	topicInput, _ := reader.ReadString('\n')
	topicInput = strings.TrimSpace(topicInput)

	fmt.Println("Enter Time--->(January 22 2017 15:04) : ")
	timeinput, _ := reader.ReadString('\n')
	timeinput = strings.TrimSpace(timeinput)
	layout := "January 2 2006 15:04"

	reminderTime, err := time.ParseInLocation(layout, timeinput, time.Local)
	if err != nil {
		fmt.Println("Invalid Time and Date Format")
		return
	}
	duration := time.Until(reminderTime)
	if duration <= 0 {
		fmt.Println("Time is alredy Past")
		return
	}
	reminders = append(reminders, Reminder{Topic: topicInput, Time: reminderTime})
	fmt.Printf("Reminder set for %s\n", reminderTime.Format("22 Jan 2017 15:04"))

	go func(topic string, triggerTime time.Time) {
		time.Sleep(time.Until(triggerTime))
		message := fmt.Sprintf("Reminder: %s (at %s)", topic, triggerTime.Format("3:04 PM"))
		err := sendSms(message)
		if err != nil {
			fmt.Println("Failed to send the SMS", err)
		} else {
			fmt.Println("SMS SENT!!", message)
		}
	}(topicInput, reminderTime)
}
func viewReminders() {
	if len(reminders) == 0 {
		fmt.Println("There is No reminder set")
		return
	}
	for i, r := range reminders {
		fmt.Printf("%d. %s at %s\n", i+1, r.Topic, r.Time.Format("02 Jan 2006 15:04"))
	}
}
func sendSms(body string) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(os.Getenv("TO_PHONE_NUMBER"))
	params.SetFrom(os.Getenv("TWILIO_FROM"))
	params.SetBody(body)

	_, err := client.Api.CreateMessage(params)
	return err

}
