/** Copyright(C),2020-2022
* Author: Jaytan
* Date: 11/23/20 8:58 PM
* Description:  db addsbirouters
 */
package dbmgr

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
)

var DBGorm *gorm.DB

// var Redisdb *redis.Client

func DbInit() {

	// DataSourceName eg: user:password@(0.0.0.0:3306)/test
	dbName := fmt.Sprintf("%s:%s@(%s:%d)/%s",
		configure.SysConf.DBConf.UserName,
		configure.SysConf.DBConf.Pwd,
		configure.SysConf.DBConf.Ip,
		configure.SysConf.DBConf.Port,
		configure.SysConf.DBConf.DbName)

	var err error
	DBGorm, err = gorm.Open(configure.SysConf.DBConf.DriverName, dbName)
	if err != nil {
		rlogger.Trace(types.ModuleUdmDB, rlogger.ERROR, nil, "connect db failed")
		panic("connect db error " + err.Error())
	}

	if err = DBGorm.DB().Ping(); err != nil {
		rlogger.Trace(types.ModuleUdmDB, rlogger.ERROR, nil, "connect db failed")
		panic("connect db error " + err.Error())
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	// If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than
	// MaxIdleConns, then MaxIdleConns will be reduced to match the new MaxOpenConns limit.
	// If n <= 0, then there is no limit on the number of open connections. The default is 0 (unlimited).
	DBGorm.DB().SetMaxOpenConns(100)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// If MaxOpenConns is greater than 0 but less than the new MaxIdleConns,
	// then the new MaxIdleConns will be reduced to match the MaxOpenConns limit.
	// If n <= 0, no idle connections are retained.The default max idle connections is currently 2.
	DBGorm.DB().SetMaxIdleConns(20)
	DBGorm.SingularTable(true)

	// Set DBGorm debug
	if configure.SysConf.DBConf.Debug {
		DBGorm.LogMode(true)
	}

	//connect to redis server
	//addr := fmt.Sprintf("%s:%d", configure.UdmConf.REDIS.SerAddr.Ipv4, configure.UdmConf.REDIS.SerAddr.Port)
	//Redisdb = redis.NewClient(&redis.Options{
	//	Addr:     addr, // use default Addr
	//	Password: "",   // no password set
	//	DB:       0,    // use default DB
	//})

	//heart beat
	//pong, err := Redisdb.Ping().Result()
	//if err != nil {
	//	logger.Trace(MODULE_ID, types.ERROR, nil, "pong(%s), err(%s)", pong, err)
	//	panic("connect redis error " + err.Error())
	//}
	//
	//logger.Trace(MODULE_ID, types.INFO, nil, "Connect to redis server, addr(%s)", addr)
	//fmt.Printf("Connect to redis server, addr(%s) \n", addr)
	return
}
