package models

type Sample struct {
	Id        int       `json:"id"`
	Data      string    `json:"data"`
}

// type CreateSampleReq struct {
// 	Data string `json:"data" binding:"required,min=20"`
// }

// type UpdateSampleReq struct {
// 	Data string `json:"data" binding:"required"`
// }
