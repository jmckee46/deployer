package awsclient

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/halt"
	"github.com/jmckee46/deployer/logger"
)

type Client struct {
	Cloudform *cloudformation.CloudFormation
	S3        *s3.S3
	ACM       *acm.ACM
}

func new() *Client {
	logger.Debug("awsclient-new", nil)

	var client *Client

	client = &Client{
		Cloudform: cloudformation.New(awsSession()),
		S3:        s3.New(awsSession()),
		ACM:       acm.New(awsSession()),
	}

	return client
}

func awsSession() *session.Session {
	sess, err := session.NewSession()

	if err != nil {
		logger.Panic("awsclient-aws-session", flaw.From(err))
	}

	_, err = sess.Config.Credentials.Get()

	if err != nil {
		halt.Panic(flaw.From(err).Wrap("cannot awsSession"))
	}

	return sess
}
