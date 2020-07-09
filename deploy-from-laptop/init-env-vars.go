package main

import (
	"fmt"
	"os"
	"strings"

	"crypto/sha256"

	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/git"
	"github.com/jmckee46/deployer/logger"
)

// initEnvVars sets laptop env-vars to match travis-ci env-vars in order to
// deploy from laptop
func initEnvVars() flaw.Flaw {
	fmt.Println("initializing local env-vars like travis-ci...")
	_, present := os.LookupEnv("AWS_ACCESS_KEY_ID")
	if !present {
		return flaw.New("AWS_ACCESS_KEY_ID must be set")
	}
	_, present = os.LookupEnv("AWS_SECRET_ACCESS_KEY")
	if !present {
		return flaw.New("AWS_SECRET_ACCESS_KEY must be set")
	}
	err := os.Setenv("AWS_REGION", "us-west-2")
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("AWS_DEFAULT_REGION", os.Getenv("AWS_REGION"))
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("ARTIFACTS_BUCKET", "deployer.global")
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("ARTIFACTS_KEY", os.Getenv("AWS_ACCESS_KEY_ID"))
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("ARTIFACTS_PATH", "artifacts")
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("ARTIFACTS_REGION", os.Getenv("AWS_REGION"))
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("ARTIFACTS_SECRET", os.Getenv("AWS_SECRET_ACCESS_KEY"))
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("DE_ARTIFACTS_PATH", os.Getenv("ARTIFACTS_PATH"))
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("DE_CI", "true")
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("DE_DOMAIN", "myAppTest.com")
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("DE_LOGGLY_TOKEN", "")
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("DE_LOG_COLORIZATION", "false")
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("DE_LOG_SERIALIZATION", "json-compact")
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("DE_MIGRATIONS_PGPASSWORD", os.Getenv(""))
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("DE_GIT_SHA", git.CurrentSha())
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("DE_GLOBAL_BUCKET", os.Getenv("ARTIFACTS_BUCKET"))
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("DE_VPC_CIDR_BASE", "192.168.0")
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("DE_SUBNET_CIDR_BLOCKS", "192.168.0.0/24,192.168.1.0/24,192.168.2.0/24")
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("HOSTED_ZONE_ID", "/hostedzone/xxxxxxxxxxxxxx")
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("PYTHONWARNINGS", "ignore")
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("DE_GIT_BRANCH", gitBranch())
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("DE_GIT_BRANCH_HASH", branchHash())
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("DE_STACK_NAME", stackName())
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("DE_LOAD_BALANCER_HOSTNAME", loadBalancerHostname())
	if err != nil {
		return flaw.From(err)
	}
	err = os.Setenv("DE_STACK_BUCKET", stackBucket())
	if err != nil {
		return flaw.From(err)
	}

	return nil
}

func gitBranch() string {
	branch, err := git.CurrentBranch()
	if err != nil {
		logger.Panic("init-env-vars while deploying from laptop", err)
	}

	return strings.ToLower(branch)
}

func branchHash() string {
	data := []byte(os.Getenv("DE_GIT_BRANCH"))
	sum := sha256.Sum256([]byte(data))

	return string(sum[:])
}

func stackName() string {
	deGitBranch := os.Getenv("DE_GIT_BRANCH")
	deGitBranchHash := os.Getenv("DE_GIT_BRANCH_HASH")

	return deGitBranch[0:8] + "-" + deGitBranchHash[0:4]
}

func loadBalancerHostname() string {
	stackName := os.Getenv("DE_STACK_NAME")
	domain := os.Getenv("DE_DOMAIN")

	return stackName + "." + domain
}

func stackBucket() string {
	stackName := os.Getenv("DE_STACK_NAME")

	return "deployer.stack." + stackName
}

/*
- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY:
- AWS_REGION=us-west-2
- AWS_DEFAULT_REGION=$AWS_REGION
# the following ARTIFACTS_* environment variables are out of sort order due to dependencies on AWS_* environment variables.
- ARTIFACTS_BUCKET=deployer.global
- ARTIFACTS_KEY=$AWS_ACCESS_KEY_ID
- ARTIFACTS_PATH=artifacts
- ARTIFACTS_REGION=$AWS_REGION
- ARTIFACTS_SECRET=$AWS_SECRET_ACCESS_KEY
- DE_ARTIFACTS_PATH=$ARTIFACTS_PATH
- DE_CI=true
- DE_DOMAIN=myAppTest.com
- DE_LOGGLY_TOKEN=
- DE_LOG_COLORIZATION=false
- DE_LOG_SERIALIZATION=json-compact
- DE_MIGRATIONS_PGPASSWORD=
- DE_GIT_SHA=$TRAVIS_COMMIT
- DE_GLOBAL_BUCKET=$ARTIFACTS_BUCKET
- DE_VPC_CIDR_BASE=192.168.0
- DE_SUBNET_CIDR_BLOCKS=192.168.0.0/24,192.168.1.0/24,192.168.2.0/24
- HOSTED_ZONE_ID=/hostedzone/Z2ZUXUMTTHLE9Y
- PYTHONWARNINGS="ignore"
# the remaining env-vars are order dependent
- DE_GIT_BRANCH=$(echo $TRAVIS_BRANCH | tr '[:upper:]' '[:lower:]')
- DE_GIT_BRANCH_HASH=$(echo -n $DE_GIT_BRANCH | shasum -a 256)
- DE_STACK_NAME=${DE_GIT_BRANCH:0:8}-${DE_GIT_BRANCH_HASH:0:4}
- DE_LOAD_BALANCER_HOSTNAME=$DE_STACK_NAME.$DE_DOMAIN
- DE_STACK_BUCKET=deployer.stack.$DE_STACK_NAME
*/
