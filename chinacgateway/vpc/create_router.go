package chinacgatewayvpc

import (
    "strconv"
)

type createRouter struct {
    req        	map[string]string
}

func CreateRouter() *createRouter {
    c := &createRouter {
        req:			make(map[string]string),
    }

    c.req["Action"] = "CreateRouter"
    return c
}

func (c *createRouter) GetRequest() map[string]string {
    return c.req
}

func (c *createRouter) GetAction() string {
    return c.req["Action"]
}

func (c *createRouter) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *createRouter) GetRegion() string {
    return c.req["Region"]
}

func (c *createRouter) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *createRouter) GetPayType() string {
    return c.req["PayType"]
}

func (c *createRouter) SetPayType(varPayType string) {
    c.req["PayType"] = varPayType
}

func (c *createRouter) GetPeriod() int {
    nPeriod, _ := strconv.Atoi(c.req["Period"])
    return nPeriod
}

func (c *createRouter) SetPeriod(varPeriod int) {
    c.req["Period"] = strconv.Itoa(varPeriod)
}

func (c *createRouter) GetName() string {
    return c.req["Name"]
}

func (c *createRouter) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *createRouter) GetDescription() string {
    return c.req["Description"]
}

func (c *createRouter) SetDescription(varDescription string) {
    c.req["Description"] = varDescription
}

func (c *createRouter) GetFirewallId() string {
    return c.req["FirewallId"]
}

func (c *createRouter) SetFirewallId(varFirewallId string) {
    c.req["FirewallId"] = varFirewallId
}

func (c *createRouter) GetInterface0NetworkId() string {
    return c.req["Interface.0.NetworkId"]
}

func (c *createRouter) SetInterface0NetworkId(varInterface0NetworkId string) {
    c.req["Interface.0.NetworkId"] = varInterface0NetworkId
}


