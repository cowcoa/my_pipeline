package lib

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func NewMyPipelineAppStage(scope constructs.Construct, id *string, props *awscdk.StageProps) awscdk.Stage {

	stage := awscdk.NewStage(scope, id, props)
	NewMyLambdaStack(stage, "LambdaStack", nil)

	return stage
}
