package wallet

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/NoNamePL/webWallet/database"
	"github.com/NoNamePL/webWallet/iternal/config"
	"github.com/NoNamePL/webWallet/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestUpdWallet(t *testing.T) {
	r := SetUpRouter()

	cfg, _ := config.NewConfig()

	db, _ := database.ConnectDB(cfg)

	h := &Handler{
		db: db,
	}
	r.POST("/wallet", h.UpdWallet)

	w := httptest.NewRecorder()

	// Create an example wallet for testing
	exampleWallet := models.Wallet{
		ValletId:      "cf936b56-f9e1-42d6-85ed-0e1322674053",
		OperationType: "DEPOSIT",
		Amount:        3000,
	}
	walletJson, _ := json.Marshal(exampleWallet)
	req, _ := http.NewRequest("POST", "/wallet", strings.NewReader(string(walletJson)))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// Compare the response body with the json data of walletJson
	assert.Equal(t, string(walletJson), w.Body.String())

}

func Testgetbalance(t *testing.T) {
	r := SetUpRouter()
	cfg, _ := config.NewConfig()

	db, _ := database.ConnectDB(cfg)

	h := &Handler{
		db: db,
	}
	r.GET("/wallets/:uuid", h.GetBalance)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("Get", "/wallets/cf936b56-f9e1-42d6-85ed-0e1322674053", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"Balance":"3000"}`, w.Body.String())
}
