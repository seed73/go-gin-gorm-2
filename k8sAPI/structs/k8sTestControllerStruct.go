package structs

type InParam struct {
	PodName string `json:"PodName" example:"pod이름소문자" binding:"required"`
}
