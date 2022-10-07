// Copyright 2022 Cisco Systems, Inc. and its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package db

import (
	"fmt"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"time"
)

const (
	dbTypeMySQL    = "mysql"
	dbTypePostgres = "postgres"
)

var (
	dbType     = "mysql" // "postgres"
	dbHost     = "localhost"
	dbPort     = "3306" // "5432"
	dbUser     = "root"
	dbPassword = "123456"
	dbDatabase = "apiregistry"
)

// ClientFlags returns cli flags for creating new db client
func ClientFlags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "db-type",
			Usage:       "db type",
			Value:       dbType,
			Destination: &dbType,
			EnvVars:     []string{"DB_TYPE"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "db-host",
			Usage:       "db host",
			Value:       dbHost,
			Destination: &dbHost,
			EnvVars:     []string{"DB_HOST"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "db-port",
			Usage:       "db port",
			Value:       dbPort,
			Destination: &dbPort,
			EnvVars:     []string{"DB_PORT"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "db-user",
			Usage:       "db user",
			Value:       dbUser,
			Destination: &dbUser,
			EnvVars:     []string{"DB_USER"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "db-password",
			Usage:       "db password",
			Value:       dbPassword,
			Destination: &dbPassword,
			EnvVars:     []string{"DB_PASSWORD"},
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "db-database",
			Usage:       "db database",
			Value:       dbDatabase,
			Destination: &dbDatabase,
			EnvVars:     []string{"DB_DATABASE"},
		}),
	}
}

// Client is a client for db service
type Client struct {
	*gorm.DB
}

var dbClient *Client
var dbMutex sync.Mutex

// NewDBClient create a new Client
func NewDBClient(cfg *shared.AppConfig) (*Client, error) {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	if dbClient != nil {
		return dbClient, nil
	}

	db, err := initDB()
	if err != nil {
		return nil, err
	}

	dbClient = &Client{DB: db}

	return dbClient, nil
}

func initDB() (*gorm.DB, error) {
	var d gorm.Dialector
	var safeDSN string

	switch dbType {
	case dbTypeMySQL:
		d, safeDSN = initMySQL()
	case dbTypePostgres:
		d, safeDSN = initPostgres()
	}

	db, err := gorm.Open(d, &gorm.Config{})
	if err != nil {
		shared.LogErrorf("failed to connect to %s DB at %s: %v", dbType, safeDSN, err.Error())
		return nil, err
	}
	shared.LogInfof("connected to %s DB at: %s", dbType, safeDSN)

	db = db.Set("gorm:auto_preload", true)
	db.Logger.LogMode(logger.Warn)

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func initMySQL() (gorm.Dialector, string) {
	dsnFormat := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=preferred"
	dsn := fmt.Sprintf(dsnFormat, dbUser, dbPassword, dbHost, dbPort, dbDatabase)
	safeDSN := fmt.Sprintf(dsnFormat, dbUser, "*", dbHost, dbPort, dbDatabase)

	return mysql.Open(dsn), safeDSN
}

func initPostgres() (gorm.Dialector, string) {
	dsnFormat := "user=%s password=%s host=%s port=%s dbname=%s sslmode=disable"
	dsn := fmt.Sprintf(dsnFormat, dbUser, dbPassword, dbHost, dbPort, dbDatabase)
	safeDSN := fmt.Sprintf(dsnFormat, dbUser, "*", dbHost, dbPort, dbDatabase)

	return postgres.Open(dsn), safeDSN
}
