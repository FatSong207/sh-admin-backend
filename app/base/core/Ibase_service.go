package core

import (
	"SH-admin/app/models/common"
)

type IBaseService[T common.Entity, TODto any] interface {
	GetById(key int64) (*T, error)
	GetByWhereStruct(t *T) (*T, error)
	GetListByWhereStruct(t *T) ([]T, error)
	GetOutDtoById(key int64) (*TODto, error)
	FindWithPager(searchDto common.SearchDto[T]) (*[]*TODto, int64, error)
	GetAll() ([]T, error)
	Insert(t *T, skipHook bool) (err error, rowsAffected int64)
	InsertBatch(ts *[]*T, skipHook bool) (err error, rowsAffected int64)
	Update(t *T, data map[string]any, skipHook bool) (rowsAffected int64, err error)
	DeleteByKeys(keys []int) (int64, error)
	DeleteAll() (rowsAffected int64, err error)
	GetBySQL(sql string) (*T, error)
	GetListBySQL(sql string) ([]T, error)
}
