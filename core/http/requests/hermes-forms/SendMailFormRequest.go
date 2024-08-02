package hermes_forms

type SendMailFormRequest struct {
	// required: true
	Code string `json:"code" binding:"required"`
	// required: true
	Recipients  []string          `json:"recipients" binding:"required"`
	PlaceHolder map[string]string `json:"placeholder" binding:"required"`
	Subject     string            `json:"subject"`
}

// swagger:parameters SendMail
type Payload struct {
	// required: true
	// in: body
	// type: object
	Body SendMailFormRequest
}
