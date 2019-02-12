package chinacgatewayvpc

import (
)

type releaseEips struct {
    req        	map[string]string
}

func ReleaseEips() *releaseEips {
    c := &releaseEips {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ReleaseEips"
    return c
}

func (c *releaseEips) GetRequest() map[string]string {
    return c.req
}

func (c *releaseEips) GetAction() string {
    return c.req["Action"]
}

func (c *releaseEips) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *releaseEips) GetRegion() string {
    return c.req["Region"]
}

func (c *releaseEips) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *releaseEips) GetEip0() string {
    return c.req["Eip.0"]
}

func (c *releaseEips) SetEip0(varEip0 string) {
    c.req["Eip.0"] = varEip0
}


