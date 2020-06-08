module github.com/anhntbk08/machinery

require (
	cloud.google.com/go/pubsub v1.0.1
	github.com/RichardKnop/logging v0.0.0-20190827224416-1a693bdd4fae
	github.com/RichardKnop/machinery v1.7.3 // indirect
	github.com/RichardKnop/redsync v1.2.0
	github.com/aws/aws-sdk-go v1.25.8
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/uuid v1.1.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271
	github.com/stretchr/testify v1.3.0
	github.com/urfave/cli v1.22.1
	go.mongodb.org/mongo-driver v1.1.2
	gopkg.in/yaml.v2 v2.2.4
)

replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999

go 1.13
