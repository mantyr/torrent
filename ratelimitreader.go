package motorrent

import (
	"io"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/time/rate"
)

type rateLimitedReader struct {
	l *rate.Limiter
	r io.Reader
}

func (me rateLimitedReader) Read(b []byte) (n int, err error) {
	if err := me.l.WaitN(context.Background(), 1); err != nil {
		panic(err)
	}
	n, err = me.r.Read(b)
	if !me.l.ReserveN(time.Now(), n-1).OK() {
		panic(n - 1)
	}
	return
}
