package api

import (
	"net/http"
	"time"

	db "github.com/aram88/simplebank/db/sqlc"
	"github.com/aram88/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password"  binding:"required,min=6"`
	FullName string `json:"full_name"  binding:"required"`
	Emial    string `json:"email"  binding:"required,email"`
}

type createUserRespons struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Emial,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if _, ok := err.(*pgconn.PgError); ok {
			ctx.JSON(http.StatusForbidden, errorResponce(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}

	rsp := createUserRespons{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt.Time,
		CreatedAt:         user.CreatedAt.Time,
	}
	ctx.JSON(http.StatusOK, rsp)
}
