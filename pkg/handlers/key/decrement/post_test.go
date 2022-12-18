package decrement_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/po1yb1ank/ccounter/pkg/handlers"
	"github.com/po1yb1ank/ccounter/pkg/handlers/key/decrement"
	"github.com/po1yb1ank/ccounter/utils"
)

func TestDecrementNotFound(t *testing.T) {
	storage := utils.NewMockStorage()
	logger := &utils.MockLogger{}
	watcher := &utils.MockWatcher{}
	tests := []struct {
		name  string
		key   string
		value int64
	}{
		{
			"positive",
			"foo",
			5,
		},
		{
			"negative",
			"bar",
			-25,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Params = []gin.Param{{Key: handlers.KEY_WILDCARD, Value: tt.key}}

			decrement.Post(storage, logger, watcher)(ctx)

			assert.Equal(t, w.Code, http.StatusNotFound, "expected not found error")
		})
	}
}

func TestDecrement(t *testing.T) {
	storage := utils.NewMockStorage()
	logger := &utils.MockLogger{}
	watcher := &utils.MockWatcher{}
	tests := []struct {
		name  string
		key   string
		value int64
	}{
		{
			"positive",
			"foo",
			5,
		},
		{
			"negative",
			"bar",
			-25,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Params = []gin.Param{{Key: handlers.KEY_WILDCARD, Value: tt.key}}
			storage.Set(ctx, tt.key, tt.value)

			decrement.Post(storage, logger, watcher)(ctx)

			res := &handlers.CounterResponse{}
			json.Unmarshal(w.Body.Bytes(), res)

			assert.Equal(t, tt.value-1, *res.Value, "decrement failed")
		})
	}
}
