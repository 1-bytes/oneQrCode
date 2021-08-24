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
	"golang.org/x/sync/errgroup"
	"log"
	"oneQrCode/bootstrap"
	configs "oneQrCode/config"
)

var g errgroup.Group

func init() {
	configs.Initialize()
}

// main ...
func main() {
	bootstrap.SetupDB()
	s := bootstrap.SetupServe(bootstrap.SetupRoute())

	// run server
	g.Go(func() error {
		return s.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
