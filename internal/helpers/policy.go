package helpers

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
)

// AddFunctionPolicy helper to add a function policy to the lambda
func AddFunctionPolicy(fn awslambda.IFunction, actions []string, resources []string) {
	fn.AddToRolePolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect:    awsiam.Effect_ALLOW,
		Actions:   jsii.Strings(actions...),
		Resources: jsii.Strings(resources...),
	}))
}

// AddGetSecretPolicy helper to create and add Get Secret Value permissions to a lambda
func AddGetSecretPolicy(stack awscdk.Stack, fn awslambda.IFunction, secretName string) {
	arn := SecretARN(stack, secretName)
	AddFunctionPolicy(fn,
		[]string{"secretsmanager:GetSecretValue"},
		[]string{*arn},
	)
}

// AddStateMachinePolicy helper to add state machine permissions to a lambda
func AddStateMachinePolicy(stack awscdk.Stack, fn awslambda.IFunction, stateMachineName string, permissions ...string) {
	arn := StateMachineARN(stack, stateMachineName)
	AddFunctionPolicy(fn,
		permissions,
		[]string{*arn},
	)
}
