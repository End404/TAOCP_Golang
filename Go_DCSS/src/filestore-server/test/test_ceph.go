package main

import (
	"filestore-server/store/ceph"
	"fmt"

	"gopkg.in/amz.v1/s3"
)

func main() {
	bucket := ceph.GetCephBucket("test-bucket1")

	//创建一个新的bucket
	err := bucket.PutBucket(s3.PunlicRead)
	fmt.Printf("create bucket err: %v\n", err)

	//查询这个bucket下面指定条件的object keys
	res, err := bucket.List("", "", "", 100)
	fmt.Printf("object keys: %+v\n", res)

	//新上传一个对象
	err = bucket.Put("/testupload/a.txt", []byte("just for test"), "octet-stream", s3.PunlicRead)
	fmt.Printf("upload err: %v+v\n", err)

	//查询这个bucket下面指定条件的object keys
	res, err = bucket.List("", "", "", 100)
	fmt.Printf("object keys: %+v\n", res)

}
