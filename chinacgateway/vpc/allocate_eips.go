package chinacgatewayvpc

import (
    "strconv"
)

type allocateEips struct {
    req        	map[string]string
}

func AllocateEips() *allocateEips {
    c := &allocateEips {
        req:			make(map[string]string),
    }

    c.req["Action"] = "AllocateEips"
    return c
}

func (c *allocateEips) GetRequest() map[string]string {
    return c.req
}

func (c *allocateEips) GetAction() string {
    return c.req["Action"]
}

func (c *allocateEips) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *allocateEips) GetRegion() string {
    return c.req["Region"]
}

func (c *allocateEips) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *allocateEips) GetPayType() string {
    return c.req["PayType"]
}

func (c *allocateEips) SetPayType(varPayType string) {
    c.req["PayType"] = varPayType
}

func (c *allocateEips) GetPeriod() int {
    nPeriod, _ := strconv.Atoi(c.req["Period"])
    return nPeriod
}

func (c *allocateEips) SetPeriod(varPeriod int) {
    c.req["Period"] = strconv.Itoa(varPeriod)
}

func (c *allocateEips) GetBandwidth() int {
    nBandwidth, _ := strconv.Atoi(c.req["Bandwidth"])
    return nBandwidth
}

func (c *allocateEips) SetBandwidth(varBandwidth int) {
    c.req["Bandwidth"] = strconv.Itoa(varBandwidth)
}

func (c *allocateEips) GetName() string {
    return c.req["Name"]
}

func (c *allocateEips) SetName(varName string) {
    c.req["Name"] = varName
}

func (c *allocateEips) GetCount() int {
    nCount, _ := strconv.Atoi(c.req["Count"])
    return nCount
}

func (c *allocateEips) SetCount(varCount int) {
    c.req["Count"] = strconv.Itoa(varCount)
}


