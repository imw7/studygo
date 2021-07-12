package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 查询

// User
// 1.定义模型
type User struct {
	gorm.Model // ID CreateAt UpdateAt DeleteAt
	Name       string
	Age        int64
}

func main() {
	// 连接MySQL数据库
	dsn := "root:password@tcp(127.0.0.1:3306)/test?charset=utf8&mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 2.把模型与数据库中的表对应起来
	_ = db.AutoMigrate(&User{})

	// 3.创建（创建完成后注释掉，避免重复创建）
	// u1 := User{Name: "edward", Age: 19}
	// u2 := User{Name: "jessie", Age: 26}
	// u3 := User{Name: "bruce", Age: 32}
	// u4 := User{Name: "ada", Age: 20}
	// u5 := User{Name: "sarah", Age: 24}
	// u6 := User{Name: "jack", Age: 57}
	// u7 := User{Name: "peter", Age: 22}
	// u8 := User{Name: "mike", Age: 19}
	// u9 := User{Name: "elizabeth", Age: 93}
	// u10 := User{Name: "clark", Age: 33}
	// users := []User{u1, u2, u3, u4, u5, u6, u7, u8, u9, u10}
	// db.Create(&users)

	// 4.查询

	// // 一般查询
	// // user := new(User)
	// // db.First(user)
	// var user User // 声明模型结构体类型变量user
	// db.First(&user) // 根据主键查询第一条记录
	// fmt.Printf("user:%v\n", user)
	var users []User
	// db.Find(&users) // 查询所有的记录
	// fmt.Printf("users:%v\n", users)
	// db.Take(&user) // 获取一条记录
	// fmt.Printf("user:%v\n", user)
	// db.Last(&user) // 根据主键查询最后一条记录
	// fmt.Printf("user:%v\n", user)
	// db.First(&users, 3) // 查询指定的某条记录(仅当主键为整型时可用)
	// fmt.Printf("user:%v\n", users)

	// // Where条件
	// // 获取第一条匹配的记录
	// db.Where("name=?", "jessie").First(&user)
	// fmt.Printf("user:%v\n", user)
	// // 获取全部匹配的记录
	// db.Where("name=?", "jessie").Find(&users)
	// fmt.Printf("user:%v\n", users)
	// // 获取除了jessie以外的全部记录
	// db.Where("name <> ?", "jessie").Find(&users)
	// fmt.Printf("user:%v\n", users)
	// IN 获取包含的记录
	// db.Where("name IN (?)", []string{"jessie", "bruce", "john"}).Find(&users)
	// fmt.Printf("user:%v\n", users)
	// // LIKE 获取包含相关字段的记录
	// db.Where("name LIKE ?", "%ar%").Find(&users)
	// fmt.Printf("user:%v\n", users)
	// // AND
	// db.Where("name=? AND age>=?", "jessie", "22").Find(&users)
	// fmt.Printf("user:%v\n", users)
	// // Time
	// db.Where("updated_at>?", "2021-07-01 00:00:00").Find(&users)
	// fmt.Printf("user:%v\n", users)
	// // BETWEEN
	// db.Where("created_at BETWEEN ? AND ?", "2021-07-01 00:00:00", "2021-07-12 00:00:00")
	// fmt.Printf("user:%v\n", users)

	// // Struct & Map 查询
	// // Struct
	// db.Where(&User{Name: "jessie", Age: 26}).First(&user)
	// fmt.Printf("user:%v\n", user)
	// // Map
	// db.Where(map[string]interface{}{"name": "peter", "age": 22}).Find(&users)
	// fmt.Printf("user:%v\n", users)
	// // 主键的切片
	// db.Where([]int64{4, 5, 6}).Find(&users)
	// fmt.Printf("user:%v\n", users)

	// // Not条件
	// db.Not("name", "eric").First(&user)
	// fmt.Printf("user:%v\n", user)
	// // Not In
	// db.Not("name", []string{"jessie", "jack"}).Find(&users)
	// fmt.Printf("user:%v\n", users)
	// // Not In slice of primary keys
	// db.Not([]int64{1, 2, 3}).First(&user)
	// fmt.Printf("user:%v\n", user)
	// db.Not([]int64{}).First(&user)
	// fmt.Printf("user:%v\n", user)
	// db.Not("name = ?", "eric").First(&user)
	// fmt.Printf("user:%v\n", user)
	// db.Not(User{Name: "eric"}).First(&user)
	// fmt.Printf("user:%v\n", user)

	// // Or条件
	// db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
	// fmt.Printf("user:%v\n", users)
	// // Struct
	// db.Where("name = 'jessie'").Or(User{Name: "jack"}).Find(&users)
	// fmt.Printf("user:%v\n", users)
	// // Map
	// db.Where("name = 'jessie'").Or(map[string]interface{}{"name": "jack"}).Find(&users)
	// fmt.Printf("user:%v\n", users)

	// // 内联条件
	// // 根据主键获取记录 (只适用于整形主键)
	// db.First(&user, 5)
	// fmt.Println(user)
	// // 根据主键获取记录, 如果它是一个非整形主键
	// db.First(&user, "id = ?", "string_primary_key")
	// fmt.Println(user)
	// db.Find(&user, "name = ?", "jessie")
	// fmt.Println(user)
	// db.Find(&users, "name <> ? AND age > ?", "jessie", 20)
	// fmt.Println(users)
	// // Struct
	// db.Find(&users, User{Age: 20})
	// fmt.Println(users)
	// // Map
	// db.Find(&users, map[string]interface{}{"age": 20})
	// fmt.Println(users)

	// // 额外查询选项
	// // 为查询 SQL 添加额外的 SQL 操作
	// db.Set("gorm:query_option", "FOR UPDATE").First(&user, 10)

	// // FirstOrInit:获取匹配的第一条记录，否则根据给定的条件初始化一个新的对象 (仅支持 struct 和 map 条件)
	// // 未找到
	// db.FirstOrInit(&user, User{Name: "non_existing"})
	// fmt.Printf("user:%v\n", user)
	// // 找到
	// db.Where(User{Name: "jessie"}).FirstOrInit(&user)
	// fmt.Printf("user:%v\n", user)
	// db.FirstOrInit(&user, map[string]interface{}{"name": "jessie"})
	// fmt.Printf("user:%v\n", user)

	// // Attrs:如果记录未找到，将使用参数创建 struct 和记录
	// // 未找到
	// db.Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrCreate(&user)
	// // 找到
	// db.Where(User{Name: "jessie"}).Attrs(User{Age: 30}).FirstOrCreate(&user)

	// // Assign:不管记录是否找到，都将参数赋值给 struct 并保存至数据库
	// // 未找到
	// // db.Where(User{Name: "mike"}).Assign(User{Age: 20}).FirstOrCreate(&user)
	// // 找到
	// db.Where(User{Name: "jessie"}).Assign(User{Age: 30}).FirstOrCreate(&user)

	// // 高级查询

	// // 子查询
	// db.Where("amount > ?", db.Table("orders").Select("AVG(amount)").Where("state = ?", "paid").SubQuery()).Find(&orders)

	// // 选择字段:Select，指定你想从数据库中检索出的字段，默认会选择全部字段。
	// db.Select("name, age").Find(&users)
	// fmt.Printf("user:%v\n", users)
	// db.Select([]string{"name", "age"}).Find(&users)
	// fmt.Printf("user:%v\n", users)
	// _, err = db.Table("users").Select("COALESCE(age,?)", 20).Rows()
	// fmt.Printf("user:%v\n", users)

	// // 排序:Order，指定从数据库中检索出记录的顺序。设置第二个参数 reorder 为 true ，可以覆盖前面定义的排序条件。
	// db.Order("age desc, name").Find(&users)
	// fmt.Printf("%v\n", users)
	// // 多字段排序
	// db.Order("age desc").Order("name").Find(&users)
	// fmt.Printf("%v\n", users)
	// // 覆盖排序
	// db.Order("age desc").Find(&users1).Order("age", true).Find(&users2)

	// // 数量:Limit，指定从数据库检索出的最大记录数。
	// db.Limit(3).Find(&users)
	// fmt.Printf("%v\n", users)
	// // -1 取消 Limit 条件
	// db.Limit(5).Find(&users1).Limit(-1).Find(&users2)
	// fmt.Printf("%v\n", users)

	// // 偏移:Offset，指定开始返回记录前要跳过的记录数。
	// db.Limit(3).Offset(3).Find(&users)
	// fmt.Printf("%v\n", users)
	// db.Limit(6).Offset(5).Find(&users).Offset(-1).Find(&users)
	// fmt.Printf("%v\n", users)

	// 总数:Count，该 model 能获取的记录总数。
	var count int64
	db.Model(&User{}).Where("name = ?", "jessie").Or("name = ?", "jackie").Find(&users).Count(&count)
	fmt.Printf("%v\n", users)
}
