package main

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

const (
	ishadowxSite = "https://get.ishadowx.net/"
)

type crawler struct {
	collector *colly.Collector
	configs   []config
}

func newCrawler() *crawler {
	return &crawler{
		collector: colly.NewCollector(),
		configs:   make([]config, 0),
	}
}

func getConfigs(c *crawler) []config {
	c.collector.OnHTML("div.hover-text", c.parse)
	err := c.collector.Visit(ishadowxSite)
	if err != nil {
		panic(err)
	}
	return c.configs
}

func (c *crawler) parse(e *colly.HTMLElement) {
	sp, err := strconv.Atoi(strings.Replace(strings.Replace(e.DOM.Find("h4").Eq(1).Text()[5:], " ", "", -1), "\n", "", -1))
	if err != nil {
		panic(err)
	}
	c.configs = append(c.configs, config{
		Enable:     true,
		Password:   strings.Replace(strings.Replace(e.DOM.Find("h4").Eq(2).Text()[9:], " ", "", -1), "\n", "", -1),
		Method:     strings.Replace(strings.Replace(e.DOM.Find("h4").Eq(3).Text()[7:], " ", "", -1), "\n", "", -1),
		Server:     strings.Replace(strings.Replace(e.DOM.Find("h4").Eq(0).Text()[11:], " ", "", -1), "\n", "", -1),
		Obfs:       "plain",
		Protocol:   "origin",
		ServerPort: sp,
	})
}
