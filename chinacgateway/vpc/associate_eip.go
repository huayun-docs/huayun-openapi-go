package chinacgatewayvpc

import (
)

type associateEip struct {
    req        	map[string]string
}

func AssociateEip() *associateEip {
    c := &associateEip {
        req:			make(map[string]string),
    }

    c.req["Action"] = "AssociateEip"
    return c
}

func (c *associateEip) GetRequest() map[string]string {
    return c.req
}

func (c *associateEip) GetAction() string {
    return c.req["Action"]
}

func (c *associateEip) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *associateEip) GetRegion() string {
    return c.req["Region"]
}

func (c *associateEip) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *associateEip) GetEip() string {
    return c.req["Eip"]
}

func (c *associateEip) SetEip(varEip string) {
    c.req["Eip"] = varEip
}

func (c *associateEip) GetLocalIpAddress() string {
    return c.req["LocalIpAddress"]
}

func (c *associateEip) SetLocalIpAddress(varLocalIpAddress string) {
    c.req["LocalIpAddress"] = varLocalIpAddress
}

func (c *associateEip) GetRouterId() string {
    return c.req["RouterId"]
}

func (c *associateEip) SetRouterId(varRouterId string) {
    c.req["RouterId"] = varRouterId
}

func (c *associateEip) GetLocalNetworkId() string {
    return c.req["LocalNetworkId"]
}

func (c *associateEip) SetLocalNetworkId(varLocalNetworkId string) {
    c.req["LocalNetworkId"] = varLocalNetworkId
}


