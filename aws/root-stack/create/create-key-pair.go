package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/halt"
	"github.com/jmckee46/deployer/logger"
)

// This file creates a key pair in aws used in the launch configuation stack template
// and stores it in ~/.ssh and modifies permissions.
func createKeyPair() flaw.Flaw {
	// get aws client
	ec2Cli := ec2.New(awsSession())

	// create input struct
	input := &ec2.CreateKeyPairInput{
		KeyName: aws.String("depler"),
	}

	// create the key pair
	output, err := ec2Cli.CreateKeyPair(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "InvalidKeyPair.Duplicate":
				return nil

			default:
				fmt.Println(aerr.Error())
			}
		} else {
			return flaw.From(err)
		}
	}

	// save key pair
	flawErr := saveKeyPair(*output.KeyMaterial)
	if flawErr != nil {
		return flawErr
	}

	// set permissions
	flawErr = setPermissions()
	if flawErr != nil {
		return flawErr
	}

	return nil
}

func awsSession() *session.Session {
	sess, err := session.NewSession()

	if err != nil {
		logger.Panic("create key pair", flaw.From(err))
	}

	_, err = sess.Config.Credentials.Get()

	if err != nil {
		halt.Panic(flaw.From(err).Wrap("cannot awsSession"))
	}

	return sess
}

func saveKeyPair(pem string) flaw.Flaw {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return flaw.From(err)
	}

	err = os.Chdir(homeDir + "/.ssh")
	if err != nil {
		return flaw.From(err)
	}

	f, err := os.Create("depler")
	if err != nil {
		return flaw.From(err)
	}

	_, err = f.WriteString(pem)
	if err != nil {
		f.Close()
		return flaw.From(err)
	}

	err = f.Close()
	if err != nil {
		return flaw.From(err)
	}

	return nil
}

func setPermissions() flaw.Flaw {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return flaw.From(err)
	}

	err = os.Chdir(homeDir + "/.ssh")
	if err != nil {
		return flaw.From(err)
	}

	err = os.Chmod("depler", 400)
	if err != nil {
		return flaw.From(err)
	}

	return nil
}
