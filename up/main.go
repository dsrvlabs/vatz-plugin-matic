package main

import (
	"log"

	pluginpb "github.com/dsrvlabs/vatz-proto/plugin/v1"
	"github.com/dsrvlabs/vatz/sdk"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/dsrvlabs/vatz-plugin-matic/up/policy"
)

const (
	pluginName = "vatz-plugin-matic"
)

var (
	executor policy.Executor
)

func init() {
	executor = policy.NewExecutor()
}

func main() {
	ctx := context.Background()
	p := sdk.NewPlugin(pluginName)

	p.Register(pluginFeature)

	if err := p.Start(ctx, "0.0.0.0", 9091); err != nil {
		log.Fatal(err)
	}
}

func pluginFeature(info, option map[string]*structpb.Value) (sdk.CallResponse, error) {
	log.Println("Execute pluginFeature")

	log.Println(info)

	resp := sdk.CallResponse{}

	val, ok := info["execute_method"]
	if !ok {
		log.Println("No mandatory field")
		return resp, nil
	}

	switch val.GetStringValue() {
	case "isBorUp":
		isUp, err := executor.IsBorUp()
		if err != nil {
			resp.Message = err.Error()
			return resp, err
		}

		log.Println(isUp, err)

		if isUp {
			resp.Severity = pluginpb.SEVERITY_INFO
			resp.State = pluginpb.STATE_SUCCESS
			resp.Message = "up"
		} else {
			resp.Severity = pluginpb.SEVERITY_CRITICAL
			resp.State = pluginpb.STATE_FAILURE
			resp.Message = "down"
		}

	case "isHeimdallUp":
		isUp, err := executor.IsHeimdallUp()
		if err != nil {
			resp.Message = err.Error()
			return resp, err
		}

		log.Println(isUp, err)
		if isUp {
			resp.Severity = pluginpb.SEVERITY_INFO
			resp.State = pluginpb.STATE_SUCCESS
			resp.Message = "up"
		} else {
			resp.Severity = pluginpb.SEVERITY_CRITICAL
			resp.State = pluginpb.STATE_FAILURE
			resp.Message = "down"
		}

	case "isHeimdallRestUp":
		isUp, err := executor.IsHeimdallRestUp()
		if err != nil {
			resp.Message = err.Error()
			return resp, err
		}

		log.Println(isUp, err)
		if isUp {
			resp.Severity = pluginpb.SEVERITY_INFO
			resp.State = pluginpb.STATE_SUCCESS
			resp.Message = "up"
		} else {
			resp.Severity = pluginpb.SEVERITY_CRITICAL
			resp.State = pluginpb.STATE_FAILURE
			resp.Message = "down"
		}
	default:
		log.Println("No matching function")

		resp.Severity = pluginpb.SEVERITY_ERROR
		resp.State = pluginpb.STATE_FAILURE
		resp.Message = "no matching function"
		break
	}

	return resp, nil
}
