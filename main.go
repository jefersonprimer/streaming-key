package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()

	e.POST("/auth", fun(c echo.Context) error {
		log.Default().Println("Running Auth...")
		body := c.Request().body
		defer body.Close()

		fields, _ := io.ReadAll(body)
		fmt.Println(string(fields))
		
		return c.String(http.StatusOK, "WORKING")
	})

	e.GET("/healthcheck", fun(c echo.Context) error {
		return c.String(http.StatusOK, "WORKING")
	})

	e.Logger.Fatal(e.Start(":8000"))
}