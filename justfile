run cmd: (build cmd)
    @./bin/{{cmd}}

build cmd:
    @go build -o ./bin/{{cmd}} ./cmd/{{cmd}}/*
