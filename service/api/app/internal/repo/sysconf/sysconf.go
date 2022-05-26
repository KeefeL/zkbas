package sysconf

import (
	"fmt"

	"github.com/zecrey-labs/zecrey-legend/pkg/multcache"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"gorm.io/gorm"
)

var (
	cacheZecreySysconfigIdPrefix   = "cache:zecrey:sysconfig:id:"
	cacheZecreySysconfigNamePrefix = "cache:zecrey:sysconfig:name:"
)

type sysconf struct {
	cachedConn sqlc.CachedConn
	table      string
	db         *gorm.DB
	cache      multcache.MultCache
}

/*
	Func: GetSysconfigByName
	Params: name string
	Return: info *Sysconfig, err error
	Description: get sysconfig by config name
*/
func (m *sysconf) GetSysconfigByName(name string) (config *SysconfInfo, err error) {
	dbTx := m.db.Table(m.table).Where("name = ?", name).Find(&config)
	if dbTx.Error != nil {
		err := fmt.Sprintf("[sysconfig.GetSysconfigByName] %s", dbTx.Error)
		logx.Error(err)
		return nil, dbTx.Error
	} else if dbTx.RowsAffected == 0 {
		err := fmt.Sprintf("[sysconfig.GetSysconfigByName] %s", ErrNotFound)
		logx.Error(err)
		return nil, ErrNotFound
	}
	return config, nil
}

/*
	Func: CreateSysconfig
	Params: config *Sysconfig
	Return: error
	Description: Insert New Sysconfig
*/

func (m *sysconf) CreateSysconfig(config *SysconfInfo) error {
	dbTx := m.db.Table(m.table).Create(config)
	if dbTx.Error != nil {
		logx.Error("[sysconfig.sysconfig] %s", dbTx.Error)
		return dbTx.Error
	}
	if dbTx.RowsAffected == 0 {
		logx.Error("[sysconfig.sysconfig] Create Invalid Sysconfig")
		return ErrInvalidSysconfig
	}
	return nil
}

func (m *sysconf) CreateSysconfigInBatches(configs []*SysconfInfo) (rowsAffected int64, err error) {
	dbTx := m.db.Table(m.table).CreateInBatches(configs, len(configs))
	if dbTx.Error != nil {
		logx.Error("[sysconfig.CreateSysconfigInBatches] %s", dbTx.Error)
		return 0, dbTx.Error
	}
	if dbTx.RowsAffected == 0 {
		logx.Error("[sysconfig.CreateSysconfigInBatches] Create Invalid Sysconfig Batches")
		return 0, ErrInvalidSysconfig
	}
	return dbTx.RowsAffected, nil
}

/*
	Func: UpdateSysconfigByName
	Params: config *Sysconfig
	Return: err error
	Description: update sysconfig by config name
*/
func (m *sysconf) UpdateSysconfig(config *SysconfInfo) error {
	dbTx := m.db.Table(m.table).Where("name = ?", config.Name).Select(NameColumn, ValueColumn, ValueTypeColumn, CommentColumn).
		Updates(config)
	if dbTx.Error != nil {
		err := fmt.Sprintf("[sysconfig.UpdateSysconfig] %s", dbTx.Error)
		logx.Error(err)
		return dbTx.Error
	} else if dbTx.RowsAffected == 0 {
		err := fmt.Sprintf("[sysconfig.UpdateSysconfig] %s", ErrNotFound)
		logx.Error(err)
		return ErrNotFound
	}
	return nil
}
