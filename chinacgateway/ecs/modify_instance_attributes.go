package chinacgatewayecs

import (
)

type modifyInstanceAttributes struct {
    req        	map[string]string
}

func ModifyInstanceAttributes() *modifyInstanceAttributes {
    c := &modifyInstanceAttributes {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ModifyInstanceAttributes"
    return c
}

func (c *modifyInstanceAttributes) GetRequest() map[string]string {
    return c.req
}

func (c *modifyInstanceAttributes) GetRegion() string {
    return c.req["Region"]
}

func (c *modifyInstanceAttributes) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *modifyInstanceAttributes) GetAction() string {
    return c.req["Action"]
}

func (c *modifyInstanceAttributes) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *modifyInstanceAttributes) GetId() string {
    return c.req["Id"]
}

func (c *modifyInstanceAttributes) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *modifyInstanceAttributes) GetName() string {
    return c.req["Name"]
}

func (c *modifyInstanceAttributes) SetName(varName string) {
    c.req["Name"] = varName
}


