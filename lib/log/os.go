package log

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	"go.uber.org/zap"
)

type osSink struct {
	w io.Writer
}

func (o osSink) Write(p []byte) (n int, err error) {
	return o.w.Write(p)
}

func (o osSink) Sync() error {
	return nil
}

func (o osSink) Close() error {
	return nil
}

func registerOSSink() {
	if err := zap.RegisterSink("os", func(url *url.URL) (zap.Sink, error) {
		h := url.Host
		if h == "" {
			h = url.Opaque
		}
		// handle os:stdout and os:stderr
		switch strings.ToLower(h) {
		case "stdout":
			return osSink{w: os.Stdout}, nil
		case "stderr":
			return osSink{w: os.Stderr}, nil
		}
		return nil, fmt.Errorf("invalid os sink %q, expected \"os:stderr\" or \"os:stdout\"", url.String())
	}); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to register os log sink: %v\n", err)
	}
}
