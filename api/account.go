package api

import (
	"net/http"

	db "github.com/aram88/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency"  binding:"required,oneof=USD EUR"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}

type getAccountRequest struct {
	ID int64 `uri:"id", binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponce(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}
