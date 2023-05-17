package data_storage

import (
	"google.golang.org/protobuf/types/known/durationpb"
)

type DataBase struct {
	Driver string `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
}

type RedisBase struct {
	Network      string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr         string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	User         string               `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Pwd          string               `protobuf:"bytes,4,opt,name=pwd,proto3" json:"pwd,omitempty"`
	Db           int32                `protobuf:"varint,5,opt,name=db,proto3" json:"db,omitempty"`
	ReadTimeout  *durationpb.Duration `protobuf:"bytes,6,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout *durationpb.Duration `protobuf:"bytes,7,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
}
