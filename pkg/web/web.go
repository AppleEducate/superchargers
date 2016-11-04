package web

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

var Schema graphql.Schema

func Run() error {
	var err error
	Schema, err = BuildSchema()
	if err != nil {
		return err
	}

	e := echo.New()
	e.Pre(redirectHTTPS)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339} method=${method} path=${path} host=${host} status=${status} bytes_in=${bytes_in} bytes_out=${bytes_out}\n",
		Skipper: func(c echo.Context) bool {
			return !strings.HasPrefix(c.Request().Host(), "localhost")
		},
	}))
	e.Use(middleware.Recover())
	e.Static("/", "public")
	e.Get("/.well-known/acme-challenge/:challenge", letsEncrypt)
	e.File("/graphiql", "public/graphiql.html")
	e.File("/faq", "public/faq.html")

	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	e.Any("/graphql", standard.WrapHandler(h))

	// Run the server
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))

	err = e.Run(standard.New(addr))
	return err
}

func letsEncrypt(c echo.Context) error {
	challenge := os.Getenv("LETS_ENCRYPT_CHALLENGE")
	param := c.Param("challenge")
	if param == challenge {
		return c.String(http.StatusOK, os.Getenv("LETS_ENCRYPT_KEY"))
	}

	return errors.New("Let's Encrypt challenge did not match")
}

func redirectHTTPS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.HasPrefix(c.Request().Host(), "localhost") {
			return next(c)
		}

		req := c.Request()
		host := req.Host()
		uri := req.URI()
		proto := req.Header().Get("X-Forwarded-Proto")
		fmt.Println(host, uri, req.Scheme(), proto)

		if proto != "https" {
			return c.Redirect(http.StatusMovedPermanently, "https://"+host+uri)
		}
		return next(c)
	}
}
