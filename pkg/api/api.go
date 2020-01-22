// Copyright 2017 Emir Ribic. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// GORSK - Go(lang) restful starter kit
//
// API Docs for GORSK v1
//
// 	 Terms Of Service:  N/A
//     Schemes: http
//     Version: 2.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Emir Ribic <ribice@gmail.com> https://ribice.ba
//     Host: localhost:8080
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer: []
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package api

import (
	"crypto/sha1"

	"go.mongodb.org/mongo-driver/mongo"
	"oceanbolt.com/iamservice/pkg/utl/zlog"

	"oceanbolt.com/iamservice/pkg/api/port"
	pl "oceanbolt.com/iamservice/pkg/api/port/logging"
	pt "oceanbolt.com/iamservice/pkg/api/port/transport"

	"oceanbolt.com/iamservice/internal/echoapi/config"
	"oceanbolt.com/iamservice/pkg/utl/middleware/jwt"
	"oceanbolt.com/iamservice/pkg/utl/rbac"
	"oceanbolt.com/iamservice/pkg/utl/secure"
	"oceanbolt.com/iamservice/pkg/utl/server"
)

// Start starts the API service
func Start(db *mongo.Database, cfg *config.Configuration) error {
	sec := secure.New(cfg.App.MinPasswordStr, sha1.New())
	rbac := rbac.New()
	jwt := jwt.New(cfg.JWT.Secret, cfg.JWT.SigningAlgorithm, cfg.JWT.Duration)
	log := zlog.New()

	e := server.New()
	e.Static("/swaggerui", cfg.App.SwaggerUIPath)

	v1 := e.Group("/v1")
	v1.Use(jwt.MWFunc())

	pt.NewHTTP(pl.New(port.Initialize(db, rbac, sec), log), v1)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
