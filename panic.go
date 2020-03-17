package evq

import "log"

func catchPanic(f func()) (err interface{}) {
	defer func() {
		err = recover()
		if err != nil {
			log.Fatalf("%s panic: %s", f, err)
		}
	}()

	f()
	return
}

//func RunPanicless(f func()) (panicless bool) {
//	defer func() {
//		err := recover()
//		panicless = err == nil
//		if err != nil {
//			log.Fatalf("%s panic: %s", f, err)
//		}
//	}()
//
//	f()
//	return
//}
//
//func RepeatUntilPanicless(f func()) {
//	for !RunPanicless(f) {
//	}
//}
