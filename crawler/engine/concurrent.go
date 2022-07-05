package engine

import "fmt"

type ConCurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan (chan Request)
	WorkReady(chan Request)
	Run()
}

//队列版
func (eg *ConCurrentEngine) Run(seeds ...Request) {
	out := make(chan ParserResult)
	eg.Scheduler.Run()

	for i := 0; i < eg.WorkerCount; i++ {
		CreateWorker(out,  eg.Scheduler)
	}

	for _, req := range seeds {
		eg.Scheduler.Submit(req)
	}

	for {
		result := <- out
		for _, item := range result.Items {
			//fmt.Printf("%s", item)
			if item == nil {

			}
		}

		for _, request := range result.Requests {
			fmt.Println(request.Url)
			eg.Scheduler.Submit(request)
		}
	}
}


func CreateWorker(out chan ParserResult, scheduler Scheduler)  {
	in := make(chan Request)
	go func() {
		for {
			scheduler.WorkReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}


//非队列版
func (eg *ConCurrentEngine) _Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParserResult)
	eg.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < eg.WorkerCount; i++ {
		//CreateWorker(in, out)
	}

	for _, req := range seeds {
		eg.Scheduler.Submit(req)
	}

	for {
		result := <- out
		for _, item := range result.Items {
			//fmt.Printf("%s", item)
			if item == nil {

			}
		}
		for _, request := range result.Requests {
			eg.Scheduler.Submit(request)
		}
	}
}
