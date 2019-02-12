package chinacgatewayvpc

import (
    "strconv"
)

type renewEips struct {
    req        	map[string]string
}

func RenewEips() *renewEips {
    c := &renewEips {
        req:			make(map[string]string),
    }

    c.req["Action"] = "RenewEips"
    return c
}

func (c *renewEips) GetRequest() map[string]string {
    return c.req
}

func (c *renewEips) GetAction() string {
    return c.req["Action"]
}

func (c *renewEips) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *renewEips) GetRegion() string {
    return c.req["Region"]
}

func (c *renewEips) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *renewEips) GetEip0() string {
    return c.req["Eip.0"]
}

func (c *renewEips) SetEip0(varEip0 string) {
    c.req["Eip.0"] = varEip0
}

func (c *renewEips) GetPayType() string {
    return c.req["PayType"]
}

func (c *renewEips) SetPayType(varPayType string) {
    c.req["PayType"] = varPayType
}

func (c *renewEips) GetPeriod() int {
    nPeriod, _ := strconv.Atoi(c.req["Period"])
    return nPeriod
}

func (c *renewEips) SetPeriod(varPeriod int) {
    c.req["Period"] = strconv.Itoa(varPeriod)
}


