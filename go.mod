module github.com/lorgioedtech/go-rest-dev

go 1.16

require (
	github.com/google/uuid v1.3.0
	github.com/jcelliott/lumber v0.0.0-20160324203708-dd349441af25 // indirect
	github.com/julienschmidt/httprouter v1.3.0
	github.com/nanobox-io/golang-scribble v0.0.0-20190309225732-aa3e7c118975
	github.com/qiangxue/go-env v1.0.1
	go.uber.org/zap v1.18.1
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/lorgioedtech/go-rest-dev/internal/http/rest => ./internal/http/rest

replace github.com/lorgioedtech/go-rest-dev/internal/storage/memory => ./internal/storage/memory

replace github.com/lorgioedtech/go-rest-dev/internal/storage/json => ./internal/storage/json

replace github.com/lorgioedtech/go-rest-dev/internal/listing => ./internal/listing

replace github.com/lorgioedtech/go-rest-dev/internal/adding => ./internal/adding
