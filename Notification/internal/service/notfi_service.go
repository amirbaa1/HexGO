package service

import (
	"Notification/internal/core/model"
	"Notification/internal/helper"
	"encoding/json"
	"log"
)

func (s *Service) SendEmail(message string) (bool, error) {
	log.Printf("Sending email with message: %s", message)

	var emailMessage model.SendMessageEmail
	err := json.Unmarshal([]byte(message), &emailMessage)
	if err != nil {
		return false, err
	}

	configEm := model.EmailConfig{
		UserName:        "amir.2002.ba@gmail.com",
		Password:        "151D8F621B3D1D8B104A6D386DA5C664DCCE",
		Server:          "smtp.elasticemail.com",
		To:              emailMessage.Email,
		Subject:         "Create Account",
		From:            "amir.2002.ba@gmail.com",
		BodyHtml:        emailMessage.Message,
		BodyText:        emailMessage.Message,
		IsTransactional: true,
	}
	err = helper.SendElasticEmail(configEm)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return false, err
	}
	log.Println("----> Email sent to recipient:", emailMessage.Email)

	return true, nil
}
