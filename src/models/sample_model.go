package models

type Sample struct {
	Id        *int      `json:"id,omitempty"`
	Data      string    `json:"data" validate:"required,min=20"`
}

// type UpdateSampleReq struct {
// 	Data string `json:"data" binding:"required"`
// }
