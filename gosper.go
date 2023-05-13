package main

import (
	"gosper/lib"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type GosperStackProps struct {
	awscdk.StackProps
}

func NewGosperStack(scope constructs.Construct, id string, props *GosperStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	lib.NewChatFunc(stack, "ChatFunc")

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewGosperStack(app, "GosperStack", &GosperStackProps{
		awscdk.StackProps{},
	})

	app.Synth(nil)
}
