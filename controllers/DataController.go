package controllers

import (
	"Beego-demo/models"
	"strconv"
)

type DataController struct {
	BaseController
	Mid int
}

func (c *DataController) Prepare() {
	c.BaseController.Prepare()
	midstr := c.Ctx.Input.Param(":mid")
	c.Data["mid"] = midstr
	mid, err := strconv.Atoi(midstr)
	if nil == err && mid > 0 {
		c.Mid = mid
	} else {
		c.setTpl()
	}
}

func (c *DataController) Index() {
	sj := models.MenuFormatStruct(c.Mid)
	if nil != sj {
		title := make(map[string]string)
		titlemap := sj.Get("schema")
		for k, _ := range titlemap.MustMap() {
			stype := titlemap.GetPath(k, "type").MustString()
			if "object" != stype && "array" != stype {
				if len(titlemap.GetPath(k, "title").MustString()) > 0 {
					title[k] = titlemap.GetPath(k, "title").MustString()
				} else {
					title[k] = k
				}
			}
		}
		c.Data["title"] = title
	}
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "data/footerjs.html"
	c.setTpl()
}
