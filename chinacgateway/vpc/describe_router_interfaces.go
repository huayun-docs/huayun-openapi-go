package chinacgatewayvpc

import (
)

type describeRouterInterfaces struct {
    req        	map[string]string
}

func DescribeRouterInterfaces() *describeRouterInterfaces {
    c := &describeRouterInterfaces {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DescribeRouterInterfaces"
    return c
}

func (c *describeRouterInterfaces) GetRequest() map[string]string {
    return c.req
}

func (c *describeRouterInterfaces) GetAction() string {
    return c.req["Action"]
}

func (c *describeRouterInterfaces) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *describeRouterInterfaces) GetRegion() string {
    return c.req["Region"]
}

func (c *describeRouterInterfaces) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *describeRouterInterfaces) GetOffset() string {
    return c.req["Offset"]
}

func (c *describeRouterInterfaces) SetOffset(varOffset string) {
    c.req["Offset"] = varOffset
}

func (c *describeRouterInterfaces) GetCount() string {
    return c.req["Count"]
}

func (c *describeRouterInterfaces) SetCount(varCount string) {
    c.req["Count"] = varCount
}

func (c *describeRouterInterfaces) GetId0() string {
    return c.req["Id.0"]
}

func (c *describeRouterInterfaces) SetId0(varId0 string) {
    c.req["Id.0"] = varId0
}

func (c *describeRouterInterfaces) GetRouterId0() string {
    return c.req["RouterId.0"]
}

func (c *describeRouterInterfaces) SetRouterId0(varRouterId0 string) {
    c.req["RouterId.0"] = varRouterId0
}

func (c *describeRouterInterfaces) GetIpAddress0() string {
    return c.req["IpAddress.0"]
}

func (c *describeRouterInterfaces) SetIpAddress0(varIpAddress0 string) {
    c.req["IpAddress.0"] = varIpAddress0
}

func (c *describeRouterInterfaces) GetNetworkId0() string {
    return c.req["NetworkId.0"]
}

func (c *describeRouterInterfaces) SetNetworkId0(varNetworkId0 string) {
    c.req["NetworkId.0"] = varNetworkId0
}


