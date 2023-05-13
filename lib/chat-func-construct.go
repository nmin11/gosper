package lib

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewChatFunc(scope constructs.Construct, id string) awslambda.Function {
	chatFunc := awslambda.NewFunction(scope, jsii.String(id), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Code:    awslambda.AssetCode_FromAsset(jsii.String("../server/lambda"), &awss3assets.AssetOptions{}),
		Handler: jsii.String("main"),
	})

	return chatFunc
}
