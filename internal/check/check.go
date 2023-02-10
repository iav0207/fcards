package check

import "github.com/iav0207/fcards/internal/out"

// Assert on a technical constraint (a failure means a bug / logical inconsistency).
func Assert(condition bool, optionalMsg ...any) {
	if !condition {
		out.Log.Panicln(optionalMsg...)
	}
}

// Check a condition and exit with a fatal error if the check fails.
// The difference with Assert is that a failure here means a misuse on the part of user
// or the data.
func Require(condition bool, fmt string, optionalArgs ...any) {
	if !condition {
		out.Log.Fatalf(fmt, optionalArgs...)
	}
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func FatalIf(err error, msg string) {
	if err != nil {
		out.Log.Fatalln(msg, err)
	}
}
