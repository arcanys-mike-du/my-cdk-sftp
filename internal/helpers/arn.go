package helpers

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

// StateMachineARN helper to create an arn for a state machine
func StateMachineARN(stack awscdk.Stack, stateMachineName string) *string {
	return stack.FormatArn(&awscdk.ArnComponents{
		Account:      awscdk.Aws_ACCOUNT_ID(),
		Partition:    awscdk.Aws_PARTITION(),
		Region:       awscdk.Aws_REGION(),
		Service:      jsii.String("states"),
		Resource:     jsii.String("stateMachine"),
		ResourceName: jsii.String(stateMachineName),
		ArnFormat:    awscdk.ArnFormat_COLON_RESOURCE_NAME,
	})
}

// SecretARN helper to create an ARN for a secret.
// Either use the exact ARN or end the name with an asterisk (*)
func SecretARN(stack awscdk.Stack, name string) *string {
	return stack.FormatArn(&awscdk.ArnComponents{
		Partition:    awscdk.Aws_PARTITION(),
		Service:      jsii.String("secretsmanager"),
		Region:       awscdk.Aws_REGION(),
		Account:      awscdk.Aws_ACCOUNT_ID(),
		Resource:     jsii.String("secret"),
		ResourceName: jsii.String(name),
		ArnFormat:    awscdk.ArnFormat_COLON_RESOURCE_NAME,
	})
}
