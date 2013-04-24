package server

import (
	"github.com/toqueteos/minero/config"
)

func ConfigCreate() *config.Config {
	c := config.NewFrom(ConfigDefault())
	c.Save("./server.conf")
	return c
}

// ConfigDefault
func ConfigDefault() config.Map {
	return config.Map{
		"server.ip":                  "",
		"server.port":                "25565",
		"server.difficulty":          "1",
		"server.gamemode":            "0",
		"server.hardcore":            "false",
		"server.max_players":         "20",
		"server.motd":                "A minero server",
		"server.online_mode":         "true",
		"server.pvp":                 "true",
		"server.query.enable":        "false",
		"server.query.port":          "25565",
		"server.rcon.enable":         "false",
		"server.rcon.port":           "25575",
		"server.rcon.password":       "",
		"server.texture_pack":        "",
		"server.view_distance":       "10",
		"server.white_list":          "false",
		"spawn.animals":              "true",
		"spawn.monsters":             "true",
		"spawn.npcs":                 "true",
		"spawn.protection.enabled":   "true",
		"spawn.protection.shape":     "square",
		"spawn.protection.size":      "10",
		"worlds.level_name":          "minero_test",
		"worlds.level_seed":          "0",
		"worlds.level_type":          "DEFAULT",
		"worlds.flight":              "false",
		"worlds.generate_structures": "true",
		"worlds.generator_settings":  "",
		"worlds.max_build_height":    "256",
		"worlds.end":                 "false",
		"worlds.nether":              "false",
		"worlds.snooper_enabled":     "true",
	}
}