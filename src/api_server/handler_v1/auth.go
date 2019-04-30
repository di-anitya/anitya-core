package handler_v1

import (
	"bytes"
	"log"
	"net/http"

	"../graphql"
	"../middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Version() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "1.0\n")
	}
}

func Token() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		tokenString, refreshTokenString, err := middleware.CreateNewTokens(username, password)

		if err == nil {
			return c.JSON(http.StatusOK, map[string]string{
				"token":        tokenString,
				"refreshToken": refreshTokenString,
			})
		}

		return echo.ErrUnauthorized
	}
}

func User() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		_ = user.Claims.(jwt.MapClaims)
		bufBody := new(bytes.Buffer)
		bufBody.ReadFrom(c.Request().Body)
		query := bufBody.String()
		log.Printf(query)
		result := graphql.ExecuteQuery(query)
		return c.JSON(http.StatusOK, result)
	}
}

func ReAuth() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*middleware.MyClaim)
		oldToken := c.FormValue("old_token")
		tokenString, refreshTokenString, err := middleware.UpdateRefreshTokenExp(claims, oldToken)
		if err == nil {
			return c.JSON(http.StatusOK, map[string]string{
				"token":        tokenString,
				"refreshToken": refreshTokenString,
			})
		}
		return echo.ErrUnauthorized
	}
}
