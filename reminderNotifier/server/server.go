package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Reminder struct {
	Topic string    `json:"topic"`
	Time  time.Time `json:"Time"`
}

var reminders []Reminder

func main() {
	http.HandleFunc("/view", viewReminders)
	http.HandleFunc("/add", addReminders)
	fmt.Println("Server is listing the PORT number 8080")
	http.ListenAndServe(":8080", nil)
}
func viewReminders(w http.ResponseWriter, r *http.Request) {
	if len(reminders) == 0 {
		fmt.Println("No reminders", http.StatusNotFound)
		http.Error(w, "Noreminders", http.StatusNotFound)
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(reminders)
}
func addReminders(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "only post method is allowed", http.StatusMethodNotAllowed)
		return
	}
	var rem Reminder
	err := json.NewDecoder(r.Body).Decode(&rem)
	if err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
	}
	reminders = append(reminders, rem)
	err = send(rem)
	if err != nil {
		fmt.Println("failed to sent the Message", err)
	} else {
		fmt.Println("message sent successfully", http.StatusCreated)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Reminder Added successfully"))

}

func send(rem Reminder) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	msg := fmt.Sprintf("⏰ Reminder: %s at %s", rem.Topic, rem.Time.Format("02 Jan 2006 3:04 PM"))

	params := &openapi.CreateMessageParams{}
	params.SetTo(os.Getenv("TO_PHONE_NUMBER"))
	params.SetFrom(os.Getenv("TWILIO_FROM"))
	params.SetBody(msg)
	_, err := client.Api.CreateMessage(params)
	return err
}
