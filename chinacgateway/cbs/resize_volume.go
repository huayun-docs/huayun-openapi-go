package chinacgatewaycbs

import (
)

type resizeVolume struct {
    req        	map[string]string
}

func ResizeVolume() *resizeVolume {
    c := &resizeVolume {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ResizeVolume"
    return c
}

func (c *resizeVolume) GetRequest() map[string]string {
    return c.req
}

func (c *resizeVolume) GetAction() string {
    return c.req["Action"]
}

func (c *resizeVolume) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *resizeVolume) GetRegion() string {
    return c.req["Region"]
}

func (c *resizeVolume) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *resizeVolume) GetId() string {
    return c.req["Id"]
}

func (c *resizeVolume) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *resizeVolume) GetSize() string {
    return c.req["Size"]
}

func (c *resizeVolume) SetSize(varSize string) {
    c.req["Size"] = varSize
}


