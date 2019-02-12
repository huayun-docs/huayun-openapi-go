package chinacgatewaycbs

import (
)

type describeSnapshots struct {
    req        	map[string]string
}

func DescribeSnapshots() *describeSnapshots {
    c := &describeSnapshots {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DescribeSnapshots"
    return c
}

func (c *describeSnapshots) GetRequest() map[string]string {
    return c.req
}

func (c *describeSnapshots) GetAction() string {
    return c.req["Action"]
}

func (c *describeSnapshots) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *describeSnapshots) GetRegion() string {
    return c.req["Region"]
}

func (c *describeSnapshots) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *describeSnapshots) GetVolumeId() string {
    return c.req["VolumeId"]
}

func (c *describeSnapshots) SetVolumeId(varVolumeId string) {
    c.req["VolumeId"] = varVolumeId
}

func (c *describeSnapshots) GetId0() string {
    return c.req["Id.0"]
}

func (c *describeSnapshots) SetId0(varId0 string) {
    c.req["Id.0"] = varId0
}

func (c *describeSnapshots) GetName() string {
    return c.req["Name"]
}

func (c *describeSnapshots) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *describeSnapshots) GetStartTime() string {
    return c.req["StartTime"]
}

func (c *describeSnapshots) SetStartTime(varStartTime string) {
    c.req["StartTime"] = varStartTime
}

func (c *describeSnapshots) GetEndTime() string {
    return c.req["EndTime"]
}

func (c *describeSnapshots) SetEndTime(varEndTime string) {
    c.req["EndTime"] = varEndTime
}

func (c *describeSnapshots) GetOffset() string {
    return c.req["Offset"]
}

func (c *describeSnapshots) SetOffset(varOffset string) {
    c.req["Offset"] = varOffset
}

func (c *describeSnapshots) GetCount() string {
    return c.req["Count"]
}

func (c *describeSnapshots) SetCount(varCount string) {
    c.req["Count"] = varCount
}

func (c *describeSnapshots) GetIsSystem() string {
    return c.req["IsSystem"]
}

func (c *describeSnapshots) SetIsSystem(varIsSystem string) {
    c.req["IsSystem"] = varIsSystem
}


