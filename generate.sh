#!/bin/bash

protoc -I ./proto/ \
    -I include/googleapis -I include/grpc-gateway \
    --go_out=. --go_opt=module=github.com/bbengfort/notes \
    --go-grpc_out=. --go-grpc_opt=module=github.com/bbengfort/notes \
    --openapiv2_out ./openapiv2 --openapiv2_opt logtostderr=true \
    proto/notes/v1/*.proto
