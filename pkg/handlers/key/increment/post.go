package increment

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
		current, err := storage.Increment(c, target)
		if err != nil {
			logger.Error(handlers.ErrorFailedToIncrement.Error())
			c.JSON(
				http.StatusInternalServerError,
				handlers.NewErrCounterResponse(handlers.ErrorFailedToIncrement.Error()),
			)
			return
		}

		go watcher.NotifyChange(target, current)

		c.JSON(http.StatusOK, handlers.NewOkGetCounterResponse(target, current))
	}
}
