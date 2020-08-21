package envvars

import (
	"fmt"
	"os"

	"github.com/jmckee46/deployer/flaw"
)

func ValidateDeploymentEnvVars() flaw.Flaw {
	fmt.Println("validating environment variables...")

	allEnvVars := []string{
		"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
		"AWS_REGION",
		"AWS_DEFAULT_REGION",
		"ARTIFACTS_BUCKET",
		"ARTIFACTS_KEY",
		"ARTIFACTS_PATH",
		"ARTIFACTS_REGION",
		"ARTIFACTS_SECRET",
		"DE_ARTIFACTS_PATH",
		"DE_CI",
		"DE_DOMAIN",
		"DE_EMAIL",
		"DE_GIT_BRANCH",
		"DE_GIT_BRANCH_HASH",
		"DE_GIT_SHA",
		"DE_LOAD_BALANCER_HOSTNAME",
		// "DE_LOGGLY_TOKEN",
		"DE_LOG_COLORIZATION",
		"DE_LOG_SERIALIZATION",
		// "DE_MIGRATIONS_PGPASSWORD",
		"DE_ROOT_BUCKET",
		"DE_STACK_BUCKET",
		"DE_STACK_NAME",
		// "DE_SUBNET_CIDR_BLOCKS",
		// "DE_VPC_CIDR_BASE",
		"HOSTED_ZONE_ID",
		"POSTGRES_PASSWORD",
		"POSTGRES_USER",
		"PYTHONWARNINGS",
	}

	for _, value := range allEnvVars {
		if os.Getenv(value) == "" {
			fmt.Println("envvar:", value)
			return flaw.New(value + "is not set properly for deployment")
		}
	}

	return nil
}
