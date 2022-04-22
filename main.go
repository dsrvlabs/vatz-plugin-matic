package main

import (
	"log"

	"github.com/dsrvlabs/vatz/sdk"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/dsrvlabs/vatz-plugin-matic/policy"
)

var (
	executor policy.Executor
)

func init() {
	executor = policy.NewExecutor()
}

func main() {
	ctx := context.Background()
	p := sdk.NewPlugin()

	p.Register(pluginFeature)

	if err := p.Start(ctx, "0.0.0.0", 9091); err != nil {
		log.Fatal(err)
	}
}

func pluginFeature(info, option map[string]*structpb.Value) error {
	log.Println("Execute pluginFeature")

	log.Println(info)

	val, ok := info["execute_method"]
	if !ok {
		log.Println("No mandatory field")
		return nil
	}

	switch val.GetStringValue() {
	case "IsBorUp":
		isUp, err := executor.IsBorUp()
		if err != nil {
			return err
		}
		_ = isUp
	default:
		log.Println("No matching function")
		break
	}

	return nil
}
