package pipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"my_pipeline/lib"
)

func NewMyPipelineStack(scope constructs.Construct, id string, props *awscdk.StackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, props)

	const connectionArn = "arn:aws:codestar-connections:ap-northeast-2:168228779762:connection/56f836ec-68cf-48be-a528-0f4e93544ceb"

	connStatement := awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect: awsiam.Effect_ALLOW,
		Actions: &[]*string{
			jsii.String("codestar-connections:UseConnection"),
		},
		Resources: &[]*string{
			jsii.String(connectionArn),
		},
	})

	// The code that defines your stack goes here
	myPipeline := pipelines.NewCodePipeline(stack, jsii.String("Pipeline"), &pipelines.CodePipelineProps{
		PipelineName: jsii.String("MyPipeline"),
		CodeBuildDefaults: &pipelines.CodeBuildOptions{
			RolePolicy: &[]awsiam.PolicyStatement{
				connStatement,
			},
			BuildEnvironment: &awscodebuild.BuildEnvironment{
				BuildImage:  awscodebuild.LinuxBuildImage_AMAZON_LINUX_2_3(),
				ComputeType: awscodebuild.ComputeType_MEDIUM,
			},
		},
		Synth: pipelines.NewShellStep(jsii.String("Synth"), &pipelines.ShellStepProps{
			Commands: &[]*string{
				jsii.String("source ./setup_build_env.sh"),
				jsii.String("echo $PATH"),
				jsii.String("echo $GOROOT"),
				jsii.String("echo $GOPATH"),
				jsii.String("go version"),
				jsii.String("node --version"),
				//jsii.String("cdk version"),
				//jsii.String("nvm current"),
				jsii.String("go mod tidy"),
				jsii.String("./cdk-cli-wrapper-dev.sh synth"),
			},
			//Input: pipelines.CodePipelineSource_GitHub(jsii.String("cowcoa/my_pipeline"), jsii.String("main"), &pipelines.GitHubSourceOptions{}),
			Input: pipelines.CodePipelineSource_Connection(jsii.String("cowcoa/my_pipeline"), jsii.String("main"), &pipelines.ConnectionSourceOptions{
				ConnectionArn: jsii.String(connectionArn),
			}),
		}),
	})

	testingStage := myPipeline.AddStage(lib.NewMyPipelineAppStage(stack, jsii.String("test"), &awscdk.StageProps{
		Env: props.Env,
	}), &pipelines.AddStageOpts{})

	testingStage.AddPost(pipelines.NewManualApprovalStep(jsii.String("approval"), &pipelines.ManualApprovalStepProps{}))

	return stack
}
