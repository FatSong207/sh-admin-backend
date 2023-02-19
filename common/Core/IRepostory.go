package Core

import (
	"SH-admin/models"
	"SH-admin/models/common"
	"gorm.io/gorm"
)

type IRepostory[T models.Entity, TODto any] interface {
	GetById(key int64) (*T, error)
	GetByWhereStruct(t *T) (*T, error)
	GetListByWhereStruct(t *T) ([]T, error)
	GetOutDtoById(key int64) (*TODto, error)
	FindWithPager(searchDto common.PageInfo, db *gorm.DB, order string, dest *[]*TODto, bind *[]*T) (int64, error)
	GetAll() ([]T, error)
	Insert(t *T, skipHook bool) (err error, rowsAffected int64)
	InsertBatch(ts *[]*T, skipHook bool) (err error, rowsAffected int64)
	Update(t *T, data map[string]any, skipHook bool) (rowsAffected int64, err error)
	DeleteByKeys(keys []int) (int64, error)
	GetBySQL(sql string) (*T, error)
	GetListBySQL(sql string) ([]T, error)
}
