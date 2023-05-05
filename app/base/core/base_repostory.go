package core

import (
	"SH-admin/app/models/common"
	"SH-admin/global"
	"gorm.io/gorm"
)

type BaseRepostory[T common.Entity, TODto any] struct {
	//IBaseRepostory[T, TKey]
}

// NewBaseRepostory 供BaseService調用
func NewBaseRepostory[T common.Entity, TODto any]() IBaseRepostory[T, TODto] {
	instance := &BaseRepostory[T, TODto]{}
	return instance
}

// GetById 根據主鍵獲取實體
func (b *BaseRepostory[T, TODto]) GetById(key int64) (*T, error) {
	t := new(T)
	err := global.DB().Where("id=?", key).First(t).Error
	if err != nil {
		return nil, err
	}
	return t, err
}

// GetByWhereStruct 根據傳入的實體當作查詢條件
func (b *BaseRepostory[T, TODto]) GetByWhereStruct(t *T) (*T, error) {
	t2 := new(T)
	err := global.DB().Where(t).First(t2).Error
	if err != nil {
		return nil, err
	}
	return t2, nil
}

// GetListByWhereStruct 根據傳入的實體當作查詢條件，查多筆
func (b *BaseRepostory[T, TODto]) GetListByWhereStruct(t *T) ([]T, error) {
	t2 := make([]T, 0)
	err := global.DB().Where(t).Find(&t2).Error
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
	err := global.DB().Table(name).Where("id=?", key).First(tOutDto).Error
	if err != nil {
		return nil, err
	}
	return tOutDto, nil
}

func (b *BaseRepostory[T, TODto]) FindWithPager(searchDto common.PageInfo, db *gorm.DB, order string, dest *[]*TODto, bind *[]*T) (int64, error) {
	limit := searchDto.PageSize
	offset := searchDto.PageSize * (searchDto.PageNum - 1)
	var t T
	name := t.TableName()
	//global.DB().Offset(offset).Limit(limit).Table(name).Where(query).Order(order).Find(dest)
	//res := global.DB().Table(name).Where(query).Find(bind)
	res := db.Table(name).Find(bind)
	total := res.RowsAffected
	db = db.Offset(offset).Limit(limit).Table(name).Order(order).Find(dest)
	return total, res.Error
}

// GetAll 獲取所有
func (b *BaseRepostory[T, TODto]) GetAll() ([]T, error) {
	t := make([]T, 0)
	err := global.DB().Find(&t).Error
	if err != nil {
		return nil, err
	}
	return t, err
}

// Insert 新增一個實體，可選是否跳過鉤子函數
func (b *BaseRepostory[T, TODto]) Insert(t *T, skipHook bool) (err error, rowsAffected int64) {
	result := global.DB().Session(&gorm.Session{SkipHooks: skipHook}).Create(t)
	if result.Error != nil {
		//panic(result.Error)
		return result.Error, 0
	}
	return nil, result.RowsAffected
}

// InsertBatch 批量新增，可選是否跳過鉤子函數
func (b *BaseRepostory[T, TODto]) InsertBatch(ts *[]*T, skipHook bool) (err error, rowsAffected int64) {
	result := global.DB().Session(&gorm.Session{SkipHooks: skipHook}).Create(ts)
	if result.Error != nil {
		//panic(result.Error)
		return result.Error, 0
	}
	return nil, result.RowsAffected
}

// Update 修改
func (b *BaseRepostory[T, TODto]) Update(t *T, data map[string]any, skipHook bool) (rowsAffected int64, err error) {
	//s := structs.New(T)
	//result := global.DB().Debug().Omit("id", "updated").Model(TIDto).Updates(m)
	result := global.DB().Debug().Model(t).Session(&gorm.Session{SkipHooks: skipHook}).Updates(data)
	if result.Error != nil {
		//panic(result.Error)
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// DeleteByKeys 根據主鍵批量刪除
func (b *BaseRepostory[T, TODto]) DeleteByKeys(keys []int) (int64, error) {
	var t = new(T)
	result := global.DB().Debug().Delete(t, keys)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// DeleteAll 刪除全部
func (b *BaseRepostory[T, TODto]) DeleteAll() (rowsAffected int64, err error) {
	var t = new(T)
	result := global.DB().Debug().Where("1 = 1").Delete(t)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (b *BaseRepostory[T, TODto]) GetBySQL(sql string) (*T, error) {
	var t = new(T)
	result := global.DB().Raw(sql).Scan(t)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return t, nil
}

func (b *BaseRepostory[T, TODto]) GetListBySQL(sql string) ([]T, error) {
	var t []T
	result := global.DB().Raw(sql).Scan(&t)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return t, nil
}
