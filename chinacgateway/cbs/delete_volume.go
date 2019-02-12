package chinacgatewaycbs

import (
)

type deleteVolume struct {
    req        	map[string]string
}

func DeleteVolume() *deleteVolume {
    c := &deleteVolume {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DeleteVolume"
    return c
}

func (c *deleteVolume) GetRequest() map[string]string {
    return c.req
}

func (c *deleteVolume) GetAction() string {
    return c.req["Action"]
}

func (c *deleteVolume) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *deleteVolume) GetRegion() string {
    return c.req["Region"]
}

func (c *deleteVolume) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *deleteVolume) GetId() string {
    return c.req["Id"]
}

func (c *deleteVolume) SetId(varId string) {
    c.req["Id"] = varId
}


