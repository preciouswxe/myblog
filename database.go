package database

import "database/sql"

// 定义一个初始化数据库的函数
func initDB() (err error) {
	dsn := "new_user:Wxe13867187633_@tcp(124.71.136.179:3306)/uwb?charset=utf8mb4&parseTime=True"

	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
