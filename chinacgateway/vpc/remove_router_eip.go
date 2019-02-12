package chinacgatewayvpc

import (
)

type removeRouterEip struct {
    req        	map[string]string
}

func RemoveRouterEip() *removeRouterEip {
    c := &removeRouterEip {
        req:			make(map[string]string),
    }

    c.req["Action"] = "RemoveRouterEip"
    return c
}

func (c *removeRouterEip) GetRequest() map[string]string {
    return c.req
}

func (c *removeRouterEip) GetAction() string {
    return c.req["Action"]
}

func (c *removeRouterEip) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *removeRouterEip) GetRegion() string {
    return c.req["Region"]
}

func (c *removeRouterEip) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *removeRouterEip) GetRouterId() string {
    return c.req["RouterId"]
}

func (c *removeRouterEip) SetRouterId(varRouterId string) {
    c.req["RouterId"] = varRouterId
}


