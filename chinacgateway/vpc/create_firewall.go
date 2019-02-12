package chinacgatewayvpc

import (
)

type createFirewall struct {
    req        	map[string]string
}

func CreateFirewall() *createFirewall {
    c := &createFirewall {
        req:			make(map[string]string),
    }

    c.req["Action"] = "CreateFirewall"
    return c
}

func (c *createFirewall) GetRequest() map[string]string {
    return c.req
}

func (c *createFirewall) GetAction() string {
    return c.req["Action"]
}

func (c *createFirewall) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *createFirewall) GetRegion() string {
    return c.req["Region"]
}

func (c *createFirewall) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *createFirewall) GetName() string {
    return c.req["Name"]
}

func (c *createFirewall) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *createFirewall) GetDescription() string {
    return c.req["Description"]
}

func (c *createFirewall) SetDescription(varDescription string) {
    c.req["Description"] = varDescription
}

func (c *createFirewall) GetType() string {
    return c.req["Type"]
}

func (c *createFirewall) SetType(varType string) {
    c.req["Type"] = varType
}


