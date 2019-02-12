package chinacgatewayecs

import (
)

type modifyImageAttributes struct {
    req        	map[string]string
}

func ModifyImageAttributes() *modifyImageAttributes {
    c := &modifyImageAttributes {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ModifyImageAttributes"
    return c
}

func (c *modifyImageAttributes) GetRequest() map[string]string {
    return c.req
}

func (c *modifyImageAttributes) GetAction() string {
    return c.req["Action"]
}

func (c *modifyImageAttributes) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *modifyImageAttributes) GetRegion() string {
    return c.req["Region"]
}

func (c *modifyImageAttributes) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *modifyImageAttributes) GetId() string {
    return c.req["Id"]
}

func (c *modifyImageAttributes) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *modifyImageAttributes) GetName() string {
    return c.req["Name"]
}

func (c *modifyImageAttributes) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *modifyImageAttributes) GetDescription() string {
    return c.req["Description"]
}

func (c *modifyImageAttributes) SetDescription(varDescription string) {
    c.req["Description"] = varDescription
}


