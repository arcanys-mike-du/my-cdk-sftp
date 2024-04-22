package helpers

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatchactions"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type SetupAlarmProps struct {
	Metric    awscloudwatch.Metric
	Topic     awssns.ITopic
	ID        string
	Name      string
	Threshold float64
}

type StateMachineAlarmProps struct {
	Topic        awssns.ITopic
	StateMachine awsstepfunctions.StateMachine
	IDPrefix     string
	NamePrefix   string
}

func SetupAlarm(scope constructs.Construct, props SetupAlarmProps) {
	enableActions := props.Topic != nil
	alarm := props.Metric.CreateAlarm(scope, jsii.String(props.ID), &awscloudwatch.CreateAlarmOptions{
		EvaluationPeriods: jsii.Number(1),
		Threshold:         jsii.Number(props.Threshold),
		ActionsEnabled:    jsii.Bool(enableActions),
		AlarmName:         jsii.String(props.Name),
		TreatMissingData:  awscloudwatch.TreatMissingData_NOT_BREACHING,
	})

	if enableActions {
		alarm.AddAlarmAction(awscloudwatchactions.NewSnsAction(props.Topic))
		alarm.AddOkAction(awscloudwatchactions.NewSnsAction(props.Topic))
	}
}
