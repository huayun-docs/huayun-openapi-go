package chinacgatewayecs

import (
)

type captureInstance struct {
    req        	map[string]string
}

func CaptureInstance() *captureInstance {
    c := &captureInstance {
        req:			make(map[string]string),
    }

    c.req["Action"] = "CaptureInstance"
    return c
}

func (c *captureInstance) GetRequest() map[string]string {
    return c.req
}

func (c *captureInstance) GetAction() string {
    return c.req["Action"]
}

func (c *captureInstance) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *captureInstance) GetRegion() string {
    return c.req["Region"]
}

func (c *captureInstance) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *captureInstance) GetInstanceId() string {
    return c.req["InstanceId"]
}

func (c *captureInstance) SetInstanceId(varInstanceId string) {
    c.req["InstanceId"] = varInstanceId
}

func (c *captureInstance) GetVolumeType0() string {
    return c.req["VolumeType.0"]
}

func (c *captureInstance) SetVolumeType0(varVolumeType0 string) {
    c.req["VolumeType.0"] = varVolumeType0
}

func (c *captureInstance) GetName() string {
    return c.req["Name"]
}

func (c *captureInstance) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *captureInstance) GetDescription() string {
    return c.req["Description"]
}

func (c *captureInstance) SetDescription(varDescription string) {
    c.req["Description"] = varDescription
}


