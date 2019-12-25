package taskrunner

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	d := func(dataChan dataChan) error {
		for i := 0; i < 30; i++ {
			dataChan <- i
			log.Printf("Dispatcher sent: %v\n", i)
		}
		return nil
	}

	e := func(dataChan dataChan) error {
	forloop:
		for {
			select {
			case d := <-dataChan:
				log.Printf("Executor received: %v \n", d)
			default:
				break forloop
			}
		}
		return errors.New("Executor done , no executor")
	}

	runner := NewRunner(30, false, d, e)
	go runner.StartAll()
	//startAll 开启Dispatch 开始监听runner.Controller  ->  监听到dispatch  开始执行dispatcher
	//runner内部发起的dispatcher执行完了之后  channel  转为  executor
	//执行executor  executor 执行完了监听到channel关闭  通过抛出error
	//在runner 结束  dispatch
	time.Sleep(3 * time.Second)
}
