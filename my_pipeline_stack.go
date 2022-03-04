package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type MyPipelineStackProps struct {
	awscdk.StackProps
}

func NewMyPipelineStack(scope constructs.Construct, id string, props *MyPipelineStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here
	pipelines.NewCodePipeline(stack, jsii.String("Pipeline"), &pipelines.CodePipelineProps{
		PipelineName: jsii.String("MyPipeline"),
		Synth: pipelines.NewShellStep(jsii.String("Synth"), &pipelines.ShellStepProps{
			Commands: &[]*string{
				jsii.String("./cdk-cli-wrapper-dev.sh synth"),
			},
			Input: pipelines.CodePipelineSource_GitHub(jsii.String("cowcoa/my_pipeline"), jsii.String("main"), &pipelines.GitHubSourceOptions{}),
		}),
	})

	return stack
}
