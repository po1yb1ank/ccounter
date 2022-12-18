package key

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/po1yb1ank/ccounter/pkg/handlers"
	"github.com/po1yb1ank/ccounter/pkg/logger"
	s "github.com/po1yb1ank/ccounter/pkg/storage"
)

// @BasePath /

// @Summary gets counter
// @Schemes
// @Description gets counter :key: by key
// @Param key path string true "Counter key"
// @Produce json
// @Success 200 {object} handlers.CounterResponse
// @failure 500 {object} handlers.CounterResponse
// @Router /{key} [post]
func Get(storage s.IStorage, logger logger.ILogger) func(*gin.Context) {
	return func(ctx *gin.Context) {
		target := ctx.Param(handlers.KEY_WILDCARD)
		value, err := storage.Current(ctx, target)

		if err != nil {
			if errors.Is(s.ErrorKeyNotFound, err) {
				err := storage.Set(ctx, target, 0)
				if err != nil {
					logger.Error(handlers.ErrorFailedToGet.Error())
					ctx.JSON(
						http.StatusNotFound,
						handlers.NewErrCounterResponse(handlers.ErrorFailedToGet.Error()),
					)
					return
				}
			}

			logger.Error(handlers.ErrorFailedToGet.Error())
			ctx.JSON(
				http.StatusInternalServerError,
				handlers.NewErrCounterResponse(handlers.ErrorFailedToGet.Error()),
			)
			return
		}

		ctx.JSON(http.StatusOK, handlers.NewOkGetCounterResponse(&target, &value))
	}
}
