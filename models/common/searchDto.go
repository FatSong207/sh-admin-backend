package common

import "SH-admin/models"

type SearchDto[T models.Entity] struct {
	Entity    T //NOTE:這裡要放T，不能放Entity
	PageInfo  PageInfo
	OrderRule OrderRule
}

type OrderRule struct {
	OrderBy map[string]string
}

func NewSearchDto[T models.Entity]() *SearchDto[T] {
	ins := SearchDto[T]{
		Entity:   *new(T),
		PageInfo: PageInfo{},
		OrderRule: OrderRule{
			map[string]string{"created": "desc"},
		},
	}
	return &ins
}
