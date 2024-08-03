package hermes_forms

type SendMailFormRequest struct {
	// required: true
	// type: string
	Code string `json:"code" binding:"required"`
	// required: true
	// type: array
	// items:
	//   type: string
	// example: ["email1@gmal.com", "email2@gmail.com"]
	Recipients []string `json:"recipients" binding:"required"`
	// required: false
	// type: object
	Placeholder map[string]string `json:"placeholder" binding:"required"`
	// required: false
	// type: string
	// extensions:
	//   x-note: All key must be capital the first character
	Subject string `json:"subject"`
}

// swagger:parameters SendMail
type Payload struct {
	// required: true
	// in: body
	// type: object
	Body SendMailFormRequest
}
