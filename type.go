package main

const ALB_LOG_REGEXP_STRING = `([^ ]*) ([^ ]*) ([^ ]*) ([^ ]*):([0-9]*) ([^ ]*)[:-]([0-9]*) ([-.0-9]*) ([-.0-9]*) ([-.0-9]*) (|[-0-9]*) (-|[-0-9]*) ([-0-9]*) ([-0-9]*) \"([^ ]*) (.*) (- |[^ ]*)\" \"([^\"]*)\" ([A-Z0-9-_]+) ([A-Za-z0-9.-]*) ([^ ]*) \"([^\"]*)\" \"([^\"]*)\" \"([^\"]*)\" ([-.0-9]*) ([^ ]*) \"([^\"]*)\" \"([^\"]*)\" \"([^ ]*)\" \"([^s]+?)\" \"([^s]+)\" \"([^ ]*)\" \"([^ ]*)\" ([^ ]*) ?(.*)?`

// https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-access-logs.html
// AlbLogRecord represents a record in the ALB log with various fields capturing
// details about the request and response processed by the ALB.

type AlbLogRecord struct {
	Type                   string
	Time                   string
	Elb                    string
	ClientIP               string
	ClientPort             string
	TargetIP               string
	TargetPort             string
	RequestProcessingTime  string
	TargetProcessingTime   string
	ResponseProcessingTime string
	ElbStatusCode          string
	TargetStatusCode       string
	ReceivedBytes          string
	SentBytes              string
	HttpMethod             string
	RequestUrl             string
	HttpVersion            string
	UserAgent              string
	SslCipher              string
	SslProtocol            string
	TargetGroupArn         string
	TraceId                string
	DomainName             string
	ChosenCertArn          string
	MatchedRulePriority    string
	RequestCreationTime    string
	ActionsExecuted        string
	RedirectUrl            string
	ErrorReason            string
	TargetPortList         string
	TargetStatusCodeList   string
	Classification         string
	ClassificationReason   string
	ConnTraceId            string
	UnknownField           string // The UnknownField is used to automatically store any newly added unknown columns.
}
