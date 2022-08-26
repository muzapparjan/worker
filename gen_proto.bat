@echo off
call protoc --go_out=plugins=grpc:. proto\worker.proto