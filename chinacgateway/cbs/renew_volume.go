package chinacgatewaycbs

import (
    "strconv"
)

type renewVolume struct {
    req        	map[string]string
}

func RenewVolume() *renewVolume {
    c := &renewVolume {
        req:			make(map[string]string),
    }

    c.req["Action"] = "RenewVolume"
    return c
}

func (c *renewVolume) GetRequest() map[string]string {
    return c.req
}

func (c *renewVolume) GetAction() string {
    return c.req["Action"]
}

func (c *renewVolume) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *renewVolume) GetRegion() string {
    return c.req["Region"]
}

func (c *renewVolume) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *renewVolume) GetId() string {
    return c.req["Id"]
}

func (c *renewVolume) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *renewVolume) GetPayType() string {
    return c.req["PayType"]
}

func (c *renewVolume) SetPayType(varPayType string) {
    c.req["PayType"] = varPayType
}

func (c *renewVolume) GetPeriod() int {
    nPeriod, _ := strconv.Atoi(c.req["Period"])
    return nPeriod
}

func (c *renewVolume) SetPeriod(varPeriod int) {
    c.req["Period"] = strconv.Itoa(varPeriod)
}


