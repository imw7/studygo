package db

import (
	"book/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// var db *sql.DB

func Init() (err error) {
	dsn := "root:password@tcp(127.0.0.1:3306)/books"
	db, err = sqlx.Connect("mysql", dsn)
	// db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	// err = db.Ping()
	// if err != nil {
	// 	return
	// }
	db.SetMaxOpenConns(100) // 最大连接
	db.SetMaxIdleConns(16)  // 最大空闲
	return
}

func ShowAllBooks() (books []*model.Book, err error) {
	sqlStr := "select id, title, price from book"
	err = db.Select(&books, sqlStr)
	// rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("query failed, err:", err)
		return
	}
	// defer func() { _ = rows.Close() }()
	// for rows.Next() {
	//	var book Book
	//	err = rows.Scan(&book.ID, &book.Title, &book.Price)
	//	if err != nil {
	//		fmt.Println("scan failed, err:", err)
	//		return
	// 	}
	// 	fmt.Printf("id:%d title:%s price:%d\n", book.ID, book.Title, book.Price)
	// }
	return
}

func InsertBook(title string, price int64) (err error) {
	sqlStr := "insert into book(title, price) value (?, ?)"
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Println("insert failed, err:", err)
		return
	}
	return
}

func DeleteBook(id int64) (err error) {
	sqlStr := "delete from book where id=?"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("delete failed, err:", err)
		return
	}
	return
}
