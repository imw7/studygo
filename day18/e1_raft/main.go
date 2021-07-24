package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// 1.实现3节点选举
// 2.改造代码成分布式选举代码，加入RPC调用
// 3.演示完整代码，自动选举 日志复制

// 定义3节点常量
const raftCount = 3

// Leader 声明Leader对象
type Leader struct {
	Term     int // 任期
	LeaderId int // 编号
}

// Raft 声明Raft结构体
type Raft struct {
	mu              sync.Mutex // 锁
	me              int        // 节点编号
	currentTerm     int        // 当前任期
	votedFor        int        // 为哪个节点投票
	state           int        // 3个状态： 0 Follower 1 Candidate 2 Leader
	lastMessageTime int64      // 发送最后一条数据的时间
	currentLeader   int        // 设置当前节点的领导
	message         chan bool  // 节点间发信息的通道
	electCh         chan bool  // 选举通道
	heartBeatCh     chan bool  // 心跳信号的通道
	heartbeatReCh   chan bool  // 返回心跳信号的通道
	timeout         int        // 超时时间
}

// 0 还没上任	-1 没有编号
var leader = Leader{Term: 0, LeaderId: -1}

func main() {
	// 过程：有3个节点，最初都是follower
	// 若有candidate状态，进行投票拉票
	// 会产生Leader

	// 创建3个节点
	for i := 0; i < raftCount; i++ {
		// 创建3个Raft节点
		CreateNode(i)
	}
	select {} // 让main函数不结束
}

func CreateNode(me int) *Raft {
	raft := &Raft{}
	raft.me = me
	// -1代表谁都不投，此时节点刚创建
	raft.votedFor = -1
	// 0 Follower
	raft.state = 0
	raft.timeout = 0
	raft.currentLeader = -1
	// 节点任期
	raft.setTerm(0)

	raft.message = make(chan bool)
	raft.electCh = make(chan bool)
	raft.heartBeatCh = make(chan bool)
	raft.heartbeatReCh = make(chan bool)

	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 选举的协程
	go raft.election()

	// 心跳检查的协程
	go raft.sendLeaderHeartbeat()

	return raft
}

func (r *Raft) setTerm(term int) {
	r.currentTerm = term
}

func (r *Raft) election() {
	// 设置标记，判断是否选出了Leader
	var result bool
	for {
		// 设置超时，150到300(ms)的随机数
		timeout := randRange(150, 300)
		r.lastMessageTime = millisecond()
		select {
		// 延迟等待1毫秒
		case <-time.After(time.Duration(timeout) * time.Millisecond):
			fmt.Println("current state is:", r.state)
		}
		result = false
		for !result {
			// 选主逻辑
			result = r.electionOneRound(&leader)
		}
	}
}

// 随机值
func randRange(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

// 获取当前时间，发送最后一条数据的时间
func millisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// 实现选主的逻辑
func (r *Raft) electionOneRound(leader *Leader) bool {
	// 定义超时
	var timeout = int64(100)
	// 投票数量
	var voteCount int
	// 定义是否开始心跳信号的产生
	var isHeartbeat bool
	// 时间
	last := millisecond()
	// 用于返回值
	isSucc := false
	// 给当前节点变成Candidate
	r.mu.Lock()
	// 修改状态
	r.becomeCandidate()
	r.mu.Unlock()
	fmt.Println("start electing leader")
	for {
		// 遍历所有节点拉选票
		for i := 0; i < raftCount; i++ {
			if i != r.me {
				// 拉选票
				go func() {
					if leader.LeaderId < 0 {
						// 设置投票
						r.electCh <- true
					}
				}()
			}
		}
		// 设置投票数量，自己首先给自己投了1票
		voteCount = 1
		// 遍历
		for i := 0; i < raftCount; i++ {
			// 计算投票数量
			select {
			case ok := <-r.electCh:
				if ok {
					// 投票数量加1
					voteCount++
					// 若选票个数，大于节点数/2，则成功
					isSucc = voteCount > raftCount/2
					if isSucc && !isHeartbeat {
						// 变化成主节点，选主成功了
						// 开始触发心跳信号检测
						isHeartbeat = true
						r.mu.Lock()
						// 变主
						r.becomeLeader()
						r.mu.Unlock()
						// 由Leader向其他节点发送心跳信号
						r.heartBeatCh <- true
						fmt.Printf("no.%d node is leader now\n", r.me)
						fmt.Println("leader start send heartbeat signal")
					}
				}
			}
		}

		// 做最后校验工作
		// 若不超时，且票数大于一半，则选举成功，break
		if timeout+last < millisecond() || voteCount > raftCount/2 || r.currentLeader > -1 {
			break
		} else {
			// 等待操作
			select {
			case <-time.After(time.Duration(10) * time.Millisecond):
			}
		}
	}
	return isSucc
}

// 修改状态成为Candidate
func (r *Raft) becomeCandidate() {
	r.state = 1
	r.setTerm(r.currentTerm + 1)
	r.votedFor = r.me
	r.currentLeader = -1
}

// 修改状态成为Leader
func (r *Raft) becomeLeader() {
	r.state = 2
	r.currentLeader = r.me
}

// Leader节点发送心跳信号
// TODO 顺便完成数据同步
// TODO 看小弟挂没挂
func (r *Raft) sendLeaderHeartbeat() {
	// 死循环
	for {
		select {
		case <-r.heartBeatCh:
			r.sendAppendEntriesImpl()
		}
	}
}

// 用于返回给Leader的确认信号
func (r *Raft) sendAppendEntriesImpl() {
	// 是主就别跑了
	if r.currentLeader == r.me {
		// 此时是Leader
		var succCount = 0 // 记录确认信号的节点个数
		for i := 0; i < raftCount; i++ {
			if i != r.me {
				go func() {
					r.heartbeatReCh <- true
				}()
			}
		}
		// 计算返回确认信号个数
		for i := 0; i < raftCount; i++ {
			select {
			case ok := <-r.heartbeatReCh:
				if ok {
					succCount++
					if succCount > raftCount/2 {
						fmt.Println("election succeed, heartbeat signal ok")
						log.Fatal("program exit")
					}
				}
			}
		}
	}
}

// Param 首字母大写，RPC规范
// 分布式通信
type Param struct {
	Msg string
}

// Communication 通信方法
func (r *Raft) Communication(p Param, arg *bool) error {
	fmt.Println(p.Msg)
	*arg = true
	return nil
}
