package chinacgatewayecs

import (
)

type terminateInstance struct {
    req        	map[string]string
}

func TerminateInstance() *terminateInstance {
    c := &terminateInstance {
        req:			make(map[string]string),
    }

    c.req["Action"] = "TerminateInstance"
    return c
}

func (c *terminateInstance) GetRequest() map[string]string {
    return c.req
}

func (c *terminateInstance) GetRegion() string {
    return c.req["Region"]
}

func (c *terminateInstance) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *terminateInstance) GetAction() string {
    return c.req["Action"]
}

func (c *terminateInstance) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *terminateInstance) GetId() string {
    return c.req["Id"]
}

func (c *terminateInstance) SetId(varId string) {
    c.req["Id"] = varId
}


