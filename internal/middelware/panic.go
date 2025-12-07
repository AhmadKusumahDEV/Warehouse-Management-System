package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type respone struct {
	StatusCode int    `json:"status_code"`
	Data       string `json:"data"`
}

func HandlePanic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if msg := recover(); msg != nil {
				if msg != 77 {
					res := respone{
						StatusCode: http.StatusBadRequest,
						Data:       "wrong password",
					}
					ctx.Header("Content-Type", "application/json")
					ctx.JSON(http.StatusBadRequest, res)
				} else if msg == 55 {
					res := respone{
						StatusCode: 403,
						Data:       "kesalahan saat updated",
					}
					ctx.Header("Content-Type", "application/json")
					ctx.JSON(http.StatusBadRequest, res)
				} else if msg == 403 {
					res := respone{
						StatusCode: 403,
						Data:       "bad request format json",
					}
					ctx.Header("Content-Type", "application/json")
					ctx.JSON(http.StatusNotFound, res)
				} else {
					res := respone{
						StatusCode: http.StatusNotFound,
						Data:       "email not found",
					}
					ctx.Header("Content-Type", "application/json")
					ctx.JSON(http.StatusNotFound, res)
				}
			}
			ctx.Abort()
		}()
		ctx.Next()
	}
}
