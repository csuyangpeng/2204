package router

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"

	"context"
	"fmt"
	"sync"
	"time"
)

type RouterMgr struct {
	//routeTbl  RouteTable  route table will be created on each worker goroutine
	ctrlChan CtrlChannel
	dataChan DataChannel
	ctx      context.Context
	cancel   context.CancelFunc

	numofWorker int
	dispatcher  *Dispatcher
}

func NewRouterMgr(appContext *types.AppContext) *RouterMgr {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)

	routerMgr := &RouterMgr{}
	routerMgr.ctrlChan = make(CtrlChannel, 10)
	routerMgr.dataChan = make(DataChannel, 10000) //TODO can be configrable
	routerMgr.numofWorker = 8
	/// for wait group:
	routerMgr.ctx, routerMgr.cancel = appContext.Ctx, appContext.Cancel

	routerMgr.dispatcher = NewDispatcher(routerMgr.numofWorker)

	return routerMgr
}

func (p *RouterMgr) String() (strbuf string) {
	strbuf += fmt.Sprintln("RouterMgr Info:")
	strbuf += fmt.Sprintln("ctrlChan: ", p.ctrlChan)
	strbuf += fmt.Sprintln("dataChan: ", p.dataChan)
	strbuf += fmt.Sprintln("ctx: ", p.ctx)
	strbuf += fmt.Sprintln("cancel: ", p.cancel)
	strbuf += fmt.Sprintln("numofWorker: ", p.numofWorker)
	strbuf += fmt.Sprintln("dispatcher: ", *p.dispatcher)
	return strbuf
}

func (p *RouterMgr) run(appContext *types.AppContext) {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)
	fmt.Println("router manager is running...")

	//start dispatch data message
	p.dispatcher.Run(p.dataChan)

	//start dispatch control message
	appWg := appContext.Wg
	go p.routeCtrlMsg(appWg, p.dispatcher, p.ctx)
}

// handle all incoming control message, notify to each workers and updated the route tables
func (p *RouterMgr) routeCtrlMsg(appWg *sync.WaitGroup, dispatcher *Dispatcher, ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)

	rlogger.Trace(types.ModuleCmnRouter, rlogger.INFO, nil, "Router Manager control Loop routine start")
	appWg.Add(1)
	defer func() {
		if p := recover(); p != nil {
			rlogger.Trace(types.ModuleCmnRouter, rlogger.ERROR, nil, "panics: %v", p)
		}
		appWg.Done()
		rlogger.Trace(types.ModuleCmnRouter, rlogger.ERROR, nil, "Router Manager control Loop routine exit")
	}()

	for {
		select {
		case <-ctxt.Done():
			return

		case msg := <-p.ctrlChan:
			switch msg.Op {
			case Register:
				err := routerRegister(msg.SrcAddr, msg.PubChannel)
				if err != nil {
					rlogger.Trace(types.ModuleCmnRouter, rlogger.ERROR, nil, "failed to regist channels")
				}
			case Deregister:
				err := routerDeregister(msg.SrcAddr)
				if err != nil {
					rlogger.Trace(types.ModuleCmnRouter, rlogger.ERROR, nil, "failed to deregist channels")
				}
			default:
				rlogger.Trace(types.ModuleCmnRouter, rlogger.ERROR, nil, "invalid control message to router")
			}

		}
	}
}

// display the route table
func DisplayRouteTable() {
	t2 := time.NewTicker(time.Second * 30)
	for {
		select {
		case <-t2.C:
			ShowRouteTable()
		}
	}
}

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan DataChannel
	//workerChannels []DataChannel
	maxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan DataChannel, maxWorkers)
	return &Dispatcher{
		WorkerPool: pool,
		maxWorkers: maxWorkers,
	}
}

func (d *Dispatcher) Run(dataChann DataChannel) {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)

	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		//store all the worker job channels
		//d.workerChannels = append(d.workerChannels, worker.JobChannel)
		// run all the worker goroutine
		worker.Start()
	}

	// dispatch all the data message to worker goroutine
	go d.dispatch(dataChann)
}

// notify to all worker for updating the route table on each worker
func (d *Dispatcher) Notify(msg *IpcMsg, jobChann DataChannel) {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)
	//for _, wchan := range d.workerChannels {
	//	rlogger.Trace(types.ModuleCmnRouter, rlogger.DEBUG, nil,  "dispatch notify message to worker goroutine")
	//	wchan <- msg
	//}
	//for i := 0; i < 2*d.maxWorkers; i++ {
	//	jobChann <- msg
	//}
}

