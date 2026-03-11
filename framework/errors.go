package framework

import "fmt"

// Param mewakili satu parameter path dinamis (misal: :id)
type Param struct {
	Key   string
	Value string
}

// Params adalah kumpulan dari Param
type Params []Param

// ByName mengambil nilai parameter berdasarkan key-nya.
func (ps Params) ByName(name string) string {
	for _, p := range ps {
		if p.Key == name {
			return p.Value
		}
	}
	return ""
}

// AppError adalah struktur kustom untuk error handling di framework.
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}

// NewAppError membuat instance AppError baru.
func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}