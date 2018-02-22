package controllers

import "github.com/revel/revel"

type Plugin struct {
	*revel.Controller
}

func (c Plugin) Index() revel.Result {
	return c.Render()
}

func (c Plugin) Search() revel.Result {
	return c.Render()
}