func (d *Dispatcher) dispatch(rchann DataChannel) {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)
	for {
		select {
		case job := <-rchann:
			// a job request has been received
			go func(job *IpcMsg) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				if len(jobChannel) == 0 {
					//go func(job *IpcMsg) {
					//	jobChannel <- job
					//}(job)
					jobChannel <- job
				} else {
					fmt.Println("job channel is busy!, length = ", len(jobChannel))
				}

			}(job)
		}
	}
}

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan DataChannel
	JobChannel DataChannel
	//SignalChannel CtrlChannel
	quit chan bool
	//routeTbl RouteTable
}

func NewWorker(workerPool chan DataChannel) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(DataChannel),
		//SignalChannel: make(CtrlChannel,1000),
		quit: make(chan bool),
		//routeTbl: make(RouteTable),
	}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (p Worker) Start() {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)

	go func() {
		rlogger.Trace(types.ModuleCmnRouter, rlogger.DEBUG, nil, "worker goroutine created. id: %d.", utils.Goid())
		for {
			p.WorkerPool <- p.JobChannel

			select {
			//case w.WorkerPool <- w.JobChannel:
			// register the current worker into the worker queue.
			case job := <-p.JobChannel:
				switch job.MsgT {
				case DP:
					p.processIpcMsg(job)
				case CP:
					rlogger.Trace(types.ModuleCmnRouter, rlogger.DEBUG, nil, "worker goroutine id: %d, register message handled", utils.Goid())
					p.processCtrlMsg(job)
				default:
					rlogger.Trace(types.ModuleCmnRouter, rlogger.WARN, nil, "invalid message type")
				}
			case <-p.quit:
				// we have received a signal to stop
				return
			}
		}
		//}
	}()

	//go func() {
	//	t2 := time.NewTicker(time.Second)
	//	for {
	//		select {
	//		case <-t2.C:
	//			w.routeTbl.showRouteTable()
	//		}
	//	}
	//}()
}

// Stop signals the worker to stop listening for work requests.
func (p Worker) Stop() {
	go func() {
		p.quit <- true
	}()
}

func (p *Worker) processCtrlMsg(msg *IpcMsg) { //ControlMsg) {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)
	//switch msg.MsgT {
	//case CP:
	//	ctrlMsg := msg.MsgD.(*ControlMsg)
	//	switch ctrlMsg.Op {
	//	case Register:
	//		err := p.routeTbl.register(ctrlMsg.SrcAddr, ctrlMsg.PubChannel)
	//		if err != nil {
	//			rlogger.Trace(types.ModuleCmnRouter, rlogger.ERROR, nil,  "failed to regist public channel for %s", ctrlMsg.SrcAddr)
	//			return
	//		}
	//	case Deregister:
	//		p.routeTbl.deregister(ctrlMsg.SrcAddr)
	//		return
	//	default:
	//		rlogger.Trace(types.ModuleCmnRouter, rlogger.WARN, nil,  "invalid opertion for control ipc message")
	//	}
	//default:
	//	rlogger.Trace(types.ModuleCmnRouter, rlogger.WARN, nil,  "invalid ipc message, msgType is %d", msg.MsgT)
	//}
}

func (p *Worker) processIpcMsg(msg *IpcMsg) {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)

	switch msg.MsgT {
	case DP:
		p.dataMsgHandler(msg)
	default:
		rlogger.Trace(types.ModuleCmnRouter, rlogger.WARN, nil, "invalid ipc message, msgType is %d", msg.MsgT)
	}
}

func (p *Worker) dataMsgHandler(msg *IpcMsg) {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)

	dataMsg := msg.MsgD.(*DataMsg)

	//sendTime := time.Now()
	destChann, err := getDestChannel(&(dataMsg.DestAddr))
	//elapsed := time.Since(sendTime)
	//if elapsed > time.Microsecond*50 {
	//	fmt.Println("get data msg : ", elapsed)
	//}

	if destChann == nil || err != nil {
		rlogger.Trace(types.ModuleCmnRouter, rlogger.DEBUG, nil,
			"failed to find destination channel for %s", &(dataMsg.DestAddr))
		return
	}

	for _, ch := range destChann {
		ch <- msg
	}

	return
}
