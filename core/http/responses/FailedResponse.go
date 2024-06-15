package responses

type FailedResponseBody struct {
	Message string `json:"message"`
}

// A ValidationError is an error that is used when the required input fails validation.
// swagger:response validationError
type ValidationError struct {
	// The error message
	// in: body
	Payload Body
}

// swagger:operation GET /users UserManagement listUsers
//
// Lists all users.
//
// This will show all available users.
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// schemes:
// - http
// - https
// responses:
//   "200":
//     description: A list of users
//     schema:
//       "$ref": "#/responses/usersResponse"
// tags:
// - User Management
