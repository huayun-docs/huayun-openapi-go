package chinacgatewayecs

import (
    "strconv"
)

type renewInstance struct {
    req        	map[string]string
}

func RenewInstance() *renewInstance {
    c := &renewInstance {
        req:			make(map[string]string),
    }

    c.req["Action"] = "RenewInstance"
    return c
}

func (c *renewInstance) GetRequest() map[string]string {
    return c.req
}

func (c *renewInstance) GetRegion() string {
    return c.req["Region"]
}

func (c *renewInstance) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *renewInstance) GetAction() string {
    return c.req["Action"]
}

func (c *renewInstance) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *renewInstance) GetId() string {
    return c.req["Id"]
}

func (c *renewInstance) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *renewInstance) GetPayType() string {
    return c.req["PayType"]
}

func (c *renewInstance) SetPayType(varPayType string) {
    c.req["PayType"] = varPayType
}

func (c *renewInstance) GetPeriod() int {
    nPeriod, _ := strconv.Atoi(c.req["Period"])
    return nPeriod
}

func (c *renewInstance) SetPeriod(varPeriod int) {
    c.req["Period"] = strconv.Itoa(varPeriod)
}


