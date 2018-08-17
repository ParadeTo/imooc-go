package engine

import "log"

type ConcurrentEngine struct {
	Scheduler      Scheduler
	WorkerCount    int
	ItemChan       chan Item
	RequestProcess Processor
	Deduplicate    Deduplicate
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)

	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		// url dedup
		if e.Deduplicate.IsDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		log.Printf("%+v", result)
		for _, item := range result.Items {
			go func(i Item) { e.ItemChan <- i }(item)
		}

		for _, r := range result.Requests {
			// url dedup
			if e.Deduplicate.IsDuplicate(r.Url) {
				continue
			}
			e.Scheduler.Submit(r)
		}
	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i am ready
			ready.WorkerReady(in)

			request := <-in
			result, err := e.RequestProcess(request) // call rpc
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
