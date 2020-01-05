module miner_service_impl

go 1.13

replace github.com/fogetu/miner_service_impl => ./

require (
	github.com/fogetu/miner_service_impl v0.0.0-00010101000000-000000000000
	github.com/fogetu/miner_service_intf v0.0.0-20181011133534-7bec62e10828
	github.com/jinzhu/gorm v1.9.11
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553
	google.golang.org/grpc v1.26.0
)
