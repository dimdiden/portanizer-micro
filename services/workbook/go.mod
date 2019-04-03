module github.com/dimdiden/portanizer-micro/services/workbook

go 1.12

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dimdiden/portanizer-micro/services/users v0.0.0-20190321143526-584168598b56
	github.com/go-kit/kit v0.8.0
	github.com/golang/protobuf v1.3.1
	github.com/kelseyhightower/envconfig v1.3.0
	github.com/lib/pq v1.0.0
	github.com/oklog/run v1.0.0
	google.golang.org/grpc v1.19.1
)

replace github.com/dimdiden/portanizer-micro/services/users => ../users
