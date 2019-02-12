package chinacgatewaycbs

import (
)

type detachVolume struct {
    req        	map[string]string
}

func DetachVolume() *detachVolume {
    c := &detachVolume {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DetachVolume"
    return c
}

func (c *detachVolume) GetRequest() map[string]string {
    return c.req
}

func (c *detachVolume) GetAction() string {
    return c.req["Action"]
}

func (c *detachVolume) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *detachVolume) GetRegion() string {
    return c.req["Region"]
}

func (c *detachVolume) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *detachVolume) GetVolumeId() string {
    return c.req["VolumeId"]
}

func (c *detachVolume) SetVolumeId(varVolumeId string) {
    c.req["VolumeId"] = varVolumeId
}

func (c *detachVolume) GetInstanceId() string {
    return c.req["InstanceId"]
}

func (c *detachVolume) SetInstanceId(varInstanceId string) {
    c.req["InstanceId"] = varInstanceId
}


