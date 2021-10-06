package status

const (
	Success             = 200
	CREATED             = 201
	InternalServerError = 500
	NotFound            = 404
	MethodNotAllowed    = 405
	BadRequest          = 400
	ErrorUnauthorized   = 401
)

var statusMessage = map[int]string{
	Success:             "Success",
	CREATED:             "Data has been created",
	InternalServerError: "Internal server error",
	NotFound:            "Your request not found",
	MethodNotAllowed:    "Method not allowed",
	BadRequest:          "Please fill required %s",
	ErrorUnauthorized:   "User unauthorized",
}

func StatusText(code int) string {
	return statusMessage[code]
}
