package internal

func Assert(condition bool, optionalMsg ...string) {
	if !condition {
		Log.Fatalln(optionalMsg)
	}
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
