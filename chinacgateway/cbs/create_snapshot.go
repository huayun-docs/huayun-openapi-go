package chinacgatewaycbs

import (
)

type createSnapshot struct {
    req        	map[string]string
}

func CreateSnapshot() *createSnapshot {
    c := &createSnapshot {
        req:			make(map[string]string),
    }

    c.req["Action"] = "CreateSnapshot"
    return c
}

func (c *createSnapshot) GetRequest() map[string]string {
    return c.req
}

func (c *createSnapshot) GetAction() string {
    return c.req["Action"]
}

func (c *createSnapshot) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *createSnapshot) GetRegion() string {
    return c.req["Region"]
}

func (c *createSnapshot) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *createSnapshot) GetVolumeId() string {
    return c.req["VolumeId"]
}

func (c *createSnapshot) SetVolumeId(varVolumeId string) {
    c.req["VolumeId"] = varVolumeId
}

func (c *createSnapshot) GetCreateType() string {
    return c.req["CreateType"]
}

func (c *createSnapshot) SetCreateType(varCreateType string) {
    c.req["CreateType"] = varCreateType
}

func (c *createSnapshot) GetName() string {
    return c.req["Name"]
}

func (c *createSnapshot) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *createSnapshot) GetDescription() string {
    return c.req["Description"]
}

func (c *createSnapshot) SetDescription(varDescription string) {
    c.req["Description"] = varDescription
}


