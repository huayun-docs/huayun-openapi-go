package chinacgatewayvpc

import (
)

type createNetwork struct {
    req        	map[string]string
}

func CreateNetwork() *createNetwork {
    c := &createNetwork {
        req:			make(map[string]string),
    }

    c.req["Action"] = "CreateNetwork"
    return c
}

func (c *createNetwork) GetRequest() map[string]string {
    return c.req
}


func (c *createNetwork) GetRegion() string {
    return c.req["Region"]
}

func (c *createNetwork) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *createNetwork) GetCidr() string {
    return c.req["Cidr"]
}

func (c *createNetwork) SetCidr(varCidr string) {
    c.req["Cidr"] = varCidr
}

func (c *createNetwork) GetDhcp() string {
    return c.req["Dhcp"]
}

func (c *createNetwork) SetDhcp(varDhcp string) {
    c.req["Dhcp"] = varDhcp
}

func (c *createNetwork) GetGateway() string {
    return c.req["Gateway"]
}

func (c *createNetwork) SetGateway(varGateway string) {
    c.req["Gateway"] = varGateway
}

func (c *createNetwork) GetName() string {
    return c.req["Name"]
}

func (c *createNetwork) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *createNetwork) GetDescription() string {
    return c.req["Description"]
}

func (c *createNetwork) SetDescription(varDescription string) {
    c.req["Description"] = varDescription
}

func (c *createNetwork) GetIpStart() string {
    return c.req["IpStart"]
}

func (c *createNetwork) SetIpStart(varIpStart string) {
    c.req["IpStart"] = varIpStart
}

func (c *createNetwork) GetIpEnd() string {
    return c.req["IpEnd"]
}

func (c *createNetwork) SetIpEnd(varIpEnd string) {
    c.req["IpEnd"] = varIpEnd
}


