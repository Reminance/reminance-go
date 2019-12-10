package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" //init()
)

var db *sql.DB

func initDB() (err error) {
	fmt.Println("开始连接数据库!")
	//连数据库
	dsn := "root:123456789@tcp(dev.seeu.ink:3306)/test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("dsn %s is valid, err:%v\n", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", dsn, err)
		return
	}
	//设置数据库连接池最大连接数
	db.SetMaxOpenConns(10)
	//设置数据库最大空闲连接数
	db.SetMaxIdleConns(5)
	return
}

type user struct {
	id       int
	username string
	password string
}

func queryOne(id int) {
	var u1 user
	//1.写查询单挑记录的mysql语句
	sqlStr := "select id, username, password from user where id = ?;"
	// //2.执行
	// rowObj := db.QueryRow(sqlStr, 1) //从连接池里拿一个链接出来查询单挑记录
	// //3.拿结果
	// rowObj.Scan(&u1.id, &u1.username, &u1.password)  //必须对rowObj对象调用scan 该方法会释放数据库连接
	db.QueryRow(sqlStr, 1).Scan(&u1.id, &u1.username, &u1.password)
	//4.打印结果
	fmt.Printf("u1:%v\n", u1)
}

func queryMulti() {
	//1.写查询单挑记录的mysql语句
	sqlStr := "select id, username, password from user;"
	// //2.执行
	// rowObj := db.QueryRow(sqlStr, 1) //从连接池里拿一个链接出来查询单挑记录
	// //3.拿结果
	// rowObj.Scan(&u1.id, &u1.username, &u1.password)  //必须对rowObj对象调用scan 该方法会释放数据库连接
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	//非常重要: 关闭rows释放持有的数据库连接
	defer rows.Close()

	//循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d username:%s password:%s\n", u.id, u.username, u.password)
	}
}

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(id, username, password) values (?,?,?)"
	ret, err := db.Exec(sqlStr, 3, "王五", "1111111")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set password=? where id = ?"
	ret, err := db.Exec(sqlStr, "222222", 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

//go 连接mysql示例
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	}
	fmt.Printf("连接数据库成功!\n")
	queryOne(1)
	queryMulti()
	insertRowDemo()
	time.Sleep(time.Duration(2) * time.Second)
	updateRowDemo()
	time.Sleep(time.Duration(2) * time.Second)
	deleteRowDemo()
}
