package key_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/po1yb1ank/ccounter/pkg/handlers"
	"github.com/po1yb1ank/ccounter/pkg/handlers/key"
	"github.com/po1yb1ank/ccounter/utils"
)

func TestGet(t *testing.T) {
	storage := utils.NewMockStorage()
	logger := &utils.MockLogger{}
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

			key.Get(storage, logger)(ctx)

			res := &handlers.CounterResponse{}
			json.Unmarshal(w.Body.Bytes(), res)

			assert.Equal(t, tt.value, *res.Value, "get failed")
		})
	}
}

func TestGetNotSet(t *testing.T) {
	storage := utils.NewMockStorage()
	logger := &utils.MockLogger{}
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

			key.Get(storage, logger)(ctx)

			res := handlers.CounterResponse{}
			json.Unmarshal(w.Body.Bytes(), &res)

			assert.Equal(t, int64(0), *res.Value, "get when not set failed")
		})
	}
}
