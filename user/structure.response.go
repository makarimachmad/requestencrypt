package user

type(
	Response struct{
		Error	bool		`json:"error"`
		Message	string		`json:"message"`
		Data	interface{}	`json:"data"`
		Erorrs	interface{}	`json:"errors"`
	}
)