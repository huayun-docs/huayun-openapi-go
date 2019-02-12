package chinacgatewaycbs

import (
)

type modifySnapshotAttributes struct {
    req        	map[string]string
}

func ModifySnapshotAttributes() *modifySnapshotAttributes {
    c := &modifySnapshotAttributes {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ModifySnapshotAttributes"
    return c
}

func (c *modifySnapshotAttributes) GetRequest() map[string]string {
    return c.req
}

func (c *modifySnapshotAttributes) GetAction() string {
    return c.req["Action"]
}

func (c *modifySnapshotAttributes) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *modifySnapshotAttributes) GetRegion() string {
    return c.req["Region"]
}

func (c *modifySnapshotAttributes) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *modifySnapshotAttributes) GetId() string {
    return c.req["Id"]
}

func (c *modifySnapshotAttributes) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *modifySnapshotAttributes) GetName() string {
    return c.req["Name"]
}

func (c *modifySnapshotAttributes) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *modifySnapshotAttributes) GetDescription() string {
    return c.req["Description"]
}

func (c *modifySnapshotAttributes) SetDescription(varDescription string) {
    c.req["Description"] = varDescription
}


