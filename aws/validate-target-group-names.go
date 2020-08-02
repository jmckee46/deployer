package awsfuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jmckee46/deployer/flaw"
)

// ValidateTargetGroupNames
func ValidateTargetGroupNames(state *state) flaw.Flaw {
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

	var result map[string]interface{}

	json.Unmarshal(templateBytes, &result)

	var resources map[string]interface{}
	var properties map[string]interface{}
	var name map[string]interface{}
	resources = result["Resources"].(map[string]interface{})

	var targetGroupNames []string
	for _, value := range resources {
		if value != nil {
			typedValue := value.(map[string]interface{})

			// if a resource has a type of "AWS::ElasticLoadBalancingV2::TargetGroup" then look in it's
			// properties map for it's name and add to targetGroupNames
			if typedValue["Type"] == "AWS::ElasticLoadBalancingV2::TargetGroup" {
				properties = typedValue["Properties"].(map[string]interface{})
				name = properties["Name"].(map[string]interface{})

				// go is convinced the value of properties["Name"] is a map not a string, so I am forced into
				// the following
				var nameValue string
				for _, v := range name {
					nameValue = v.(string)
				}

				targetGroupNames = append(targetGroupNames, nameValue)
			}
		}
	}

	var cleanNames []string
	stackNameLength := len(os.Getenv("DE_STACK_NAME"))
	stackName := os.Getenv("DE_STACK_NAME")

	// remove template substitution characters
	for _, targetGroupName := range targetGroupNames {
		splitName := strings.Split(targetGroupName, "}")
		if len(splitName) > 1 {
			cleanNames = append(cleanNames, splitName[1])
		}
	}

	// check name length
	for _, cleanName := range cleanNames {
		if len(cleanName)+stackNameLength > 31 {
			msg := fmt.Sprintf("target group name, %s, is longer than 31 characters", stackName+cleanName)
			return flaw.New(msg)
		}
	}
	return nil
}
