package folgonsqproducer

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	nsq "github.com/nsqio/go-nsq"
)

// activityLog is the default logger for the exec Activity
var log = logger.GetLogger("activity-tibco-folgonsqproducer")
var producer = newProducer("127.0.0.1:4150", "127.0.0.1:4151")

type NsqpActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &NsqpActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *NsqpActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *NsqpActivity) Eval(context activity.Context) (done bool, err error) {

	topic, _ := context.GetInput("Topic").(string)
	message, _ := context.GetInput("Message").(string)

	producer.Publish(topic, []byte(message))
	log.Info("Published message")

	return true, nil
}

func newProducer(NsqdTcpAddress string, NsqdHttpAddress string) *nsq.Producer {
	config := nsq.NewConfig()
	newProducer, err := nsq.NewProducer(NsqdTcpAddress, config)
	if err != nil {
		log.Info("Could not create NSQ producer. Exiting.")
	}
	return newProducer
}
