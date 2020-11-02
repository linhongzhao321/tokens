package http

type HttpStat struct {
	code int
	msg  string
}

// status   > 0
// success = 0
// error   < 0
var (
	SUCCESS = HttpStat{
		0,
		"success",
	}

	PING = HttpStat{
		1,
		"ping",
	}

	ERR = HttpStat{
		-1,
		"not found err",
	}
)
