package taskrunner

type Runner struct {
	Controller controllerChan
	Error      controllerChan
	Data       dataChan
	DataSize   int
	Longlived  bool
	Dispatcher fn
	Executor   fn
}

func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		Longlived:  longlived,
		DataSize:   size,
		Dispatcher: d,
		Executor:   e,
	}
}

func (runner *Runner) startDispatch() {
	defer func() {
		if !runner.Longlived {
			close(runner.Controller)
			close(runner.Data)
			close(runner.Error)
		}
	}()

	for {
		select {
		case c := <-runner.Controller:
			if c == READY_TO_DISPATCH {
				err := runner.Dispatcher(runner.Data)
				if err != nil {
					runner.Error <- CLOSABLE
				} else {
					runner.Controller <- READY_TO_EXECUTE
				}
			}
			if c == READY_TO_EXECUTE {
				err := runner.Executor(runner.Data)
				if err != nil {
					runner.Error <- CLOSABLE
				} else {
					runner.Controller <- READY_TO_DISPATCH
				}
			}
		case e := <-runner.Error:
			if e == CLOSABLE {
				return
			}
		default:

		}
	}
}

func (runner *Runner) StartAll() {
	runner.Controller <- READY_TO_DISPATCH
	runner.startDispatch()
}
