package handlers

import (
	"github.com/jsiebens/ionscale/internal/version"
	"github.com/labstack/echo/v4"
)

func IndexHandler(code int) echo.HandlerFunc {
	return func(c echo.Context) error {
		info, s := version.GetReleaseInfo()
		data := map[string]interface{}{
			"Version":  info,
			"Revision": s,
		}
		return c.Render(code, "index.html", data)
	}
}
