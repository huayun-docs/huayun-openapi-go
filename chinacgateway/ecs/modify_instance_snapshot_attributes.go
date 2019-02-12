package chinacgatewayecs

import (
)

type modifyInstanceSnapshotAttributes struct {
    req        	map[string]string
}

func ModifyInstanceSnapshotAttributes() *modifyInstanceSnapshotAttributes {
    c := &modifyInstanceSnapshotAttributes {
        req:			make(map[string]string),
    }

    c.req["Action"] = "ModifyInstanceSnapshotAttributes"
    return c
}

func (c *modifyInstanceSnapshotAttributes) GetRequest() map[string]string {
    return c.req
}

func (c *modifyInstanceSnapshotAttributes) GetAction() string {
    return c.req["Action"]
}

func (c *modifyInstanceSnapshotAttributes) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *modifyInstanceSnapshotAttributes) GetRegion() string {
    return c.req["Region"]
}

func (c *modifyInstanceSnapshotAttributes) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *modifyInstanceSnapshotAttributes) GetId() string {
    return c.req["Id"]
}

func (c *modifyInstanceSnapshotAttributes) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *modifyInstanceSnapshotAttributes) GetName() string {
    return c.req["Name"]
}

func (c *modifyInstanceSnapshotAttributes) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *modifyInstanceSnapshotAttributes) GetDescription() string {
    return c.req["Description"]
}

func (c *modifyInstanceSnapshotAttributes) SetDescription(varDescription string) {
    c.req["Description"] = varDescription
}


