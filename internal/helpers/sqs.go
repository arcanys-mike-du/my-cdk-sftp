package helpers

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type StandardQueueProps struct {
	DLQ               awssqs.Queue
	AlarmTopic        awssns.ITopic
	ID                string
	Name              string
	VisibilityTimeout int
	MaxReceiveCount   int
	MakeDLQ           bool
}

func NewStandardQueue(scope constructs.Construct, props StandardQueueProps) awssqs.Queue {
	var dlq awssqs.Queue
	if props.MakeDLQ && props.DLQ == nil {
		dlq = NewStandardDLQ(scope, StandardQueueProps{
			AlarmTopic:      props.AlarmTopic,
			ID:              props.ID + "DLQ",
			Name:            props.Name + "-dlq",
			MaxReceiveCount: props.MaxReceiveCount,
			MakeDLQ:         true,
		})
	} else if props.DLQ != nil {
		dlq = props.DLQ
	}

	var dlqProps *awssqs.DeadLetterQueue
	if dlq != nil {
		dlqProps = &awssqs.DeadLetterQueue{
			MaxReceiveCount: jsii.Number(props.MaxReceiveCount),
			Queue:           dlq,
		}
	}

	q := awssqs.NewQueue(scope, jsii.String(props.ID), &awssqs.QueueProps{
		DeadLetterQueue:   dlqProps,
		QueueName:         jsii.String(props.Name),
		RetentionPeriod:   awscdk.Duration_Days(jsii.Number(7)),
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(props.VisibilityTimeout)),
	})

	return q
}

func NewStandardDLQ(scope constructs.Construct, props StandardQueueProps) awssqs.Queue {
	q := awssqs.NewQueue(scope, jsii.String(props.ID), &awssqs.QueueProps{
		QueueName: jsii.String(props.Name),
	})

	SetupAlarm(scope, SetupAlarmProps{
		ID:   fmt.Sprintf("%sAlarm", props.ID),
		Name: fmt.Sprintf("%salarm", props.Name),
		Metric: q.MetricApproximateNumberOfMessagesVisible(&awscloudwatch.MetricOptions{
			Statistic: awscloudwatch.Stats_AVERAGE(),
			Period:    awscdk.Duration_Minutes(jsii.Number(5)),
		}),
		Topic:     props.AlarmTopic,
		Threshold: 1,
	})

	return q
}
