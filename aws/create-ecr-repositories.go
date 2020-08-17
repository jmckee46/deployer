package awsfuncs

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/jmckee46/deployer/images"

	"github.com/jmckee46/deployer/flaw"
)

// CreateEcrRepositories creates ecr repositories
func CreateEcrRepositories(state *State) flaw.Flaw {
	fmt.Println("    creating ECR repositories...")

	// get aws client
	ecrCli := state.AWSClient.ECR

	// get list of images ready to deploy
	images, err := images.List()
	if err != nil {
		return err
	}

	state.ImagesToDeploy = images

	for _, image := range images {
		repositoryName := os.Getenv("DE_STACK_NAME") + "-" + image

		// create input struct
		input := &ecr.CreateRepositoryInput{
			RepositoryName: aws.String(repositoryName),
		}

		// create ecr repository
		_, err := ecrCli.CreateRepository(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case ecr.ErrCodeRepositoryAlreadyExistsException:
					return nil
				default:
					fmt.Println(aerr.Error())
					return flaw.From(err)
				}
			} else {
				return flaw.From(err)
			}
		}
	}

	return nil
}
