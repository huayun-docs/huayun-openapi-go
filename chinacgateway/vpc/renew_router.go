package chinacgatewayvpc

import (
    "strconv"
)

type renewRouter struct {
    req        	map[string]string
}

func RenewRouter() *renewRouter {
    c := &renewRouter {
        req:			make(map[string]string),
    }

    c.req["Action"] = "RenewRouter"
    return c
}

func (c *renewRouter) GetRequest() map[string]string {
    return c.req
}

func (c *renewRouter) GetAction() string {
    return c.req["Action"]
}

func (c *renewRouter) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *renewRouter) GetRegion() string {
    return c.req["Region"]
}

func (c *renewRouter) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *renewRouter) GetId() string {
    return c.req["Id"]
}

func (c *renewRouter) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *renewRouter) GetPayType() string {
    return c.req["PayType"]
}

func (c *renewRouter) SetPayType(varPayType string) {
    c.req["PayType"] = varPayType
}

func (c *renewRouter) GetPeriod() int {
    nPeriod, _ := strconv.Atoi(c.req["Period"])
    return nPeriod
}

func (c *renewRouter) SetPeriod(varPeriod int) {
    c.req["Period"] = strconv.Itoa(varPeriod)
}


