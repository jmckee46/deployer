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

	rootBucketParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("DeRootBucket"),
		ParameterValue: aws.String(os.Getenv("DE_ROOT_BUCKET")),
	}
	gitShaParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("DeGitSha"),
		ParameterValue: aws.String(os.Getenv("DE_GIT_SHA")),
	}
	loadBalancerHostnameParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("DeLoadBalancerHostname"),
		ParameterValue: aws.String(os.Getenv("DE_LOAD_BALANCER_HOSTNAME")),
	}
	// logglyTokenParameter := cloudformation.Parameter{
	// 	ParameterKey:   aws.String("DeLogglyToken"),
	// 	ParameterValue: aws.String(os.Getenv("DE_LOGGLY_TOKEN")),
	// }
	stackBucketParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("DeStackBucket"),
		ParameterValue: aws.String(os.Getenv("DE_STACK_BUCKET")),
	}
	stackNameParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("DeStackName"),
		ParameterValue: aws.String(os.Getenv("DE_STACK_NAME")),
	}
	subnetCidrBlockParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("DeSubnetCidrBlocks"),
		ParameterValue: aws.String(os.Getenv("DE_SUBNET_CIDR_BLOCKS")),
	}
	tlsCertificationArnParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("AnTlsCertificateArn"),
		ParameterValue: aws.String(arn),
	}
	VpcCidrBaseParameter := cloudformation.Parameter{
		ParameterKey:   aws.String("AnVpcCidrBase"),
		ParameterValue: aws.String(os.Getenv("DE_VPC_CIDR_BASE")),
	}

	state.StackParametersStackState = []*cloudformation.Parameter{
		&rootBucketParameter,
		&gitShaParameter,
		&loadBalancerHostnameParameter,
		&stackBucketParameter,
		&stackNameParameter,
		&subnetCidrBlockParameter,
		&tlsCertificationArnParameter,
		&VpcCidrBaseParameter,
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

	state.StackParametersPasswords = []*cloudformation.Parameter{
		&migrationsPgPasswordParameter,
		&postgresUserParameter,
		&postgresPasswordParameter,
	}

	state.StackParametersAll = append(state.StackParametersPasswords, state.StackParametersStackState...)

	return nil
}
