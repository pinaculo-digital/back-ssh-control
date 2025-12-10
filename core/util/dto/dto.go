package util_dto

import (
	"math"
)

type QueryPaginationDto struct {
	Page  int `form:"page" json:"page" validate:"gte=1" example:"1"`
	Limit int `form:"limit" json:"limit" validate:"gte=1,lte=5000" example:"10"`
}
type ResponsePaginatedDto[T any] struct {
	QueryPaginationDto
	Total      int `json:"total"`
	TotalPages int `json:"totalPages"`
	Data       T   `json:"data"`
}

func (pagination *ResponsePaginatedDto[T]) SetTotalPages() {

	calc := float64(pagination.Total) / float64(pagination.Limit)
	pagination.TotalPages = int(math.Ceil(calc))

}

type Response struct {
	StatusCode int    `json:"statusCode" example:"200"`
	Message    string `json:"message,omitempty" example:"Tudo certo!!"`
	Data       any    `json:"data,omitempty"`
}
