package main

import (
	nr "github.com/yvasiyarov/newrelic_platform_go"
)

func InitTableStatsComponent(server_id string, verbose bool) *nr.PluginComponent {

	component := nr.NewPluginComponent(server_id, "com.github.maciejmrowiec.pg_monitor", verbose)

	return component
}
