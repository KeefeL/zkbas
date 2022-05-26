package sysconf

import (
	"context"
	"sync"

	"github.com/zecrey-labs/zecrey-legend/pkg/multcache"
	"github.com/zecrey-labs/zecrey-legend/service/api/app/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Sysconf interface {
	GetSysconfigByName(name string) (info *SysconfInfo, err error)
	CreateSysconfig(config *SysconfInfo) error
	CreateSysconfigInBatches(configs []*SysconfInfo) (rowsAffected int64, err error)
	UpdateSysconfig(config *SysconfInfo) error
}

var singletonValue *sysconf
var once sync.Once

func New(c config.Config) Sysconf {
	once.Do(func() {
		conn := sqlx.NewSqlConn("postgres", c.Postgres.DataSource)
		gormPointer, err := gorm.Open(postgres.Open(c.Postgres.DataSource))
		if err != nil {
			logx.Errorf("gorm connect db error, err = %s", err.Error())
		}
		singletonValue = &sysconf{
			cachedConn: sqlc.NewConn(conn, c.CacheRedis),
			table:      `sys_config`,
			db:         gormPointer,
			cache:      multcache.NewGoCache(context.Background(), 100, 10),
		}
	})
	return singletonValue
}
