package Core

import (
	"SH-admin/models"
)

type BaseService[T models.Entity, TODto any] struct {
	//IService[T, TKey]
	baseRepo IRepostory[T, TODto] //供BaseApi調用
}

// NewBaseService 供BaseApi調用
func NewBaseService[T models.Entity, TODto any]() IService[T, TODto] {
	ins := &BaseService[T, TODto]{
		baseRepo: NewBaseRepostory[T, TODto](),
	}
	return ins
}

// GetById 根據主鍵獲取實體
func (b *BaseService[T, TODto]) GetById(key int64) (*T, error) {
	t, err := b.baseRepo.GetById(key)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// GetByWhereStruct 根據傳入的實體當作查詢條件
func (b *BaseService[T, TODto]) GetByWhereStruct(t *T) (*T, error) {
	t, err := b.baseRepo.GetByWhereStruct(t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// GetListByWhereStruct 根據傳入的實體當作查詢條件，查多筆
func (b *BaseService[T, TODto]) GetListByWhereStruct(t *T) ([]T, error) {
	ts, err := b.baseRepo.GetListByWhereStruct(t)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

// GetOutDtoById 根據Id獲取實體OutDto
func (b *BaseService[T, TODto]) GetOutDtoById(key int64) (*TODto, error) {
	t, err := b.baseRepo.GetOutDtoById(key)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// FindWithPager 列表分頁
func (b *BaseService[T, TODto]) FindWithPager(searchDto models.SearchDto[T]) (*[]*T, int64, error) {
	var query = searchDto.Entity
	var dest = make([]*T, 0)
	var bind = make([]*T, 0)

	t, err := b.baseRepo.FindWithPager(searchDto.PageInfo, query, &dest, &bind)
	if err != nil {
		return nil, 0, err
	}
	return &dest, t, nil
}

// GetAll 獲取所有
func (b *BaseService[T, TODto]) GetAll() ([]T, error) {
	all, err := b.baseRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return all, nil
}

// Insert 新增一個實體，可選是否跳過鉤子函數
func (b *BaseService[T, TODto]) Insert(t *T, skipHook bool) (err error, rowsAffected int64) {
	err, rowsAffected = b.baseRepo.Insert(t, skipHook)
	return err, rowsAffected
}

// InsertBatch 批量新增，可選是否跳過鉤子函數
func (b *BaseService[T, TODto]) InsertBatch(ts *[]*T, skipHook bool) (err error, rowsAffected int64) {
	err, rowsAffected = b.baseRepo.InsertBatch(ts, skipHook)
	return err, rowsAffected
}

// Update 修改
func (b *BaseService[T, TODto]) Update(t *T, data map[string]any, skipHook bool) (rowsAffected int64, err error) {
	rowsAffected, err = b.baseRepo.Update(t, data, skipHook)
	return rowsAffected, err
}
