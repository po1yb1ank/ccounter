package key

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/po1yb1ank/ccounter/pkg/handlers"
	"github.com/po1yb1ank/ccounter/pkg/logger"
	"github.com/po1yb1ank/ccounter/pkg/storage"
)

func Get(storage storage.IStorage, logger logger.ILogger) func(*gin.Context) {
	return func(c *gin.Context) {
		target := c.Param(handlers.KEY_WILDCARD)
		value, err := storage.Current(c, target)
		if err != nil {
			logger.Error(handlers.ErrorFailedToGet.Error())
			c.JSON(
				http.StatusInternalServerError,
				handlers.NewErrCounterResponse(handlers.ErrorFailedToGet.Error()),
			)
			return
		}

		c.JSON(http.StatusOK, handlers.NewOkGetCounterResponse(target, value))
	}
}
