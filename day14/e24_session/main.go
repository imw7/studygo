package main

// Session中间件开发
// 设计一个通用的Session服务，支持内存存储和redis存储

// session模块设计
// 1.本质上 k-v 系统，通过key进行增删改查
// 2.session可以存储在内存或者redis（2个版本）
// session接口设计
// 1.Set()
// 2.Get()
// 3.Del()
// 4.Save(): session存储，redis实现延迟加载
// SessionMgr接口设计
// 1.Init(): 初始化，加载redis地址
// 2.CreateSession(): 创建一个新的session
// 3.GetSession(): 通过sessionId获取对应的session对象

// MemorySession设计
// 1.定义MemorySession对象（字段：sessionId、存k-v的map、读写锁）
// 2.构造函数，为了获取对象
// 3.Set()
// 4.Get()
// 5.Del()
// 6.Save()
// MemorySessionMgr设计
// 1.定义MemorySessionMgr对象（字段：存放所有session的map，读写锁）
// 2.构造函数
// 3.Init(): 初始化，加载redis地址
// 4.CreateSession(): 创建一个新的session
// 5.GetSession(): 通过sessionId获取对应的session对象

// RedisSession设计
// 1.定义一个RedisSession对象（字段：sessionId，存k-v的map，读写锁，redis连接池，记录内存中map是否被修改的标记）
// 2.构造函数
// 3.Set()：将session存到内存中的map
// 4.Get()：取数据，实现延迟加载
// 5.Del()
// 6.Save()：将session存到redis
// RedisSessionMgr设计
// 1.定义RedisSessionMgr对象（字段：redis地址、redis密码、连接池、读写锁、大map）
// 2. 构造函数
// 3. Init()
// 4. CreateSession()
// 5. GetSession()
func main() {

}
