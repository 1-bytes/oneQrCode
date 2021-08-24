// Copyright (C) 2021 Enzo
//
// This program is free software:
// you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the
// Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
// without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License along with this program. If not,
// see <https://www.gnu.org/licenses/>

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"oneQrCode/bootstrap"
	configs "oneQrCode/config"
	"oneQrCode/pkg/config"
)

func init() {
	configs.Initialize()
}

// main .
func main() {
	bootstrap.SetupDB()

	r := gin.New()
	bootstrap.SetupRoute(r)

	err := r.Run(
		fmt.Sprintf("%s:%s",
			config.GetString("app.url"),
			config.GetString("app.port")))

	if err != nil {
		log.Fatal("Run the Service failed")
	}
}
