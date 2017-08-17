package dry

import "log"

func Check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func Assert(ok bool, msg ...interface{}) {
	if !ok {
		out := []interface{}{"assertion failed"}
		out = append(out, msg...)
		log.Panicln(out...)
	}
}
