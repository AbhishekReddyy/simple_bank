package api

import (
	"os"
	mockdb "simplebank/db/mock"
	db "simplebank/db/sqlc"
	"simplebank/util"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	config := &util.Config{
		AccessTokenDuration: time.Minute,
	}
	server := NewServer(*config, store)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func mockForTest(t *testing.T) *mockdb.MockStore {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	return mockdb.NewMockStore(ctrl)
}
