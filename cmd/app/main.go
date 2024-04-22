package main

import (
	awscdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
	"my-cdk-sftp/constants"
	"my-cdk-sftp/lib/my_cdk_sftp_stack"
)

func main() {
	app := awscdk.NewApp(nil)

	my_cdk_sftp_stack.NewMyCdkSftpStack(app, "MyCdkSftpStack", &my_cdk_sftp_stack.MyCdkSftpStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(constants.Account),
		Region:  jsii.String(constants.Region),
	}
}
