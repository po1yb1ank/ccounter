package decrement

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/po1yb1ank/ccounter/pkg/handlers"
	"github.com/po1yb1ank/ccounter/pkg/logger"
	s "github.com/po1yb1ank/ccounter/pkg/storage"
	"github.com/po1yb1ank/ccounter/pkg/watcher"
)

// @BasePath /

// @Summary decrement counter
// @Schemes
// @Description decrement counter :key
// @Param key path string true "Counter key"
// @Produce json
// @Success 200 {object} handlers.CounterResponse
// @failure 404 {object} handlers.CounterResponse
// @failure 500 {object} handlers.CounterResponse
// @Router /{key}/decrement [post]
func Post(
	storage s.IStorage,
	logger logger.ILogger,
	watcher watcher.IWatcher,
) func(*gin.Context) {
	return func(c *gin.Context) {
		target := c.Param(handlers.KEY_WILDCARD)
		current, err := storage.Decrement(c, target)
		if errors.Is(err, s.ErrorKeyNotFound) {
			c.JSON(
				http.StatusNotFound,
				handlers.NewErrCounterResponse(err.Error()),
			)
			return
		}
		if err != nil {
			logger.Error(handlers.ErrorFailedToDecrement.Error())
			c.JSON(
				http.StatusInternalServerError,
				handlers.NewErrCounterResponse(handlers.ErrorFailedToDecrement.Error()),
			)
			return
		}

		err = watcher.NotifyChange(target, current)
		if err != nil {
			logger.Error(handlers.ErrorFailedToNotify.Error() + err.Error())
		}

		c.JSON(http.StatusOK, handlers.NewOkDecrementCounterResponse(&target, &current))
	}
}
