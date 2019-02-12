package chinacgatewaycbs

import (
    "strconv"
)

type describeVolumes struct {
    req        	map[string]string
}

func DescribeVolumes() *describeVolumes {
    c := &describeVolumes {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DescribeVolumes"
    return c
}

func (c *describeVolumes) GetRequest() map[string]string {
    return c.req
}

func (c *describeVolumes) GetAction() string {
    return c.req["Action"]
}

func (c *describeVolumes) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *describeVolumes) GetRegion() string {
    return c.req["Region"]
}

func (c *describeVolumes) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *describeVolumes) GetOffset() int {
    nOffset, _ := strconv.Atoi(c.req["Offset"])
    return nOffset
}

func (c *describeVolumes) SetOffset(varOffset int) {
    c.req["Offset"] = strconv.Itoa(varOffset)
}

func (c *describeVolumes) GetCount() int {
    nCount, _ := strconv.Atoi(c.req["Count"])
    return nCount
}

func (c *describeVolumes) SetCount(varCount int) {
    c.req["Count"] = strconv.Itoa(varCount)
}

func (c *describeVolumes) GetId0() string {
    return c.req["Id.0"]
}

func (c *describeVolumes) SetId0(varId0 string) {
    c.req["Id.0"] = varId0
}

func (c *describeVolumes) GetName() string {
    return c.req["Name"]
}

func (c *describeVolumes) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *describeVolumes) GetPayType() string {
    return c.req["PayType"]
}

func (c *describeVolumes) SetPayType(varPayType string) {
    c.req["PayType"] = varPayType
}

func (c *describeVolumes) GetProductType() string {
    return c.req["ProductType"]
}

func (c *describeVolumes) SetProductType(varProductType string) {
    c.req["ProductType"] = varProductType
}

func (c *describeVolumes) GetStatus() string {
    return c.req["Status"]
}

func (c *describeVolumes) SetStatus(varStatus string) {
    c.req["Status"] = varStatus
}

func (c *describeVolumes) GetUserId0() string {
    return c.req["UserId.0"]
}

func (c *describeVolumes) SetUserId0(varUserId0 string) {
    c.req["UserId.0"] = varUserId0
}

func (c *describeVolumes) GetProductStatus() string {
    return c.req["ProductStatus"]
}

func (c *describeVolumes) SetProductStatus(varProductStatus string) {
    c.req["ProductStatus"] = varProductStatus
}

func (c *describeVolumes) GetInstanceName() string {
    return c.req["InstanceName"]
}

func (c *describeVolumes) SetInstanceName(varInstanceName string) {
    c.req["InstanceName"] = varInstanceName
}

func (c *describeVolumes) GetDueStartTime() string {
    return c.req["DueStartTime"]
}

func (c *describeVolumes) SetDueStartTime(varDueStartTime string) {
    c.req["DueStartTime"] = varDueStartTime
}

func (c *describeVolumes) GetDueEndTime() string {
    return c.req["DueEndTime"]
}

func (c *describeVolumes) SetDueEndTime(varDueEndTime string) {
    c.req["DueEndTime"] = varDueEndTime
}


