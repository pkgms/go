/**
 * Created by zc on 2020/10/27.
 */
package database

import (
	"github.com/go-sql-driver/mysql"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type Config struct {
	Addr   string `json:"addr" yaml:"addr"`
	User   string `json:"user" yaml:"user"`
	Pwd    string `json:"pwd" yaml:"pwd"`
	DBName string `json:"dbname" yaml:"dbname"`
	Debug  bool   `json:"debug" yaml:"debug"`
}

func (c *Config) Clone() *Config {
	return &Config{
		Addr:   c.Addr,
		User:   c.User,
		Pwd:    c.Pwd,
		DBName: c.DBName,
	}
}

func New(cfg *Config) (*gorm.DB, error) {
	config := mysql.Config{
		Addr:                 cfg.Addr,
		User:                 cfg.User,
		Passwd:               cfg.Pwd,
		DBName:               cfg.DBName,
		Net:                  "tcp",
		Collation:            "utf8mb4_general_ci",
		ParseTime:            true,
		Loc:                  time.UTC,
		AllowNativePasswords: true,
	}
	dbConnect, err := gorm.Open(
		gormmysql.Open(config.FormatDSN()),
		&gorm.Config{
			SkipDefaultTransaction: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	if cfg.Debug {
		dbConnect = dbConnect.Debug()
	}
	return dbConnect, nil
}
