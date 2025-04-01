#!/bin/sh

protoc --go_out=. --go-grpc_out=. shop_product.proto payment.proto
