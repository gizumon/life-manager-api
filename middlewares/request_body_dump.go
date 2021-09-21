package middlewares

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func RequestBodyDump(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
}
