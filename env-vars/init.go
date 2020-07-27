package envvars

import (
	"os"

	"github.com/jmckee46/deployer/halt"
)

var AwsAccessKeyID string
var AwsContainerCredentialsRelativeURI string
var AwsEndpoint string
var AwsRegion string
var AwsSecretAccessKey string
var GitBranch string
var GitSha string
var JwtSigningKey string
var DePgPassword string
var DePgUser string
var LogColorization string
var LogDebugMessages string
var LogSerialization string
var MigrationsPath string
var PGDatabase string
var PGHost string
var PGPassword string
var PGPort string
var PGUser string
var PostgresPassword string
var PostgresUser string
var StackName string

var Mocked bool
var NotMocked bool

func init() {

	envVarNames := []string{
		// "AWS_ACCESS_KEY_ID",
		// "AWS_CONTAINER_CREDENTIALS_RELATIVE_URI",
		// "AWS_DEFAULT_REGION",
		// "AWS_ENDPOINT",
		// "AWS_REGION",
		// "AWS_SECRET_ACCESS_KEY",
		// "DE_CI",
		// "DE_LOCAL",
		"DE_GIT_SHA",
		// "DE_JWT_SIGNING_KEY",
		"DE_LOG_COLORIZATION",
		// "DE_LOG_DEBUG_MESSAGES",
		"DE_LOG_SERIALIZATION",
		// "DE_MIGRATIONS_PATH",
		// "DE_PGPASSWORD",
		// "DE_PGUSER",
		"DE_STACK_NAME",
		// "PGDATABASE",
		// "PGHOST",
		// "PGPASSWORD",
		// "PGPORT",
		// "PGUSER",
		"POSTGRES_PASSWORD",
		"POSTGRES_USER",
	}

	for _, envVarName := range envVarNames {
		if os.Getenv(envVarName) == "" {
			halt.PanicWith(envVarName + " must be set")
		}
	}

	AwsAccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	AwsContainerCredentialsRelativeURI = os.Getenv("AWS_CONTAINER_CREDENTIALS_RELATIVE_URI")
	AwsEndpoint = os.Getenv("AWS_ENDPOINT")
	AwsRegion = os.Getenv("AWS_REGION")
	AwsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	GitSha = os.Getenv("DE_GIT_SHA")
	LogColorization = os.Getenv("DE_LOG_COLORIZATION")
	LogDebugMessages = os.Getenv("DE_LOG_DEBUG_MESSAGES")
	LogSerialization = os.Getenv("DE_LOG_SERIALIZATION")
	MigrationsPath = os.Getenv("DE_MIGRATIONS_PATH")
	PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	PostgresUser = os.Getenv("POSTGRES_USER")
	StackName = os.Getenv("DE_STACK_NAME")

	if os.Getenv("DE_LOCAL") == "true" || os.Getenv("DE_CI") == "true" {
		Mocked = true
	} else {
		NotMocked = true
	}
}
