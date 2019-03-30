module github.com/dimdiden/portanizer-micro/services/auth

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dimdiden/portanizer-micro/services/users v0.0.0-20190321143526-584168598b56
	github.com/go-kit/kit v0.8.0
	github.com/golang/protobuf v1.3.0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/kelseyhightower/envconfig v1.3.0
	github.com/oklog/run v1.0.0
	google.golang.org/grpc v1.19.0
)

replace github.com/dimdiden/portanizer-micro/services/users => ../users
