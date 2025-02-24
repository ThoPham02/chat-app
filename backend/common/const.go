package common

// common code
const (
	SUCCESS_CODE = 0
	SUCCESS_MESS = "Success"

	UNKNOW_CODE = 1
	UNKNOW_MESS = "Unknown error"

	DB_ERR_CODE = 2
	DB_ERR_MESS = "Database error"
)

// error code
const (
	USER_NOT_FOUND_CODE = 1001
	USER_NOT_FOUND_MESS = "User not found"

	PASSWORD_INCORRECT_CODE = 1002
	PASSWORD_INCORRECT_MESS = "Password is incorrect"

	USER_EXISTED_CODE = 1003
	USER_EXISTED_MESS = "User already existed"
)