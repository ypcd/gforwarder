package gstunnellib

import "errors"

func ChanClose(c interface{}) {
	chan1, ok := c.(chan []byte)
	if !ok {
		CheckErrorEx_panic(errors.New("Error: The c is not chan []byte."))
	}
	defer Panic_Recover(g_logger)
	select {
	case _, ok = <-chan1:
		if ok {
			close(chan1)
		}
	default:
		close(chan1)
	}
}

func ChanClean(c chan []byte) {
	for range c {
	}
}
