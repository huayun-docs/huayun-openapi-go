package chinacgatewayecs

import (
)

type changeInstancePassword struct {
    req        	map[string]string
}

func ChangeInstancePassword() *changeInstancePassword {
    c := &changeInstancePassword {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ChangeInstancePassword"
    return c
}

func (c *changeInstancePassword) GetRequest() map[string]string {
    return c.req
}

func (c *changeInstancePassword) GetRegion() string {
    return c.req["Region"]
}

func (c *changeInstancePassword) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *changeInstancePassword) GetAction() string {
    return c.req["Action"]
}

func (c *changeInstancePassword) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *changeInstancePassword) GetId() string {
    return c.req["Id"]
}

func (c *changeInstancePassword) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *changeInstancePassword) GetPassword() string {
    return c.req["Password"]
}

func (c *changeInstancePassword) SetPassword(varPassword string) {
    c.req["Password"] = varPassword
}


