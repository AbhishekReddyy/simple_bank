package api

import (
	"os"
	mockdb "simplebank/db/mock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func mockForTest(t *testing.T) *mockdb.MockStore {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	return mockdb.NewMockStore(ctrl)
}
