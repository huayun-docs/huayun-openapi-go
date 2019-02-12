package chinacgatewayvpc

import (
)

type modifyRouterAttributes struct {
    req        	map[string]string
}

func ModifyRouterAttributes() *modifyRouterAttributes {
    c := &modifyRouterAttributes {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ModifyRouterAttributes"
    return c
}

func (c *modifyRouterAttributes) GetRequest() map[string]string {
    return c.req
}

func (c *modifyRouterAttributes) GetAction() string {
    return c.req["Action"]
}

func (c *modifyRouterAttributes) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *modifyRouterAttributes) GetRegion() string {
    return c.req["Region"]
}

func (c *modifyRouterAttributes) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *modifyRouterAttributes) GetId() string {
    return c.req["Id"]
}

func (c *modifyRouterAttributes) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *modifyRouterAttributes) GetName() string {
    return c.req["Name"]
}

func (c *modifyRouterAttributes) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *modifyRouterAttributes) GetDescription() string {
    return c.req["Description"]
}

func (c *modifyRouterAttributes) SetDescription(varDescription string) {
    c.req["Description"] = varDescription
}


