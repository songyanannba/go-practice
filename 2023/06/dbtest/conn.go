package dbtest

import (
	"fmt"
	"github.com/gocql/gocql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var err error

func InitDB() {
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	//	logger.Config{
	//		SlowThreshold:             time.Second, // 慢 SQL 阈值
	//		LogLevel:                  logger.Warn, // 日志级别
	//		IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
	//		Colorful:                  true,        // 禁用彩色打印
	//	},
	//)

	var err error
	//dsn := "root:数据库密码@tcp(127.0.0.1:3306)/数据库名字?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"数据库密码",
		"127.0.0.1",
		3306,
		"orm_test",
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("数据库连接失败")
	}
	fmt.Println("conn suc ...")
}





//cassandra conn
var Session *gocql.Session


func InitCassandra() {

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "schema1"

	/*cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "user",
		Password: "password"
	}*/

	Session, err = cluster.CreateSession()
	if err != nil {
		panic(" cql InitCassandra err")
	}

	fmt.Println(" cql succ")
}
