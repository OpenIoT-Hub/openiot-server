package errno

const (
	Success = 200

	InvalidParams = 400

	Error                       = 500
	ErrorDatabaseQuery          = 50001
	ErrorDatabaseRecordNotFound = 50002
	ErrorUserNotExist           = 50003
)
