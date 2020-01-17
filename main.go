/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/fogetu/go_system/system_config"
	"github.com/fogetu/miner_service_impl/classes/impl"
	"github.com/fogetu/miner_service_intf/mine_intf"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	var port string
	port = system_config.Configer().String("httpport")
	port = ":" + port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logs.Error("failed to listen: %v", err)
	}
	log.Printf("impl is started: listen port: %s", string(port))
	s := grpc.NewServer()
	pooService := impl.PoolImpl{}
	mine_intf.RegisterPoolServer(s, pooService)
	if err := s.Serve(lis); err != nil {
		logs.Error("failed to serve: %v", err)
	}
}
