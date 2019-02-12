package chinacgatewayecs

import (
)

type modifyKeyPairAttributes struct {
    req        	map[string]string
}

func ModifyKeyPairAttributes() *modifyKeyPairAttributes {
    c := &modifyKeyPairAttributes {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ModifyKeyPairAttributes"
    return c
}

func (c *modifyKeyPairAttributes) GetRequest() map[string]string {
    return c.req
}

func (c *modifyKeyPairAttributes) GetAction() string {
    return c.req["Action"]
}

func (c *modifyKeyPairAttributes) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *modifyKeyPairAttributes) GetRegion() string {
    return c.req["Region"]
}

func (c *modifyKeyPairAttributes) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *modifyKeyPairAttributes) GetId() string {
    return c.req["Id"]
}

func (c *modifyKeyPairAttributes) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *modifyKeyPairAttributes) GetName() string {
    return c.req["Name"]
}

func (c *modifyKeyPairAttributes) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *modifyKeyPairAttributes) GetDescription() string {
    return c.req["Description"]
}

func (c *modifyKeyPairAttributes) SetDescription(varDescription string) {
    c.req["Description"] = varDescription
}


