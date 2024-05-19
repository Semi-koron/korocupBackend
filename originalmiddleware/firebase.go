package originalmiddleware

import (
	"net/http"
	"strings"

	"context"

	firebase "firebase.google.com/go/v4"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

// FirebaseInit はFirebaseプロジェクトを初期化します。
func FirebaseInit() (*firebase.App, error) {
	opt := option.WithCredentialsFile("/app/korocup2024-firebase-adminsdk-37xv8-f9c212121c.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	return app, nil
}

// FirebaseAuthMiddleware はFirebase AuthのJWTトークンを検証するミドルウェアです。
func FirebaseAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
		}

		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		app, err := FirebaseInit()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Firebase init error")
		}
		auth, err := app.Auth(context.Background())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Firebase auth error")
		}
		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		// トークンからUIDを取得し、リクエストコンテキストに追加
		c.Set("uid", token.UID)

		return next(c)
	}
}
