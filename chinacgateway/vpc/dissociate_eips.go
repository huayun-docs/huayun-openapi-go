package chinacgatewayvpc

import (
)

type dissociateEips struct {
    req        	map[string]string
}

func DissociateEips() *dissociateEips {
    c := &dissociateEips {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DissociateEips"
    return c
}

func (c *dissociateEips) GetRequest() map[string]string {
    return c.req
}

func (c *dissociateEips) GetAction() string {
    return c.req["Action"]
}

func (c *dissociateEips) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *dissociateEips) GetRegion() string {
    return c.req["Region"]
}

func (c *dissociateEips) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *dissociateEips) GetEip0() string {
    return c.req["Eip.0"]
}

func (c *dissociateEips) SetEip0(varEip0 string) {
    c.req["Eip.0"] = varEip0
}


