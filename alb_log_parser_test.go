package alb_log_parser

import (
	"reflect"
	"testing"
)

func TestAlbLogParser_ParseAlbLog(t *testing.T) {
	type args struct {
		log string
	}
	tests := []struct {
		name    string
		p       AlbLogParser
		args    args
		want    AlbLogRecord
		wantErr bool
	}{
		{
			name: "Positive Case1 (without UnknownField)",
			p:    NewAlbLogParser(),
			args: args{
				log: `https 2024-11-20T03:50:02.120802Z app/test-app-alb/1234567890abcdef 192.168.0.1:12345 10.0.0.1:80 0.002 0.035 0.000 200 200 264 868 "GET https://example.com:443/api/test HTTP/1.1" "Test-User-Agent" ECDHE-RSA-AES128-GCM-SHA256 TLSv1.2 arn:aws:elasticloadbalancing:region:account-id:targetgroup/test-tg/abcdef123456 "Root=1-2-3-4-5" "example.com" "arn:aws:acm:region:account-id:certificate/abcdef123456" 0 2024-11-20T03:50:02.083000Z "forward" "-" "-" "10.0.0.1:80" "200" "-" "-" TID_test`,
			},
			want: AlbLogRecord{
				Type:                   "https",
				Time:                   "2024-11-20T03:50:02.120802Z",
				Elb:                    "app/test-app-alb/1234567890abcdef",
				ClientIP:               "192.168.0.1",
				ClientPort:             "12345",
				TargetIP:               "10.0.0.1",
				TargetPort:             "80",
				RequestProcessingTime:  "0.002",
				TargetProcessingTime:   "0.035",
				ResponseProcessingTime: "0.000",
				ElbStatusCode:          "200",
				TargetStatusCode:       "200",
				ReceivedBytes:          "264",
				SentBytes:              "868",
				HttpMethod:             "GET",
				RequestUrl:             "https://example.com:443/api/test",
				HttpVersion:            "HTTP/1.1",
				UserAgent:              "Test-User-Agent",
				SslCipher:              "ECDHE-RSA-AES128-GCM-SHA256",
				SslProtocol:            "TLSv1.2",
				TargetGroupArn:         "arn:aws:elasticloadbalancing:region:account-id:targetgroup/test-tg/abcdef123456",
				TraceId:                "Root=1-2-3-4-5",
				DomainName:             "example.com",
				ChosenCertArn:          "arn:aws:acm:region:account-id:certificate/abcdef123456",
				MatchedRulePriority:    "0",
				RequestCreationTime:    "2024-11-20T03:50:02.083000Z",
				ActionsExecuted:        "forward",
				RedirectUrl:            "-",
				ErrorReason:            "-",
				TargetPortList:         "10.0.0.1:80",
				TargetStatusCodeList:   "200",
				Classification:         "-",
				ClassificationReason:   "-",
				ConnTraceId:            "TID_test",
				UnknownField:           "",
			},
			wantErr: false,
		},
		{
			name: "Positive Case 2 (with one UnknownField)",
			p:    NewAlbLogParser(),
			args: args{
				log: `https 2024-11-20T03:50:02.120802Z app/test-app-alb/1234567890abcdef 192.168.0.1:12345 10.0.0.1:80 0.002 0.035 0.000 200 200 264 868 "GET https://example.com:443/api/test HTTP/1.1" "Test-User-Agent" ECDHE-RSA-AES128-GCM-SHA256 TLSv1.2 arn:aws:elasticloadbalancing:region:account-id:targetgroup/test-tg/abcdef123456 "Root=1-2-3-4-5" "example.com" "arn:aws:acm:region:account-id:certificate/abcdef123456" 0 2024-11-20T03:50:02.083000Z "forward" "-" "-" "10.0.0.1:80" "200" "-" "-" TID_test new_field`,
			},
			want: AlbLogRecord{
				Type:                   "https",
				Time:                   "2024-11-20T03:50:02.120802Z",
				Elb:                    "app/test-app-alb/1234567890abcdef",
				ClientIP:               "192.168.0.1",
				ClientPort:             "12345",
				TargetIP:               "10.0.0.1",
				TargetPort:             "80",
				RequestProcessingTime:  "0.002",
				TargetProcessingTime:   "0.035",
				ResponseProcessingTime: "0.000",
				ElbStatusCode:          "200",
				TargetStatusCode:       "200",
				ReceivedBytes:          "264",
				SentBytes:              "868",
				HttpMethod:             "GET",
				RequestUrl:             "https://example.com:443/api/test",
				HttpVersion:            "HTTP/1.1",
				UserAgent:              "Test-User-Agent",
				SslCipher:              "ECDHE-RSA-AES128-GCM-SHA256",
				SslProtocol:            "TLSv1.2",
				TargetGroupArn:         "arn:aws:elasticloadbalancing:region:account-id:targetgroup/test-tg/abcdef123456",
				TraceId:                "Root=1-2-3-4-5",
				DomainName:             "example.com",
				ChosenCertArn:          "arn:aws:acm:region:account-id:certificate/abcdef123456",
				MatchedRulePriority:    "0",
				RequestCreationTime:    "2024-11-20T03:50:02.083000Z",
				ActionsExecuted:        "forward",
				RedirectUrl:            "-",
				ErrorReason:            "-",
				TargetPortList:         "10.0.0.1:80",
				TargetStatusCodeList:   "200",
				Classification:         "-",
				ClassificationReason:   "-",
				ConnTraceId:            "TID_test",
				UnknownField:           "new_field",
			},
			wantErr: false,
		},
		{
			name: "Positive Case 3 (with some UnknownFields)",
			p:    NewAlbLogParser(),
			args: args{
				log: `https 2024-11-20T03:50:02.120802Z app/test-app-alb/1234567890abcdef 192.168.0.1:12345 10.0.0.1:80 0.002 0.035 0.000 200 200 264 868 "GET https://example.com:443/api/test HTTP/1.1" "Test-User-Agent" ECDHE-RSA-AES128-GCM-SHA256 TLSv1.2 arn:aws:elasticloadbalancing:region:account-id:targetgroup/test-tg/abcdef123456 "Root=1-2-3-4-5" "example.com" "arn:aws:acm:region:account-id:certificate/abcdef123456" 0 2024-11-20T03:50:02.083000Z "forward" "-" "-" "10.0.0.1:80" "200" "-" "-" TID_test new_field1 "new_field2"`,
			},
			want: AlbLogRecord{
				Type:                   "https",
				Time:                   "2024-11-20T03:50:02.120802Z",
				Elb:                    "app/test-app-alb/1234567890abcdef",
				ClientIP:               "192.168.0.1",
				ClientPort:             "12345",
				TargetIP:               "10.0.0.1",
				TargetPort:             "80",
				RequestProcessingTime:  "0.002",
				TargetProcessingTime:   "0.035",
				ResponseProcessingTime: "0.000",
				ElbStatusCode:          "200",
				TargetStatusCode:       "200",
				ReceivedBytes:          "264",
				SentBytes:              "868",
				HttpMethod:             "GET",
				RequestUrl:             "https://example.com:443/api/test",
				HttpVersion:            "HTTP/1.1",
				UserAgent:              "Test-User-Agent",
				SslCipher:              "ECDHE-RSA-AES128-GCM-SHA256",
				SslProtocol:            "TLSv1.2",
				TargetGroupArn:         "arn:aws:elasticloadbalancing:region:account-id:targetgroup/test-tg/abcdef123456",
				TraceId:                "Root=1-2-3-4-5",
				DomainName:             "example.com",
				ChosenCertArn:          "arn:aws:acm:region:account-id:certificate/abcdef123456",
				MatchedRulePriority:    "0",
				RequestCreationTime:    "2024-11-20T03:50:02.083000Z",
				ActionsExecuted:        "forward",
				RedirectUrl:            "-",
				ErrorReason:            "-",
				TargetPortList:         "10.0.0.1:80",
				TargetStatusCodeList:   "200",
				Classification:         "-",
				ClassificationReason:   "-",
				ConnTraceId:            "TID_test",
				UnknownField:           "new_field1 \"new_field2\"",
			},
			wantErr: false,
		},
		{
			name: "Negative Case 4 (invalid log format)",
			p:    NewAlbLogParser(),
			args: args{
				log: `invalid log format`,
			},
			want:    AlbLogRecord{},
			wantErr: true,
		},
		{
			name: "Negative Case 5 (missing fields)",
			p:    NewAlbLogParser(),
			args: args{
				log: `https 2024-11-20T03:50:02.120802Z app/test-app-alb/1234567890abcdef`,
			},
			want:    AlbLogRecord{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.ParseAlbLog(tt.args.log)
			if (err != nil) != tt.wantErr {
				t.Errorf("AlbLogParser.ParseAlbLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AlbLogParser.ParseAlbLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
