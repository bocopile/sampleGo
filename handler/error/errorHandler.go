package errors

type ErrorCode int

const (
	ErrCodeInvalidUserID ErrorCode = iota + 1
	ErrCodeUserNotFound
	ErrCodeInternalError
	// 추가적인 에러 코드를 여기에 정의
)

type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

var errorMessages = map[ErrorCode]string{
	ErrCodeInvalidUserID: "Invalid user ID format, it must be a numeric string",
	ErrCodeUserNotFound:  "The specified user was not found",
	ErrCodeInternalError: "An internal error occurred",
	// 추가적인 에러 메시지를 여기에 정의
}

func NewAppError(code ErrorCode) *AppError {
	message, exists := errorMessages[code]
	if !exists {
		message = "Unknown error"
	}
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// Error 메서드 추가
func (e *AppError) Error() string {
	return e.Message
}

// 특정 에러를 반환하는 헬퍼 함수
func NewUserNotFoundError() *AppError {
	return NewAppError(ErrCodeUserNotFound)
}
