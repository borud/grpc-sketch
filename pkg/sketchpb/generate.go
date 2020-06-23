package sketchpb

//go:generate protoc -I=../../api/proto/v1 --go_out=plugins=grpc:. ../../api/proto/v1/sketch.proto
