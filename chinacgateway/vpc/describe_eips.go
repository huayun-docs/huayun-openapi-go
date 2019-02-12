package chinacgatewayvpc

import (
    "strconv"
)

type describeEips struct {
    req        	map[string]string
}

func DescribeEips() *describeEips {
    c := &describeEips {
        req:			make(map[string]string),
    }

    c.req["Action"] = "DescribeEips"
    return c
}

func (c *describeEips) GetRequest() map[string]string {
    return c.req
}

func (c *describeEips) GetAction() string {
    return c.req["Action"]
}

func (c *describeEips) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *describeEips) GetRegion() string {
    return c.req["Region"]
}

func (c *describeEips) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *describeEips) GetOffset() int {
    nOffset, _ := strconv.Atoi(c.req["Offset"])
    return nOffset
}

func (c *describeEips) SetOffset(varOffset int) {
    c.req["Offset"] = strconv.Itoa(varOffset)
}

func (c *describeEips) GetCount() int {
    nCount, _ := strconv.Atoi(c.req["Count"])
    return nCount
}

func (c *describeEips) SetCount(varCount int) {
    c.req["Count"] = strconv.Itoa(varCount)
}

func (c *describeEips) GetEip0() string {
    return c.req["Eip.0"]
}

func (c *describeEips) SetEip0(varEip0 string) {
    c.req["Eip.0"] = varEip0
}

func (c *describeEips) GetName() string {
    return c.req["Name"]
}

func (c *describeEips) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *describeEips) GetStartTime() string {
    return c.req["StartTime"]
}

func (c *describeEips) SetStartTime(varStartTime string) {
    c.req["StartTime"] = varStartTime
}

func (c *describeEips) GetEndTime() string {
    return c.req["EndTime"]
}

func (c *describeEips) SetEndTime(varEndTime string) {
    c.req["EndTime"] = varEndTime
}

func (c *describeEips) GetProductStatus() string {
    return c.req["ProductStatus"]
}

func (c *describeEips) SetProductStatus(varProductStatus string) {
    c.req["ProductStatus"] = varProductStatus
}

func (c *describeEips) GetPayType() string {
    return c.req["PayType"]
}

func (c *describeEips) SetPayType(varPayType string) {
    c.req["PayType"] = varPayType
}


