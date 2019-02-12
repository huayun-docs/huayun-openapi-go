package chinacgatewayvpc

import (
)

type assignRouterEip struct {
    req        	map[string]string
}

func AssignRouterEip() *assignRouterEip {
    c := &assignRouterEip {
        req:			make(map[string]string),
    }

    c.req["Action"] = "AssignRouterEip"
    return c
}

func (c *assignRouterEip) GetRequest() map[string]string {
    return c.req
}

func (c *assignRouterEip) GetAction() string {
    return c.req["Action"]
}

func (c *assignRouterEip) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *assignRouterEip) GetRegion() string {
    return c.req["Region"]
}

func (c *assignRouterEip) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *assignRouterEip) GetRouterId() string {
    return c.req["RouterId"]
}

func (c *assignRouterEip) SetRouterId(varRouterId string) {
    c.req["RouterId"] = varRouterId
}

func (c *assignRouterEip) GetEip() string {
    return c.req["Eip"]
}

func (c *assignRouterEip) SetEip(varEip string) {
    c.req["Eip"] = varEip
}


