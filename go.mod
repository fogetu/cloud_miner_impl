module miner_service_impl

go 1.13

replace github.com/fogetu/miner_service_impl => ./

require github.com/fogetu/miner_service_intf v0.0.0-20200113151624-96513cf11c88

require (
	github.com/astaxie/beego v1.12.0
	github.com/fogetu/go_system v0.0.0-20200117024926-65de75c476fe
	github.com/fogetu/miner_service_impl v0.0.0-00010101000000-000000000000
	github.com/jinzhu/gorm v1.9.11
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553
	google.golang.org/grpc v1.26.0
)
