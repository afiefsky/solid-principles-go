package main

import "fmt"

// NotificationService is an abstraction for sending notifications
type NotificationService interface {
	Send(message string)
}

// EmailService sends notifications via email
type EmailService struct{}

func (e EmailService) Send(message string) {
	fmt.Println("Sending Email:", message)
}

// Notifier depends on the abstraction (NotificationService) instead of a specific implementation
type Notifier struct {
	service NotificationService
}

func (n Notifier) Notify(message string) {
	n.service.Send(message)
}

func main() {
	emailService := EmailService{}
	notifier := Notifier{service: emailService}

	notifier.Notify("Hello, Go!")
}
