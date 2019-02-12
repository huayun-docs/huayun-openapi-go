package chinacgatewayecs

import (
)

type rebuildInstance struct {
    req        	map[string]string
}

func RebuildInstance() *rebuildInstance {
    c := &rebuildInstance {
        req:			make(map[string]string),
    }

    c.req["Action"] = "RebuildInstance"
    return c
}

func (c *rebuildInstance) GetRequest() map[string]string {
    return c.req
}

func (c *rebuildInstance) GetRegion() string {
    return c.req["Region"]
}

func (c *rebuildInstance) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *rebuildInstance) GetAction() string {
    return c.req["Action"]
}

func (c *rebuildInstance) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *rebuildInstance) GetId() string {
    return c.req["Id"]
}

func (c *rebuildInstance) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *rebuildInstance) GetImageId() string {
    return c.req["ImageId"]
}

func (c *rebuildInstance) SetImageId(varImageId string) {
    c.req["ImageId"] = varImageId
}

func (c *rebuildInstance) GetPassword() string {
    return c.req["Password"]
}

func (c *rebuildInstance) SetPassword(varPassword string) {
    c.req["Password"] = varPassword
}


