package main

import "time"

type SNSRecords struct {
	Records []struct {
		EventSource          string `json:"EventSource"`
		EventVersion         string `json:"EventVersion"`
		EventSubscriptionArn string `json:"EventSubscriptionArn"`
		SNS                  struct {
			Type              string    `json:"Type"`
			MessageID         string    `json:"MessageId"`
			TopicArn          string    `json:"TopicArn"`
			Subject           string    `json:"Subject"`
			Message           string    `json:"Message"`
			Timestamp         time.Time `json:"Timestamp"`
			SignatureVersion  string    `json:"SignatureVersion"`
			Signature         string    `json:"Signature"`
			SigningCertURL    string    `json:"SigningCertUrl"`
			UnsubscribeURL    string    `json:"UnsubscribeUrl"`
			MessageAttributes struct {
			} `json:"MessageAttributes"`
		} `json:"Sns"`
	} `json:"Records"`
}

type CloudWatchAlarmMessage struct {
	AlarmName        string `json:"AlarmName"`
	AlarmDescription string `json:"AlarmDescription"`
	AWSAccountID     string `json:"AWSAccountId"`
	NewStateValue    string `json:"NewStateValue"`
	NewStateReason   string `json:"NewStateReason"`
	StateChangeTime  string `json:"StateChangeTime"`
	Region           string `json:"Region"`
	OldStateValue    string `json:"OldStateValue"`
	Trigger          struct {
		MetricName string `json:"MetricName"`
		Namespace  string `json:"Namespace"`
		Statistic  string `json:"Statistic"`
		Dimensions []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"Dimensions"`
		Period             float64 `json:"Period"`
		EvaluationPeriods  float64 `json:"EvaluationPeriods"`
		ComparisonOperator string  `json:"ComparisonOperator"`
		Threshold          float64 `json:"Threshold"`
	} `json:"Trigger"`
}
