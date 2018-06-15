package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)

	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item %v", item)
		}

		for _, r := range result.Requests {
			e.Scheduler.Submit(r) // 产生了很多 requests, 但是只能往里面放一个，所以这里不能往下继续运行
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			// 循环等待
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			// 循环等待
			out <- result
		}
	}()
}
