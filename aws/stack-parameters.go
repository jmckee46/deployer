package awsfuncs

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/jmckee46/deployer/flaw"
)

// StackParameters returns the stack parameters required to create/update a stack
func StackParameters(state *State) flaw.Flaw {
	fmt.Println("  gathering stack parameters...")

	arn, err := GetAcmCertificateArn()
	if err != nil {
		return err
	}

	gitShaParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("DeGitSha"),
		ParameterValue: aws.String(os.Getenv("DE_GIT_SHA")),
	}
	loadBalancerHostnameParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("DeLoadBalancerHostname"),
		ParameterValue: aws.String(os.Getenv("DE_LOAD_BALANCER_HOSTNAME")),
	}
	// stackBucketParameter := cloudformation.Parameter{
	// 	ParameterKey:   aws.String("DeStackBucket"),
	// 	ParameterValue: aws.String(os.Getenv("DE_STACK_BUCKET")),
	// }
	stackNameParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("DeStackName"),
		ParameterValue: aws.String(os.Getenv("DE_STACK_NAME")),
	}
	subnetCidrBlockParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("DeSubnetCidrBlocks"),
		ParameterValue: aws.String(os.Getenv("DE_SUBNET_CIDR_BLOCKS")),
	}
	tlsCertificationArnParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("DeTlsCertificateArn"),
		ParameterValue: aws.String(arn),
	}
	vpcCidrBaseParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("DeVpcCidrBase"),
		ParameterValue: aws.String(os.Getenv("DE_VPC_CIDR_BASE")),
	}
	dockerRegistry := cloudformation.Parameter{
		ParameterKey:   aws.String("DeDockerRegistry"),
		ParameterValue: aws.String(state.GetDockerRegistry()),
	}

	state.StackParametersStackState = []*cloudformation.Parameter{
		&gitShaParameter,
		&loadBalancerHostnameParameter,
		// &stackBucketParameter,
		&stackNameParameter,
		&subnetCidrBlockParameter,
		&tlsCertificationArnParameter,
		&vpcCidrBaseParameter,
		&dockerRegistry,
	}

	migrationsPgPasswordParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("DeMigrationsPgPassword"),
		ParameterValue: aws.String(os.Getenv("DE_MIGRATIONS_PGPASSWORD")),
	}
	postgresUserParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("PostgresUser"),
		ParameterValue: aws.String(os.Getenv("POSTGRES_USER")),
	}
	postgresPasswordParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("PostgresPassword"),
		ParameterValue: aws.String(os.Getenv("POSTGRES_PASSWORD")),
	}
	logglyTokenParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("DeLogglyToken"),
		ParameterValue: aws.String(os.Getenv("DE_LOGGLY_TOKEN")),
	}

	state.StackParametersPasswords = []*cloudformation.Parameter{
		&migrationsPgPasswordParameter,
		&postgresUserParameter,
		&postgresPasswordParameter,
		&logglyTokenParameter,
	}

	state.StackParametersAll = append(state.StackParametersPasswords, state.StackParametersStackState...)

	return nil
}
