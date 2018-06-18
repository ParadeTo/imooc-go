package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)

	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v", itemCount, item)
			itemCount++
		}

		for _, r := range result.Requests {
			e.Scheduler.Submit(r) // 产生了很多 requests, 但是只能往里面放一个，所以这里不能往下继续运行
		}
	}
}

func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			// tell scheduler i am ready
			s.WorkerReady(in)

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
