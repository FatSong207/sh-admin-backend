package Core

import (
	"SH-admin/global"
	"SH-admin/models"
	"SH-admin/models/common"
	"gorm.io/gorm"
)

type BaseRepostory[T models.Entity, TODto any] struct {
	//IRepostory[T, TKey]
}

// NewBaseRepostory 供BaseService調用
func NewBaseRepostory[T models.Entity, TODto any]() IRepostory[T, TODto] {
	instance := &BaseRepostory[T, TODto]{}
	return instance
}

// GetById 根據主鍵獲取實體
func (b *BaseRepostory[T, TODto]) GetById(key int64) (*T, error) {
	t := new(T)
	err := global.Db.Where("id=?", key).First(t).Error
	if err != nil {
		return nil, err
	}
	return t, err
}

// GetByWhereStruct 根據傳入的實體當作查詢條件
func (b *BaseRepostory[T, TODto]) GetByWhereStruct(t *T) (*T, error) {
	t2 := new(T)
	err := global.Db.Where(t).First(t2).Error
	if err != nil {
		return nil, err
	}
	return t2, nil
}

// GetListByWhereStruct 根據傳入的實體當作查詢條件，查多筆
func (b *BaseRepostory[T, TODto]) GetListByWhereStruct(t *T) ([]T, error) {
	t2 := make([]T, 0)
	err := global.Db.Where(t).Find(&t2).Error
	if err != nil {
		return nil, err
	}
	return t2, nil
}

// GetOutDtoById 根據Id獲取實體OutDto
func (b *BaseRepostory[T, TODto]) GetOutDtoById(key int64) (*TODto, error) {
	tOutDto := new(TODto)
	t := new(T)
	name := (*t).TableName()
	err := global.Db.Table(name).Where("id=?", key).First(tOutDto).Error
	if err != nil {
		return nil, err
	}
	return tOutDto, nil
}

func (b *BaseRepostory[T, TODto]) FindWithPager(searchDto common.PageInfo, query T, dest *[]*T, bind *[]*T) (int64, error) {
	limit := searchDto.PageSize
	offset := searchDto.PageSize * (searchDto.PageNum - 1)
	var t T
	name := t.TableName()
	global.Db.Offset(offset).Limit(limit).Table(name).Where(query).Find(dest)
	res := global.Db.Table(name).Where(query).Find(bind)
	return res.RowsAffected, res.Error
}

// GetAll 獲取所有
func (b *BaseRepostory[T, TODto]) GetAll() ([]T, error) {
	t := make([]T, 0)
	err := global.Db.Find(&t).Error
	if err != nil {
		return nil, err
	}
	return t, err
}

// Insert 新增一個實體，可選是否跳過鉤子函數
func (b *BaseRepostory[T, TODto]) Insert(t *T, skipHook bool) (err error, rowsAffected int64) {
	result := global.Db.Session(&gorm.Session{SkipHooks: skipHook}).Create(t)
	if result.Error != nil {
		panic(result.Error)
	}
	return result.Error, result.RowsAffected
}

// InsertBatch 批量新增，可選是否跳過鉤子函數
func (b *BaseRepostory[T, TODto]) InsertBatch(ts *[]*T, skipHook bool) (err error, rowsAffected int64) {
	result := global.Db.Session(&gorm.Session{SkipHooks: skipHook}).Create(ts)
	if result.Error != nil {
		panic(result.Error)
	}
	return result.Error, result.RowsAffected
}

// Update 修改
func (b *BaseRepostory[T, TODto]) Update(t *T, data map[string]any, skipHook bool) (rowsAffected int64, err error) {
	//s := structs.New(T)
	//result := global.Db.Debug().Omit("id", "updated").Model(TIDto).Updates(m)
	result := global.Db.Debug().Model(t).Session(&gorm.Session{SkipHooks: skipHook}).Updates(data)
	if result.Error != nil {
		panic(result.Error)
	}
	return result.RowsAffected, result.Error
}
