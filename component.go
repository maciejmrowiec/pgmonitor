package main

import (
	"database/sql"
	ext "github.com/maciejmrowiec/pgmonitor/newrelic"
	platform "github.com/yvasiyarov/newrelic_platform_go"
)

func InitTupleComponent(db *sql.DB, hostname string, verbose bool) platform.IComponent {

	component := ext.NewDynamicPluginComponent(hostname, "com.github.maciejmrowiec.pgmonitor", verbose)

	component.AddDynamicMetrica(NewTableMetric(db, "tuple/table/$/percent/active", TupleActivePercent, "%"))
	component.AddDynamicMetrica(NewTableMetric(db, "tuple/table/$/percent/dead", TupleDeadPercent, "%"))
	component.AddDynamicMetrica(NewTableMetric(db, "tuple/table/$/percent/free", TupleFreePercent, "%"))

	return component
}

func InitTupleSummaryComponent(db *sql.DB, hostname string, verbose bool) platform.IComponent {

	component := platform.NewPluginComponent(hostname, "com.github.maciejmrowiec.pgmonitor", verbose)

	component.AddMetrica(NewTableAverageSummaryMetric(db, "tuple/summary/percent/active", TupleActivePercent, "%"))
	component.AddMetrica(NewTableAverageSummaryMetric(db, "tuple/summary/percent/dead", TupleDeadPercent, "%"))
	component.AddMetrica(NewTableAverageSummaryMetric(db, "tuple/summary/percent/free", TupleFreePercent, "%"))

	return component

}

func InitTableSizeComponent(db *sql.DB, hostname string, verbose bool) platform.IComponent {

	component := ext.NewDynamicPluginComponent(hostname, "com.github.maciejmrowiec.pgmonitor", verbose)

	component.AddDynamicMetrica(NewTableMetric(db, "disksize/table/$/table", TableDiskSize, "B"))
	component.AddDynamicMetrica(NewTableMetric(db, "disksize/table/$/index", TableIndexesDiskSize, "B"))

	component.AddDynamicMetrica(NewTableMetric(db, "disksize/summary/table/$", TableDiskSize, "B"))
	component.AddDynamicMetrica(NewTableMetric(db, "disksize/summary/index/$", TableIndexesDiskSize, "B"))

	return component
}
