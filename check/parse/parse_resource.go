package parse

import "fmt"

func GetResourseList(role_policy map[string]interface{}) ([]interface{}, error) {

	policyMandatory := [2]string{"PolicyName", "PolicyDocument"}

	polDocMandatory := [2]string{"Version", "Statement"}

	statementMandatory := [3]string{"Effect", "Action", "Resource"}

	resourceList := []interface{}{}

	_, err := checkForMandatory(role_policy, policyMandatory[:])
	if err != nil {
		return resourceList, err
	}

	polDoc, _ := role_policy["PolicyDocument"].(map[string]interface{})

	_, err = checkForMandatory(polDoc, polDocMandatory[:])
	if err != nil {
		return resourceList, err
	}

	stats, _ := polDoc["Statement"].([]interface{})
	for _, item := range stats {
		to_check, _ := item.(map[string]interface{})

		_, err = checkForMandatory(to_check, statementMandatory[:])
		if err != nil {
			return resourceList, err
		}
		resourceList = append(resourceList, to_check["Resource"])
	}

	return resourceList, nil

}

func checkForMandatory(content map[string]interface{}, needed []string) (bool, error) {

	for _, keyword := range needed {
		_, ok := content[keyword]
		if !ok {
			return false, fmt.Errorf("%s not in json", keyword)
		}
	}
	return true, nil
}
