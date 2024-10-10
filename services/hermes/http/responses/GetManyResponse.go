package responses

import "UTCTSS-microservices/services/feedback/models"

// swagger:response GetManyResponse
type GetManyResponse struct {
	// in: body
	Body GetManyResponseBody `json:"body"`
}

type GetManyResponseBody struct {
	Limit  int               `json:"limit"`
	Offset int               `json:"offset"`
	Data   []models.Feedback `json:"data"`
}
