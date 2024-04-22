package helpers

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

var buildFlags = jsii.Strings("-ldflags \"-s -w\"", "-trimpath", "-tags lambda.norpc")

type GoFunctionProps struct {
	EnvVars        *map[string]*string
	ID             string
	FunctionName   string
	EntryPath      string
	TimeoutSeconds int
	Role           awsiam.Role
}

func GoFunc(scope constructs.Construct, props GoFunctionProps) awscdklambdagoalpha.GoFunction {
	return awscdklambdagoalpha.NewGoFunction(scope, jsii.String(props.ID), &awscdklambdagoalpha.GoFunctionProps{
		FunctionName: jsii.String(props.FunctionName),
		Entry:        jsii.String(props.EntryPath),
		MemorySize:   jsii.Number(1024),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(props.TimeoutSeconds)),
		Tracing:      awslambda.Tracing_ACTIVE,
		Architecture: awslambda.Architecture_ARM_64(),
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			GoBuildFlags: buildFlags,
		},
		Environment:  props.EnvVars,
		LogRetention: awslogs.RetentionDays_THREE_MONTHS,
		Role:         props.Role,
	})
}
