package dto

type APIResponse[T any] struct {
	Success bool      `json:"success"`
	Data    T         `json:"data,omitempty"`
	Error   *APIError `json:"error,omitempty"`
	Meta    *Meta     `json:"meta,omitempty"`
}

type APIError struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	DebugMessage error  `json:"debug_message,omitempty"`
	Details      any    `json:"details,omitempty"`
}

type Meta struct {
	Timestamp int64  `json:"timestamp"`
	RequestID string `json:"request_id"`
}
