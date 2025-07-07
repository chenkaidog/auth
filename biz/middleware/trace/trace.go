package trace

import (
	"auth/biz/util/id_gen"
	"auth/biz/util/trace_info"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

const (
	headerKeyLogId = "X-Log-ID"
)

func New() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		logID := c.Request.Header.Get(headerKeyLogId)
		if logID == "" {
			logID = id_gen.NewID()
		}
		ctx = trace_info.WithLogId(ctx, logID)
		c.Next(ctx)
		c.Header(headerKeyLogId, logID)
	}
}
