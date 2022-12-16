package reset

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/po1yb1ank/ccounter/pkg/handlers"
	"github.com/po1yb1ank/ccounter/pkg/logger"
	"github.com/po1yb1ank/ccounter/pkg/storage"
	"github.com/po1yb1ank/ccounter/pkg/watcher"
)

func Post(storage storage.IStorage, logger logger.ILogger, watcher *watcher.Watcher) func(*gin.Context) {
	return func(c *gin.Context) {
		target := c.Param(handlers.KEY_WILDCARD)
		if err := storage.Reset(c, target); err != nil {
			logger.Error(handlers.ErrorFailedToReset.Error())
			c.JSON(
				http.StatusInternalServerError,
				handlers.NewErrCounterResponse(handlers.ErrorFailedToReset.Error()),
			)
			return
		}
		go watcher.NotifyChange(target, 0)
		c.JSON(http.StatusOK, handlers.NewOkGetCounterResponse("test", 0))
	}
}
