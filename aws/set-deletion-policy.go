package awsfuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jmckee46/deployer/flaw"
)

// SetDeletionPolicy cannot be done in a cloudformation template at this time. "Currently, AWS
// CloudFormation supports the Fn::If intrinsic function in the metadata attribute, update policy
// attribute, and property values in the Resources section and Outputs sections of a template."
// https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-conditions.html
func SetDeletionPolicy(state *State) flaw.Flaw {
	fmt.Println("  setting database deletion policy...")

	// open template
	template, err := os.Open(state.RenderedTemplateLocal)
	if err != nil {
		return flaw.From(err)
	}
	defer template.Close()

	// read in template
	templateBytes, err := ioutil.ReadAll(template)
	if err != nil {
		return flaw.From(err)
	}

	// unmarshal the file
	var result map[string]interface{}
	json.Unmarshal(templateBytes, &result)

	var resources map[string]interface{}

	resources = result["Resources"].(map[string]interface{})

	// if a resource has a type of "AWS::RDS::DBInstance" then set a
	// deletion policy based on stack name
	for _, value := range resources {
		typedValue := value.(map[string]interface{})

		if typedValue["Type"] == "AWS::RDS::DBInstance" {
			if os.Getenv("DE_STACK_NAME") == "master" {
				typedValue["DeletionPolicy"] = "Snapshot"
			} else {
				typedValue["DeletionPolicy"] = "Delete"
			}
		}
	}

	updatedResult, err := json.Marshal(result)
	if err != nil {
		return flaw.From(err)
	}

	err = ioutil.WriteFile(state.RenderedTemplateLocal, updatedResult, 0755)
	if err != nil {
		return flaw.From(err)
	}

	return nil
}
