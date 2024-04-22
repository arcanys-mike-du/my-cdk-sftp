package my_cdk_sftp_stack

import (
	awscdk "github.com/aws/aws-cdk-go/awscdk/v2"
	awsiam "github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	awstransfer "github.com/aws/aws-cdk-go/awscdk/v2/awstransfer"
	constructs "github.com/aws/constructs-go/constructs/v10"
	jsii "github.com/aws/jsii-runtime-go"
	"my-cdk-sftp/contants"
	"my-cdk-sftp/helpers"
)

type MyCdkSftpStackProps struct {
	awscdk.StackProps
}

func NewMyCdkSftpStack(scope constructs.Construct, id string, props *MyCdkSftpStackProps) awscdk.Stack {
	var s awscdk.Stack
	s = awscdk.NewStack(scope, &id, &props.StackProps)

	// Create an S3 bucket for use
	awss3.NewBucket(s, jsii.String("MySftpBucket"), &awss3.BucketProps{
		BucketName:    jsii.String(contants.S3BucketName), // Use the constant here
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,       // Adjust the removal policy as needed
	})

	// Create an IAM role for the SFTP server
	role := awsiam.NewRole(s, jsii.String(contants.SftpRoleName), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("transfer.amazonaws.com"), nil),
		ManagedPolicies: &[]awsiam.IManagedPolicy{
			awsiam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("AmazonS3FullAccess")),
		},
	})

	// Create the IAM role for the Lambda function
	lambdaExecutionRole := awsiam.NewRole(s, jsii.String("CustomAuthorizerLambdaRole"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), nil),
		InlinePolicies: &map[string]awsiam.PolicyDocument{
			"inlinePolicy": awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
				Statements: &[]awsiam.PolicyStatement{
					awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
						Actions:   &[]*string{jsii.String("logs:CreateLogGroup"), jsii.String("logs:CreateLogStream"), jsii.String("logs:PutLogEvents")},
						Resources: &[]*string{jsii.String("*")},
						Effect:    awsiam.Effect_ALLOW,
					}),
				},
			}),
		},
	})

	lambdaFn := helpers.GoFunc(s, helpers.GoFunctionProps{
		ID:             "CustomAuthorizerFunction",
		FunctionName:   "CustomAuthorizerFunction",
		EntryPath:      "./cmd/lambdas/authorizer",
		TimeoutSeconds: 100,
		Role:           lambdaExecutionRole,
		EnvVars: &map[string]*string{
			"SFTP_ROLE_ARN":       role.RoleArn(),
			"S3_BUCKET_NAME":      jsii.String(contants.S3BucketName),
			"SFTP_HOME_DIRECTORY": jsii.String(contants.SftpHomeDirectory),
		},
	})

	// Use the Lambda function's ARN as the custom identity provider for the SFTP server
	awstransfer.NewCfnServer(s, jsii.String("MySftpServer"), &awstransfer.CfnServerProps{
		EndpointType:         jsii.String("PUBLIC"),
		IdentityProviderType: jsii.String("AWS_LAMBDA"), // Use AWS_LAMBDA for Lambda-based custom provider
		IdentityProviderDetails: &map[string]*string{
			"Function": lambdaFn.FunctionArn(),
			"Role":     lambdaExecutionRole.RoleArn(),
		},
		Protocols: &[]*string{
			jsii.String("SFTP"),
		},
		LoggingRole: role.RoleArn(),
	})

	return s
}
