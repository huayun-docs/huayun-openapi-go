package chinacgatewayvpc

import (
)

type modifyNetworkAttributes struct {
    req        	map[string]string
}

func ModifyNetworkAttributes() *modifyNetworkAttributes {
    c := &modifyNetworkAttributes {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ModifyNetworkAttributes"
    return c
}

func (c *modifyNetworkAttributes) GetRequest() map[string]string {
    return c.req
}

func (c *modifyNetworkAttributes) GetAction() string {
    return c.req["Action"]
}

func (c *modifyNetworkAttributes) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *modifyNetworkAttributes) GetRegion() string {
    return c.req["Region"]
}

func (c *modifyNetworkAttributes) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *modifyNetworkAttributes) GetId() string {
    return c.req["Id"]
}

func (c *modifyNetworkAttributes) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *modifyNetworkAttributes) GetName() string {
    return c.req["Name"]
}

func (c *modifyNetworkAttributes) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *modifyNetworkAttributes) GetDescription() string {
    return c.req["Description"]
}

func (c *modifyNetworkAttributes) SetDescription(varDescription string) {
    c.req["Description"] = varDescription
}


