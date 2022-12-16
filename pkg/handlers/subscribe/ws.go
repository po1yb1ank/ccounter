package subscribe

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/po1yb1ank/ccounter/pkg/handlers"
	"github.com/po1yb1ank/ccounter/pkg/logger"
	"github.com/po1yb1ank/ccounter/pkg/storage"
	"github.com/po1yb1ank/ccounter/pkg/watcher"
)

func Ws(storage storage.IStorage, logger logger.ILogger, watcher *watcher.Watcher) func(*gin.Context) {
	wsupgrader := websocket.Upgrader{
		ReadBufferSize:  handlers.DEFAULT_WS_BUFFER_SIZE,
		WriteBufferSize: handlers.DEFAULT_WS_BUFFER_SIZE,
	}

	return func(c *gin.Context) {
		conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to set websocket upgrade: %+v", err))
			c.JSON(http.StatusBadRequest, handlers.NewErrWSConnectionResponse())
			return
		}
		watcher.AddSubscriber(conn)
	}
}
