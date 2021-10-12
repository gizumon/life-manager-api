package middlewares

import (
	"fmt"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go/v4"
	"github.com/labstack/echo/v4"
	"gizumon.com/life-manager-api/constants"
	"google.golang.org/api/option"
)

func InitFirebaseRDB() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			temp := os.Getenv(constants.FIREBASE_RDB_URL_KEY)
			fmt.Println(temp)
			ctx := c.Request().Context()
			conf := &firebase.Config{
				DatabaseURL: os.Getenv(constants.FIREBASE_RDB_URL_KEY),
			}
			// Fetch the service account key JSON file contents
			keyPath := filepath.Join(constants.BASE_PATH, os.Getenv(constants.GOOGLE_APPLICATION_CREDENTIALS_PATH_KEY))
			opt := option.WithCredentialsFile(keyPath)

			// Initialize the app with a service account, granting admin privileges
			app, err := firebase.NewApp(ctx, conf, opt)
			if err != nil {
				fmt.Errorf("Error initializing app:", err)
				return err
			}

			client, err := app.Database(ctx)
			if err != nil {
				fmt.Errorf("Error initializing database client:", err)
				return err
			}

			c.Set(constants.FIREBASE_RDB_CLIENT_KEY, client)
			return next(c)
		}
	}
}
