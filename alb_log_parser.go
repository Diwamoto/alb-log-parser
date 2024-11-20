package main

import (
	"errors"
	"regexp"
)

type AlbLogParser struct {
	regexp regexp.Regexp
}

func NewAlbLogParser() AlbLogParser {
	return AlbLogParser{
		regexp: *regexp.MustCompile(ALB_LOG_REGEXP_STRING),
	}
}

// ParseAlbLog parses a given ALB log string and returns a slice of strings
// containing the matched groups based on the ALB_LOG_REGEXP_STRING regular expression.
// If no matches are found, it returns an empty slice and no error.
//
// Parameters:
//   - log: A string containing the ALB log to be parsed.
//
// Returns:
//   - A slice of strings containing the matched groups from the log.
//   - An error if the parsing fails (currently always returns nil).
func (p AlbLogParser) ParseAlbLog(log string) (AlbLogRecord, error) {

	// Ignore the 0th element as it contains the entire matched string
	matches := p.regexp.FindStringSubmatch(log)
	if matches == nil {
		return AlbLogRecord{}, errors.New("no matches found")
	}
	matches = matches[1:]

	albLog := AlbLogRecord{
		Type:                   matches[0],
		Time:                   matches[1],
		Elb:                    matches[2],
		ClientIP:               matches[3],
		ClientPort:             matches[4],
		TargetIP:               matches[5],
		TargetPort:             matches[6],
		RequestProcessingTime:  matches[7],
		TargetProcessingTime:   matches[8],
		ResponseProcessingTime: matches[9],
		ElbStatusCode:          matches[10],
		TargetStatusCode:       matches[11],
		ReceivedBytes:          matches[12],
		SentBytes:              matches[13],
		HttpMethod:             matches[14],
		RequestUrl:             matches[15],
		HttpVersion:            matches[16],
		UserAgent:              matches[17],
		SslCipher:              matches[18],
		SslProtocol:            matches[19],
		TargetGroupArn:         matches[20],
		TraceId:                matches[21],
		DomainName:             matches[22],
		ChosenCertArn:          matches[23],
		MatchedRulePriority:    matches[24],
		RequestCreationTime:    matches[25],
		ActionsExecuted:        matches[26],
		RedirectUrl:            matches[27],
		ErrorReason:            matches[28],
		TargetPortList:         matches[29],
		TargetStatusCodeList:   matches[30],
		Classification:         matches[31],
		ClassificationReason:   matches[32],
		ConnTraceId:            matches[33],
	}

	// Check if there is an unknown field
	if matches[34] != "" {
		albLog.UnknownField = matches[34]
	}

	return albLog, nil

}
