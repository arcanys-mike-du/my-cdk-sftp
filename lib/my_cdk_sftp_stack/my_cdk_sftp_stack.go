package my_cdk_sftp_stack

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awstransfer"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type MyCdkSftpStackProps struct {
	awscdk.StackProps
}

func NewMyCdkSftpStack(scope constructs.Construct, id string, props *MyCdkSftpStackProps) awscdk.Stack {
	var s awscdk.Stack
	s = awscdk.NewStack(scope, &id, &props.StackProps)

	// Create an S3 bucket for use using
	awss3.NewBucket(s, jsii.String("MySftpBucket"), &awss3.BucketProps{
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY, // Adjust the removal policy as needed
	})

	// Create an IAM role for the SFTP server
	awsiam.NewRole(s, jsii.String("SftpRole"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("transfer.amazonaws.com"), nil),
		ManagedPolicies: &[]awsiam.IManagedPolicy{
			awsiam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("AmazonS3FullAccess")),
		},
	})

	// Create the SFTP server
	awstransfer.NewCfnServer(s, jsii.String("MySftpServer"), &awstransfer.CfnServerProps{
		EndpointType:         jsii.String("PUBLIC"),
		IdentityProviderType: jsii.String("SERVICE_MANAGED"),
		Protocols: &[]*string{
			jsii.String("SFTP"),
		},
	})
	return s
}
