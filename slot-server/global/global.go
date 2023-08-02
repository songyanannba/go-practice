package global

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
	"strings"
	"sync"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"slot-server/utils/timer"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/spf13/viper"
	"gorm.io/gorm"
	"slot-server/config"
)

var (
	GVA_DB      *gorm.DB // 读写库
	GVA_READ_DB *gorm.DB // 只读库 用于查询 不可用于写操作
	GVA_DBList  map[string]*gorm.DB
	GVA_REDIS   *redis.ClusterClient
	GVA_CONFIG  config.Server
	GVA_VP      *viper.Viper
	// GVA_LOG    *oplogging.Logger
	GVA_LOG                 *ZapLogger
	GVA_Timer               timer.Timer = timer.NewTimerTask()
	GVA_Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
	SvName     string

	Json = sonic.ConfigFastest
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GVA_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}

type ZapLogger struct {
	*zap.Logger
}

func (z ZapLogger) Skip(s int) *ZapLogger {
	z.Logger = z.WithOptions(zap.AddCallerSkip(s))
	return &z
}

func (z ZapLogger) Fatal(v ...interface{}) {
	z.WithOptions(zap.AddCallerSkip(1)).Fatal(fmt.Sprint(v...))
}

func (z ZapLogger) Fatalf(format string, v ...interface{}) {
	z.WithOptions(zap.AddCallerSkip(1)).Fatal(fmt.Sprintf(format, v...))
}

func (z ZapLogger) Println(v ...interface{}) {
	z.WithOptions(zap.AddCallerSkip(1)).Info(fmt.Sprint(v...))
}

func (z ZapLogger) Infof(format string, a ...any) {
	z.WithOptions(zap.AddCallerSkip(1)).Info(fmt.Sprintf(format, a...))
}

func GetListenUrl(port string) string {
	if !strings.Contains(port, ":") {
		return fmt.Sprintf("%s:%s", GVA_CONFIG.System.ListenIp, port)
	}
	return port
}

func GetConnectUrl(port string) string {
	if !strings.Contains(port, ":") {
		return fmt.Sprintf("%s:%s", GVA_CONFIG.System.ConnectIp, port)
	}
	return port
}
