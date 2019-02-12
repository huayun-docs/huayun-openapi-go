package chinacgatewaycbs

import (
)

type setVolumeBootable struct {
    req        	map[string]string
}

func SetVolumeBootable() *setVolumeBootable {
    c := &setVolumeBootable {
        req:			make(map[string]string),
    }

    c.req["Action"] = "SetVolumeBootable"
    return c
}

func (c *setVolumeBootable) GetRequest() map[string]string {
    return c.req
}

func (c *setVolumeBootable) GetAction() string {
    return c.req["Action"]
}

func (c *setVolumeBootable) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *setVolumeBootable) GetRegion() string {
    return c.req["Region"]
}

func (c *setVolumeBootable) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *setVolumeBootable) GetVolumeId() string {
    return c.req["VolumeId"]
}

func (c *setVolumeBootable) SetVolumeId(varVolumeId string) {
    c.req["VolumeId"] = varVolumeId
}

func (c *setVolumeBootable) GetImageId() string {
    return c.req["ImageId"]
}

func (c *setVolumeBootable) SetImageId(varImageId string) {
    c.req["ImageId"] = varImageId
}


