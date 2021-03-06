#!/bin/bash

# Get script location.
SHELL_PATH=$(cd "$(dirname "$0")";pwd)

CDK_CMD=$1
CDK_ACC="$(aws sts get-caller-identity --output text --query 'Account')"
CDK_REGION="$(aws configure get region)"

if [ -z $CDK_REGION ]
then
    CDK_REGION=$AWS_DEFAULT_REGION
fi

echo "CDK_ACC: $CDK_ACC"
echo "CDK_REGION: $CDK_REGION"

if [ -z $AWS_DEFAULT_REGION ]
then
    echo "Run bootstrap"
    export CDK_NEW_BOOTSTRAP=1 
    npx cdk bootstrap aws://${CDK_ACC}/${CDK_REGION} --cloudformation-execution-policies arn:aws:iam::aws:policy/AdministratorAccess
fi

# Deploy pre-process.
echo "Run cdk-cli-wrapper.sh"
$SHELL_PATH/cdk-cli-wrapper.sh ${CDK_ACC} ${CDK_REGION} "$@"

# Destroy post-process.
if [ "$CDK_CMD" == "destroy" ]; then
    rm -rf $SHELL_PATH/cdk.out/
fi
