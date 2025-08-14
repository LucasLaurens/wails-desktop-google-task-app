package errors

import "log"

type FatalError struct {
	Err  error
	Args []interface{}
}

func Fatal(msg string, fatalError FatalError) {
	if fatalError.Err != nil {
		log.Fatalf(
			msg,
			fatalError.Args[0],
			fatalError.Err,
		)
	}
}
