package xpostgres

import (
	"database/sql"
	"errors"
	"fmt"
	logger "github.com/ipfs/go-log"
	"go-pjt-base/pkg/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

var (
	log                      = logger.Logger("xpostgres")
	ERR_DB_INSTANCE_IS_EMPTY = errors.New("database instance is empty")
)

var (
	cli *PostgresClient
)

type PostgresClient struct {
	db  *gorm.DB
	cfg conf.Postgres
}

func NewPostgresClient(cfg conf.Postgres) *PostgresClient {
	cli = &PostgresClient{cfg: cfg}
	cli.db, _ = ConnectDB(cfg)
	return cli
}

func GetDB() *gorm.DB {
	if cli.db == nil {
		cli.db, _ = ConnectDB(cli.cfg)
	}
	return cli.db
}

func GetTX() *gorm.DB {
	return GetDB().Begin()
}

func Transaction(handle func(tx *gorm.DB) (err error)) (err error) {
	var (
		db *gorm.DB
	)
	db = GetDB()
	if db == nil {
		err = ERR_DB_INSTANCE_IS_EMPTY
		return
	}
	tx := db.Begin(&sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	err = handle(tx)
	if err != nil {
		err = tx.Rollback().Error
		return
	}
	err = tx.Commit().Error
	return
}

func ConnectDB(cfg conf.Postgres) (db *gorm.DB, err error) {
	var (
		args string
		opts *gorm.Config
	)
	port, err := strconv.Atoi(cfg.Port)
	if err != nil {
		panic(err.Error())
		return nil, err
	}
	args = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		cfg.Address,
		cfg.Username,
		cfg.Password,
		cfg.Db,
		port)

	opts = &gorm.Config{
		SkipDefaultTransaction: false,
		PrepareStmt:            false,
	}

	db, err = gorm.Open(postgres.Open(args), opts)
	if err != nil {
		panic(err.Error())
		return
	}
	db = db.Debug()

	return
}
