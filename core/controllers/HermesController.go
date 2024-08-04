package controllers

import (
	"UTCTSS-microservices/common"
	hermes_forms "UTCTSS-microservices/core/http/requests/hermes-forms"
	"UTCTSS-microservices/core/models"
	"UTCTSS-microservices/core/repositories"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"text/template"
)

type HermesController struct {
	Repository repositories.MailTemplateRepository
}

// swagger:operation POST /hermes/send-mail Hermes SendMail
//
// # Send mail
//
// ---
// produces:
// - application/json
// consumes:
// - application/json
// responses:
//	 200:
//	  description: send email
//	  schema:
//	    type: object
//	    properties:
//	      message:
//	        type: string
//	 400:
//	  "$ref": "#/responses/ValidationError"

func (controller *HermesController) SendMail(context *gin.Context) {
	input := hermes_forms.SendMailFormRequest{}
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mailTemplate := models.MailTemplate{}
	err := controller.Repository.GetByCode(&mailTemplate, input.Code)
	if err != nil {
		context.JSON(common.ErrorHandlerHttpResponse(err))
		return
	}

	if input.Subject != "" {
		mailTemplate.Subject = input.Subject
	}

	message, err := controller.handleMailMessage(mailTemplate, input.Placeholder)

	if err != nil {
		context.JSON(common.ErrorHandlerHttpResponse(err))
		return
	}

	go sendMail(message, input.Recipients)

	context.JSON(http.StatusCreated, gin.H{"message": "Email sent successfully!"})
}

func (controller *HermesController) handleMailMessage(mailTemplate models.MailTemplate, placeholder map[string]string) (string, error) {
	for key, value := range placeholder {
		mailTemplate.Content = strings.Replace(mailTemplate.Content, "{{."+key+"}}", value, -1)
	}

	var message bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message.Write([]byte(fmt.Sprintf("Subject: "+mailTemplate.Subject+"\n%s\n\n", mimeHeaders)))

	mailTemplate1, _ := template.ParseFiles("storage/app/mail-layouts/" + mailTemplate.Layout + ".html")

	err := mailTemplate1.Execute(&message, struct {
		Content string
	}{Content: mailTemplate.Content})

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return message.String(), nil
}

func sendMail(message string, recipients []string) {
	mailHost := os.Getenv("MAIL_HOST")
	mailPort := os.Getenv("MAIL_PORT")
	sender := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")

	auth := smtp.PlainAuth("", sender, password, mailHost)
	err := smtp.SendMail(mailHost+":"+mailPort, auth, sender, recipients, []byte(message))
	if err != nil {
		fmt.Println(err)
		return
	}
}
