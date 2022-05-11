package opdb

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
import "gopkg.in/ini.v1"

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
}

//解析ini文件并反射到结构体中
func parserConfig(dbCfg *DBConfig) {
	cfg, err := ini.Load("./opdb/config.ini")
	if err != nil {
		log.Panicf("Fail to read file: %v\n", err)
	}
	dbCfg.Host = cfg.Section("DB").Key("host").String()
	dbCfg.Port, _ = cfg.Section("DB").Key("port").Int()
	dbCfg.User = cfg.Section("DB").Key("user").String()
	dbCfg.Pwd = cfg.Section("DB").Key("pwd").String()
	dbCfg.Database = cfg.Section("DB").Key("database").String()
}

//读取配置文件内容
func LoadDBConfig() *DBConfig {
	dbCfg := &DBConfig{}
	//解析ini文件并反射到结构体中
	parserConfig(dbCfg)
	return dbCfg
}

var DB *gorm.DB

//从配置文件中读取数据库的配置信息并连接数据库
func InitMySqlConn() (err error) {
	//读取配置文件内容
	dbCfg := LoadDBConfig()
	//拼凑连接数据库的语句
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbCfg.User, dbCfg.Pwd, dbCfg.Host, dbCfg.Port, dbCfg.Database)
	log.Println(connStr)
	//连接数据库
	DB, err = gorm.Open("mysql", connStr)
	if err != nil {
		return err
	}
	//检测数据库是否活跃
	return DB.DB().Ping()
}

//如果表不存在,通过传入的结构体建表
func InitModel() {
	DB.AutoMigrate(&UserInfo{})
	DB.AutoMigrate(&Video{})
	DB.AutoMigrate(&Favorite{})
	DB.AutoMigrate(&Comment{})
	DB.AutoMigrate(&Relation{})
	DB.Model(&Video{}).AddForeignKey("fk_vi_userinfo_id", "user_infos(id)", "RESTRICT", "RESTRICT")

	DB.Model(&Favorite{}).AddForeignKey("user_info_id", "user_infos(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Favorite{}).AddForeignKey("video_id", "videos(id)", "RESTRICT", "RESTRICT")

	DB.Model(&Comment{}).AddForeignKey("user_info_id", "user_infos(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Comment{}).AddForeignKey("video_id", "videos(id)", "RESTRICT", "RESTRICT")

	DB.Model(&Relation{}).AddForeignKey("user_info_id", "user_infos(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Relation{}).AddForeignKey("user_info_to_id", "user_infos(id)", "RESTRICT", "RESTRICT")
}
