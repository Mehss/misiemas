package utils

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"tripatra-dct-service-config/database/model"
	userModel "tripatra-dct-service-config/database/model/user"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

// SendEmail handles sending an email using SMTP
func SendEmail(mail model.MailModel) (*model.Response, error) {
	response := model.Response{}

	// Load environment variables (SMTP configuration is here now)
	godotenv.Load()
	SMTP_HOST := os.Getenv("SMTP_HOST")
	SMTP_PORT, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	SMTP_NAME := os.Getenv("SMTP_NAME")
	SMTP_USER := os.Getenv("SMTP_USER")
	SMTP_PASS := os.Getenv("SMTP_PASS")

	// Create a new mail message
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", fmt.Sprintf("%s<%s>", SMTP_NAME, SMTP_USER))
	mailer.SetHeader("To", mail.To...)

	// Add CC recipients if present
	if len(mail.CC) > 0 {
		mailer.SetHeader("Cc", mail.CC...)
	}

	// Add BCC recipients if present
	if len(mail.BCC) > 0 {
		mailer.SetHeader("Bcc", mail.BCC...)
	}

	// Set the subject and body of the email
	mailer.SetHeader("Subject", mail.Subject)
	mailer.SetBody("text/html", mail.Body)

	// Attach files if present
	if len(mail.Attachments) > 0 {
		for _, attachment := range mail.Attachments {
			fileContent, errFile := base64.StdEncoding.DecodeString(attachment.Attachment)
			if errFile != nil {
				log.Println("Error decoding attachment:", errFile)
				continue
			}
			mailer.Attach(attachment.FileName, gomail.SetCopyFunc(func(w io.Writer) error {
				_, err := w.Write(fileContent)
				return err
			}))
		}
	}

	// Create a new SMTP dialer
	dialer := gomail.NewDialer(
		SMTP_HOST,
		SMTP_PORT,
		SMTP_USER,
		SMTP_PASS,
	)

	// Send the email
	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Println("Error sending email:", err)
		response.Status = false
		response.Message = err.Error()
		return &response, err
	}

	// Email sent successfully
	response.Status = true
	response.Message = "Mail sent!"
	log.Println("Mail sent successfully!")
	return &response, nil
}

func GetUserEmailByRoleConstant(ctx context.Context, tx *gorm.DB, roleIDS []uint) ([]userModel.User, error) {
	var listUser []userModel.User

	// Use WithContext to associate the query with the request context and use IN clause for multiple role IDs
	err := tx.WithContext(ctx).Where("role_id IN (?)", roleIDS).Find(&listUser).Error
	if err != nil {
		return listUser, err // Return the user list even if it's empty, along with the error
	}

	return listUser, nil
}

func GetEmailByUserID(ctx context.Context, tx *gorm.DB, userID uint) (string, error) {
	var email string

	// Use WithContext to associate the query with the request context
	err := tx.WithContext(ctx).Model(&userModel.User{}).Where("id = ?", userID).Pluck("email", &email).Error
	if err != nil {
		return email, err
	}

	return email, nil
}
