package wallet

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/NoNamePL/webWallet/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type handler struct {
	db *sql.DB
}

func RegisterRouter(router *gin.Engine, db *sql.DB) {
	h := &handler{
		db: db,
	}

	routers := router.Group("/api/v1")
	routers.POST("/wallet", h.updWallet)
	routers.GET("/wallets/:uuid", h.getBalance)
}

func (db *handler) getBalance(ctx *gin.Context) {
	id := ctx.Param("uuid")

	stmt, err := db.db.Prepare("SELECT amount FROM wallet WHERE valletId = ($1)")

	if err != nil {
		wrongQuery(ctx)
		return
	}

	var resBalance string

	err = stmt.QueryRow(id).Scan(&resBalance)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("not row in db, err: %s", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"Status": "not row in db",
		})
		return
	}

	ctx.JSON(http.StatusOK, resBalance)

}

func (db *handler) updWallet(ctx *gin.Context) {
	var modelWallet models.Wallet
	if err := ctx.BindJSON(&modelWallet); err != nil {
		wrongQuery(ctx)
		return
	}

	// check to DIPOSIT or WITHDRAW
	if modelWallet.OperationType == "WITHDRAW" {
		modelWallet.Amount *= -1
	} else if modelWallet.OperationType != "DEPOSIT" {
		wrongQuery(ctx)
		return
	}

	// Check to for a value
	checkValleId, err := db.db.Prepare(`
		SELECT amount FROM wallet WHERE valletId = ($1)
	`)

	if err != nil {
		fmt.Println("can't unsert or update query:", err.Error())
		wrongQuery(ctx)
		return
	}

	var gotAmount int

	err = checkValleId.QueryRow(modelWallet.ValletId).Scan(&gotAmount)
	if errors.Is(err, sql.ErrNoRows) {
		// Insert data
		stmt, err := db.db.Prepare(`
			INSERT INTO wallet (valletId,amount)
			VALUES ($1, $2);
		`)
		if err != nil {
			wrongQuery(ctx)
			return
		}

		_, err = stmt.Exec(modelWallet.ValletId, modelWallet.Amount)

		if err != nil {
			wrongQuery(ctx)
			return
		}
	} else if err == nil {
		stmt, err := db.db.Prepare(`
			UPDATE wallet
			SET amount = $1
			WHERE valletId = $2;
		`)

		if err != nil {
			wrongQuery(ctx)
			return
		}

		_, err = stmt.Exec(modelWallet.Amount+gotAmount, modelWallet.ValletId)

		if err != nil {
			wrongQuery(ctx)
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "you are create/change wallet",
		"model":  modelWallet,
	})
}

func wrongQuery(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"Status": "BadRequest",
	})
}
