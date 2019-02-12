package chinacgatewayecs

import (
)

type createInstanceSnapshot struct {
    req        	map[string]string
}

func CreateInstanceSnapshot() *createInstanceSnapshot {
    c := &createInstanceSnapshot {
        req:			make(map[string]string),
    }

    c.req["Action"] = "CreateInstanceSnapshot"
    return c
}

func (c *createInstanceSnapshot) GetRequest() map[string]string {
    return c.req
}

func (c *createInstanceSnapshot) GetAction() string {
    return c.req["Action"]
}

func (c *createInstanceSnapshot) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *createInstanceSnapshot) GetRegion() string {
    return c.req["Region"]
}

func (c *createInstanceSnapshot) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *createInstanceSnapshot) GetInstanceId() string {
    return c.req["InstanceId"]
}

func (c *createInstanceSnapshot) SetInstanceId(varInstanceId string) {
    c.req["InstanceId"] = varInstanceId
}

func (c *createInstanceSnapshot) GetCreateType() string {
    return c.req["CreateType"]
}

func (c *createInstanceSnapshot) SetCreateType(varCreateType string) {
    c.req["CreateType"] = varCreateType
}

func (c *createInstanceSnapshot) GetName() string {
    return c.req["Name"]
}

func (c *createInstanceSnapshot) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *createInstanceSnapshot) GetDescription() string {
    return c.req["Description"]
}

func (c *createInstanceSnapshot) SetDescription(varDescription string) {
    c.req["Description"] = varDescription
}


