package imagex_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	imagex "github.com/volcengine/volc-sdk-golang/service/imagex/v2"
)

func Test_DescribeImageXMirrorRequestHTTPCodeByTime(t *testing.T) {
	instance := imagex.NewInstance()

	instance.SetCredential(base.Credentials{
		AccessKeyID:     "ak",
		SecretAccessKey: "sk",
	})

	param := &imagex.DescribeImageXMirrorRequestHTTPCodeByTimeReq{
		DescribeImageXMirrorRequestHTTPCodeByTimeQuery: &imagex.DescribeImageXMirrorRequestHTTPCodeByTimeQuery{},
		DescribeImageXMirrorRequestHTTPCodeByTimeBody:  &imagex.DescribeImageXMirrorRequestHTTPCodeByTimeBody{},
	}

	resp, err := instance.DescribeImageXMirrorRequestHTTPCodeByTime(context.Background(), param)

	if err != nil {
		fmt.Printf("error %v", err)
	} else {
		t, _ := json.Marshal(resp)
		fmt.Printf("success %v", string(t))
	}
}
