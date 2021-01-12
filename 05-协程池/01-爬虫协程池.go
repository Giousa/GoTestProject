package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

//声明需要爬虫的对象
type Spider struct {
	url string
}

//爬虫开始请求
func (s *Spider) start() {
	fmt.Println("准备爬虫：",s.url)
	//time.Sleep(time.Second*10)
	s.spiderUrl()
}

//任务
type SpiderJob struct {
	Spider Spider
}

//任务队列
var SpiderJobQueue chan SpiderJob

//执行工作对象
type SpiderWorker struct {
	name string
	spiderWorkPool chan chan SpiderJob
	spiderJobChannel chan SpiderJob
	quit chan bool
}

//创建工人
func NewSpiderWorker(workPool chan chan SpiderJob, name string)  SpiderWorker{

	//这里workPool来源于任务管理员，所有工人都共用这同一个
	//但是，所有工人，都有各自的spiderJobChannel，简单来说，就是每个工人，都有各自的工作
	spiderWorker := SpiderWorker{
		name: name,
		spiderWorkPool: workPool,
		spiderJobChannel: make(chan SpiderJob),
		quit: make(chan bool),
	}

	fmt.Printf("新工人：【%s】 创建成功!\n",name)
	return spiderWorker
}


//工人 工作
func (w *SpiderWorker) startWork() {
	fmt.Printf("新工人：【%s】 开始工作::\n",w.name)
	go func() {
		for{
			//将工作ch，注册到工人ch中
			//w.spiderWorkPool  来源于 taskAdminer 任务管理员
			//每个工人，都会把自己的工作ch，放入工人ch中进行管理
			//注意：这里的<- w.spiderJobChannel 是写入下面的才是读取
			w.spiderWorkPool <- w.spiderJobChannel
			fmt.Printf("新工人：【%s】 将工作ch注册到工人ch中!\n",w.name)

			//监控工作
			select {
				case job := <- w.spiderJobChannel:
					fmt.Printf("【%s】工人: 接收到新任务\n",w.name)
					job.Spider.start()
				case <- w.quit:
					fmt.Printf("【%s】工人: 任务完成\n",w.name)
					return

			}
		}
	}()
}

func (w *SpiderWorker) stopWork() {
	go func() {
		w.quit <- true
	}()
}

//任务管理员,分发任务
type TaskAdminer struct {
	name string
	maxSpiderWorkers int//最大工人数
	spiderWorkPool chan chan SpiderJob//注册和工人数量一样的通道
}


//创建任务管理员
//仅创建一个任务管理员，多个工人，工人都使用同一个任务管理员下的spiderWorkPool
func NewTaskAdminer(maxWorks int) TaskAdminer {
	return TaskAdminer{
		name: "任务管理员",
		maxSpiderWorkers: maxWorks,
		spiderWorkPool: make(chan chan SpiderJob,maxWorks),//将工人放到一个池中,可以理解成一个部门中
	}
}

//任务管理员分配任务
func (t *TaskAdminer) dispatchTask() {
	fmt.Println("任务管理员")
	for  {
		select {
			case job := <- SpiderJobQueue:
				fmt.Println("任务管理员，接收到一个新任务")

				//匿名传参，防止影响外部job值
				go func(job SpiderJob) {
					//spiderWorkPool chan chan SpiderJob
					// <- t.spiderWorkPool  读取后 是 ：chan SpiderJob 还是一个通道，可以往这个通道里面写入

					//简单来说，就是讲任务管理员接收的任务，放入工人ch下的工作ch中
					//工人在创建，开始工作的时候，startWork，就已经将工作ch存入工人ch中
					//因为工人ch是有缓冲池的，所以是从里面选择了一个工作ch，将工作放进去，工人能读取到，就可以工作了
					spiderJobChannel := <- t.spiderWorkPool
					spiderJobChannel <- job

				}(job)
			default:
				//fmt.Println("ok !!")

		}
	}
}

//任务管理员开始运行
func (t *TaskAdminer) taskRun() {

	//创建对应工人数
	fmt.Println("开始创建工人::")
	for i := 0;i < t.maxSpiderWorkers;i++{
		workdName := "work_"+strconv.Itoa(i)
		spiderWorker := NewSpiderWorker(t.spiderWorkPool,workdName)
		spiderWorker.startWork()
	}

	//监控
	go t.dispatchTask()
	//t.dispatchTask()
}

//初始化配置
func initSpiderConfig()  {
	maxWorkers := 2
	maxQueues := 4
	taskAdminer := NewTaskAdminer(maxWorkers)
	fmt.Println("初始化，任务管理员")
	SpiderJobQueue = make(chan SpiderJob,maxQueues)
	fmt.Println("初始化，工作队列")

	//持续运行任务
	fmt.Println("开始运行任务::")
	taskAdminer.taskRun()
}


func main() {

	initSpiderConfig()

	//爬虫任务开始

	for i := 0; i < 20; i++ {
		s := Spider{
			url: "http://aa036.space/das/18/" + strconv.Itoa(i) + ".htm",
		}

		spiderJob := SpiderJob{
			Spider: s,
		}

		SpiderJobQueue <- spiderJob

		fmt.Println("任务:",i)

		time.Sleep(time.Second*3)
	}

	close(SpiderJobQueue)

	fmt.Println("---爬虫结束---")
}

func (s *Spider) spiderUrl() {
	requestSpiderUrl(s.url)
}

func requestSpiderUrl(url string)  {
	resp,_ := http.Get(url)

	defer resp.Body.Close()

	buf := make([]byte,1024)
	var result string
	for{
		n,err := resp.Body.Read(buf)
		if n == 0{
			break
		}

		if err != nil && err != io.EOF{
			fmt.Println("read err : ",err)
			break
		}

		result += string(buf[:n])
	}

	fmt.Println(result)
}
