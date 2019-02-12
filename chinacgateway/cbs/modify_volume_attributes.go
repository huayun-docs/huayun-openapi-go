package chinacgatewaycbs

import (
)

type modifyVolumeAttributes struct {
    req        	map[string]string
}

func ModifyVolumeAttributes() *modifyVolumeAttributes {
    c := &modifyVolumeAttributes {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ModifyVolumeAttributes"
    return c
}

func (c *modifyVolumeAttributes) GetRequest() map[string]string {
    return c.req
}

func (c *modifyVolumeAttributes) GetAction() string {
    return c.req["Action"]
}

func (c *modifyVolumeAttributes) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *modifyVolumeAttributes) GetRegion() string {
    return c.req["Region"]
}

func (c *modifyVolumeAttributes) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *modifyVolumeAttributes) GetId() string {
    return c.req["Id"]
}

func (c *modifyVolumeAttributes) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *modifyVolumeAttributes) GetName() string {
    return c.req["Name"]
}

func (c *modifyVolumeAttributes) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *modifyVolumeAttributes) GetDescription() string {
    return c.req["Description"]
}

func (c *modifyVolumeAttributes) SetDescription(varDescription string) {
    c.req["Description"] = varDescription
}


