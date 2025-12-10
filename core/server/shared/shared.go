package shared

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/gin-contrib/cors"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var Builder = goqu.Dialect("postgres")
var Cors = cors.New(cors.Config{
	AllowMethods:     []string{"*"},
	AllowHeaders:     []string{"*"},
	ExposeHeaders:    []string{"*"},
	MaxAge:           1800,
	AllowCredentials: true,
	AllowAllOrigins:  true,
})

func SetDB(db *sqlx.DB) {
	DB = db
}
func GetDB() *sqlx.DB {
	return DB
}
