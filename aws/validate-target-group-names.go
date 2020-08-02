package awsfuncs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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

	fmt.Println("StackName:", os.Getenv("DE_STACK_NAME"))
	fmt.Println("targetGroupNames:", targetGroupNames)
	stackNameLength := len(os.Getenv("DE_STACK_NAME"))
	for targetGroupName := range targetGroupNames {
		NEED TO SPLIT OUT THE NAME AND CHECK LENGTH
		if len(stackNameLength + )
	}

	return nil
}
