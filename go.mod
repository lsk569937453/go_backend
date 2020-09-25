module go_backend

go 1.14

replace (
	github.com/golang/protobuf => github.com/golang/protobuf v1.3.5
	github.com/jhump/protoreflect => github.com/jhump/protoreflect v1.7.0
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55
	//google.golang.org/protobuf => github.com/golang/protobuf v1.4.2
	google.golang.org/grpc => google.golang.org/grpc v1.31.1
)

require (
	github.com/Unknwon/goconfig v0.0.0-20200817131228-2444c9802e76
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.2.0
	github.com/go-redis/redis/v8 v8.0.0-beta.10
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/jhump/protoreflect v1.7.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/robfig/cron/v3 v3.0.1
	github.com/satori/go.uuid v1.2.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/valyala/fasthttp v1.16.0
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.31.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)
