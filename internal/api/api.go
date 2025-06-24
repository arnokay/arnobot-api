package api

import "github.com/labstack/echo/v4"

type Controllers struct {
	PlatformController *PlatformController
}

func (c *Controllers) Routes(parentGroup *echo.Group) {
	c.PlatformController.Routes(parentGroup)
}
