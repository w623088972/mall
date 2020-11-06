package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
	"myself/mall/util"
	"strconv"
	"sync"
	"sync/atomic"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDb(),
	}
}

func (db *Database) Close() {
	DB.Self.Close()
}

func GetSelfDb() *gorm.DB {
	return InitSelfDb()
}

func InitSelfDb() *gorm.DB {
	db, err := openDb(
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"),
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
	)
	if err != nil {
		panic(err.Error())
		return nil
	}

	return db
}

type dbInfo struct {
	Host     string `gorm:"column:host"`
	Port     string `gorm:"column:port"`
	Database string `gorm:"column:database"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

type DbInstance struct {
	*gorm.DB
	lastUseTime int
	info        *dbInfo
	members     map[int]bool
}
type DbManager struct {
	dbs        map[string]*DbInstance
	platformDB *DbInstance
	*sync.Mutex
	num uint64
}

var (
	dbMgr = &DbManager{Mutex: new(sync.Mutex),
		dbs: make(map[string]*DbInstance, 0),
	}
)

func Init(host, port, database, username, password string, log bool) {
	db, err := openDb(host, port, database, username, password)
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(log)

	dbMgr.platformDB = &DbInstance{
		info: &dbInfo{
			Host:     host,
			Port:     port,
			Database: database,
			Username: username,
			Password: password,
		},
		DB: db,
	}
}

func openDb(host, port, database, username, password string) (*gorm.DB, error) {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host+port, database, "utf8", true, "Local")
	db, err := gorm.Open("mysql", config)
	if err != nil {
		fmt.Println("连接失败")
		return nil, err
	}

	setupDb(db)

	return db, nil
}

func setupDb(db *gorm.DB) {
	db.LogMode(viper.GetBool("mysql.log")) //true打印详细日志
	//db.SetLogger(gorm.Logger{gin.DefaultWriter})
	db.DB().SetMaxOpenConns(64) //用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(32) //用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

func GetDb(projectId int, userId int) (*DbInstance, error) {
	if projectId == 0 {
		return dbMgr.platformDB, nil
	}
	pid := fmt.Sprint(projectId)
	num := atomic.AddUint64(&dbMgr.num, 1)
	now := util.StrNow()
	log.Printf("%s GetDb for projectId:%d userId:%d getNum:%d \n", now, projectId, userId, num)
	defer func() {
		now := util.StrNow()
		log.Printf("%s GetDb for projectId:%d done getNum:%d\n", now, projectId, num)
	}()
	//TODO: 上锁的时间可能比较长(连接mysql)
	dbMgr.Lock()
	db, ok := dbMgr.dbs[strconv.Itoa(projectId)]
	if ok {
		dbMgr.Unlock()
		return db, nil
	}
	dbMgr.Unlock()
	/*
		if ok {
			_, ok := db.members[userId]
			if !ok {
				return nil, fmt.Errorf("projectId:%d accountId:%d does not match", projectId, userId)
			}
			return db, nil
		}
	*/

	tmp := struct {
		DbAddr     string `gorm:"column:db_addr"`
		DbName     string `gorm:"column:db_name"`
		DbUser     string `gorm:"column:db_user"`
		DbPassword string `gorm:"column:db_password"`
		TenantId   int    `gorm:"column:tenant_id"`
		AccountId  int    `gorm:"column:account_id"`
	}{}
	log.Printf("%s GetDb for projectId:%d userId:%d before query num:%d", now, projectId, userId, num)

	//TODO: 没必要联查tenant
	if err := dbMgr.platformDB.Raw(`SELECT tenant.account_id,tenant.id tenant_id,db_addr,db_name,db_user,db_password 
			FROM tenant_project INNER JOIN tenant ON(tenant_project.tenant_id = tenant.id)
			WHERE tenant_project.id = ?`, projectId).Scan(&tmp).Error; err != nil {
		return nil, err
	}
	log.Printf("GetDb %#v\n", tmp)
	//tmp.DbAddr = "vintop-online.mysql.polardb.rds.aliyuncs.com:3306"
	gormDB, err := openDb(tmp.DbAddr, "", tmp.DbName, tmp.DbUser, tmp.DbPassword)
	if err != nil {
		return nil, err
	}
	db = &DbInstance{
		DB: gormDB,
	}

	log.Printf("%s GetDb for projectId:%d userId:%d before lock num:%d", now, projectId, userId, num)
	dbMgr.Lock()
	defer dbMgr.Unlock()
	if db, ok := dbMgr.dbs[pid]; ok {
		return db, nil
	}
	dbMgr.dbs[pid] = db

	return db, nil
}
