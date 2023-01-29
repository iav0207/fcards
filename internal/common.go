package internal

func Assert(condition bool, optionalMsg ...any) {
	if !condition {
		Log.Fatalln(optionalMsg...)
	}
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func FatalIf(err error, msg string) {
	if err != nil {
		Log.Fatalln(msg, err)
	}
}
