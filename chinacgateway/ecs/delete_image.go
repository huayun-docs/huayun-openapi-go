package chinacgatewayecs

import (
)

type deleteImage struct {
    req        	map[string]string
}

func DeleteImage() *deleteImage {
    c := &deleteImage {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DeleteImage"
    return c
}

func (c *deleteImage) GetRequest() map[string]string {
    return c.req
}

func (c *deleteImage) GetAction() string {
    return c.req["Action"]
}

func (c *deleteImage) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *deleteImage) GetRegion() string {
    return c.req["Region"]
}

func (c *deleteImage) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *deleteImage) GetId() string {
    return c.req["Id"]
}

func (c *deleteImage) SetId(varId string) {
    c.req["Id"] = varId
}


