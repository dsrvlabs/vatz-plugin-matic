package main

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall"
	"testing"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/dsrvlabs/vatz-plugin-matic/mocks"
	pb "github.com/dsrvlabs/vatz-plugin-matic/plugin"
	"github.com/stretchr/testify/assert"
)

func TestGrpc(t *testing.T) {
	ch := make(chan os.Signal, 1)

	go func() {
		time.Sleep(time.Millisecond * 100)
		ch <- syscall.SIGTERM
	}()

	startServer(ch)
}

// TODO: Test invalid parameters.
func TestExecuteUp(t *testing.T) {
	// Mock
	mockExecutor := mocks.Executor{}
	executor = &mockExecutor

	mockExecutor.On("IsBorUp").Return(true, nil)

	// Prepare test
	s := pluginServer{}

	ctx := context.Background()
	req := &pb.ExecuteRequest{
		ExecuteInfo: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"function": structpb.NewStringValue("IsBorUp"),
			},
		},
	}

	// Test
	resp, err := s.Execute(ctx, req)

	// Asserts
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	mockExecutor.AssertExpectations(t)
}

func TestRequestJson(t *testing.T) {
	req := pb.ExecuteRequest{
		ExecuteInfo: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"function": structpb.NewStringValue("IsBorUp"),
			},
		},
	}

	d, err := json.Marshal(&req)

	fmt.Println(err)
	fmt.Println(string(d))
}
