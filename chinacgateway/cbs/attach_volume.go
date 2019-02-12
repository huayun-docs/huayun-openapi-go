package chinacgatewaycbs

import (
)

type attachVolume struct {
    req        	map[string]string
}

func AttachVolume() *attachVolume {
    c := &attachVolume {
        req:			make(map[string]string),
    }

    c.req["Action"] = "AttachVolume"
    return c
}

func (c *attachVolume) GetRequest() map[string]string {
    return c.req
}

func (c *attachVolume) GetAction() string {
    return c.req["Action"]
}

func (c *attachVolume) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *attachVolume) GetRegion() string {
    return c.req["Region"]
}

func (c *attachVolume) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *attachVolume) GetVolumeId() string {
    return c.req["VolumeId"]
}

func (c *attachVolume) SetVolumeId(varVolumeId string) {
    c.req["VolumeId"] = varVolumeId
}

func (c *attachVolume) GetInstanceId() string {
    return c.req["InstanceId"]
}

func (c *attachVolume) SetInstanceId(varInstanceId string) {
    c.req["InstanceId"] = varInstanceId
}



