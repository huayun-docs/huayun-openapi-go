package chinacgatewayvpc

import (
)

type createRouterInterface struct {
    req        	map[string]string
}

func CreateRouterInterface() *createRouterInterface {
    c := &createRouterInterface {
        req:			make(map[string]string),
    }

    c.req["Action"] = "CreateRouterInterface"
    return c
}

func (c *createRouterInterface) GetRequest() map[string]string {
    return c.req
}

func (c *createRouterInterface) GetAction() string {
    return c.req["Action"]
}

func (c *createRouterInterface) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *createRouterInterface) GetRegion() string {
    return c.req["Region"]
}

func (c *createRouterInterface) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *createRouterInterface) GetRouterId() string {
    return c.req["RouterId"]
}

func (c *createRouterInterface) SetRouterId(varRouterId string) {
    c.req["RouterId"] = varRouterId
}

func (c *createRouterInterface) GetNetworkId() string {
    return c.req["NetworkId"]
}

func (c *createRouterInterface) SetNetworkId(varNetworkId string) {
    c.req["NetworkId"] = varNetworkId
}


