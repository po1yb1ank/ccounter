package reset

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/po1yb1ank/ccounter/pkg/handlers"
	"github.com/po1yb1ank/ccounter/pkg/logger"
	"github.com/po1yb1ank/ccounter/pkg/storage"
	"github.com/po1yb1ank/ccounter/pkg/watcher"
)

// @BasePath /

// @Summary resets counter
// @Schemes
// @Description resets counter :key
// @Param key path string true "Counter key"
// @Produce json
// @Success 200 {object} handlers.CounterResponse
// @failure 500 {object} handlers.CounterResponse
// @Router /{key}/reset [post]
func Post(
	storage storage.IStorage,
	logger logger.ILogger,
	watcher watcher.IWatcher,
) func(*gin.Context) {
	return func(c *gin.Context) {
		target := c.Param(handlers.KEY_WILDCARD)
		if err := storage.Set(c, target, 0); err != nil {
			logger.Error(handlers.ErrorFailedToReset.Error())
			c.JSON(
				http.StatusInternalServerError,
				handlers.NewErrCounterResponse(handlers.ErrorFailedToReset.Error()),
			)
			return
		}
		var value int64 = 0
		err := watcher.NotifyChange(target, value)
		if err != nil {
			logger.Error(handlers.ErrorFailedToNotify.Error() + err.Error())
		}

		c.JSON(http.StatusOK, handlers.NewOkResetCounterResponse(&target, &value))
	}
}
