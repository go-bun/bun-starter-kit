package drivers

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func PostgresqlDriver(dns string) *bun.DB {
	hsqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dns)))
	return bun.NewDB(hsqldb, pgdialect.New())
}
