package taskrunner

const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE  = "e"
	CLOSABLE          = "c"
)

type controllerChan chan string

type dataChan chan interface{}

type fn func(dataChan dataChan) error
