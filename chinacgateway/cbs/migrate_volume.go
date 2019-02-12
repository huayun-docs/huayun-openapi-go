package chinacgatewaycbs

import (
)

type migrateVolume struct {
    req        	map[string]string
}

func MigrateVolume() *migrateVolume {
    c := &migrateVolume {
        req:			make(map[string]string),
    }

    c.req["Action"] = "MigrateVolume"
    return c
}

func (c *migrateVolume) GetRequest() map[string]string {
    return c.req
}

func (c *migrateVolume) GetAction() string {
    return c.req["Action"]
}

func (c *migrateVolume) SetAction(varAction string) {
    c.req["Action"] = varAction
}

func (c *migrateVolume) GetRegion() string {
    return c.req["Region"]
}

func (c *migrateVolume) SetRegion(varRegion string) {
    c.req["Region"] = varRegion
}

func (c *migrateVolume) GetId() string {
    return c.req["Id"]
}

func (c *migrateVolume) SetId(varId string) {
    c.req["Id"] = varId
}

func (c *migrateVolume) GetType() string {
    return c.req["Type"]
}

func (c *migrateVolume) SetType(varType string) {
    c.req["Type"] = varType
}


