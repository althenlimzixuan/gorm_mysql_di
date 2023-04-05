package gorm_mysql_di

import (
	"github.com/althenlimzixuan/gorm_mysql_di/config"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormMySqlServiceInt interface {
	AutoMigrate(interface{}) (bool, error)
	AutoMigrates(...interface{}) (bool, error)
	Ping() (bool, error)
	First(interface{}, ...interface{}) *gorm.DB
	GetDB() *gorm.DB
	Where(interface{}, ...interface{}) *gorm.DB
	Unscope() *gorm.DB
	Exec(string, ...interface{}) *gorm.DB
}

type GormMySqlService struct {
	Config *config.GormMySqlConfig
	DB     *gorm.DB
}

func ProvideGormMySqlService() GormMySqlServiceInt {
	cfg := config.GormMySqlConfig{}
	envconfig.Process(config.ConfigHeader, &cfg)

	db, err := gorm.Open(mysql.Open(cfg.Url), &gorm.Config{})

	if err != nil {
		logrus.Fatalln(err)
	}

	if !cfg.LogMode {
		db.Logger = logger.Default.LogMode(logger.Silent)
	}

	return &GormMySqlService{Config: &cfg, DB: db}
}

func (svc *GormMySqlService) AutoMigrate(entity_int interface{}) (bool, error) {

	err := svc.DB.AutoMigrate(entity_int)

	return err == nil, err
}

// Migrate Multiple Entities
func (svc *GormMySqlService) AutoMigrates(entities_int ...interface{}) (bool, error) {

	err := svc.DB.AutoMigrate(entities_int[:]...)

	return err == nil, err
}

func (svc *GormMySqlService) GetDB() *gorm.DB {
	return svc.DB
}

func (svc *GormMySqlService) Ping() (bool, error) {
	db, err := svc.DB.DB()

	if err != nil {
		logrus.Error(err)
		return false, err
	}

	err = db.Ping()

	if err != nil {
		logrus.Error(err)
		return false, err
	}

	return true, err
}

func (svc *GormMySqlService) First(dest interface{}, conds ...interface{}) *gorm.DB {
	svc.DB = svc.DB.First(dest, conds[:]...)
	return svc.DB
}

func (svc *GormMySqlService) Where(dest interface{}, conds ...interface{}) *gorm.DB {
	svc.DB = svc.DB.Where(dest, conds[:]...)
	return svc.DB
}

func (svc *GormMySqlService) Unscope() *gorm.DB {
	svc.DB = svc.DB.Unscoped()
	return svc.DB
}

func (svc *GormMySqlService) Exec(query string, values ...interface{}) *gorm.DB {
	svc.DB = svc.DB.Exec(query, values[:]...)
	return svc.DB
}
