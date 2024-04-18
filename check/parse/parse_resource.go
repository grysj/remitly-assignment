package parse

import "fmt"

func GetResourseList(role_policy map[string]interface{}) (map[string]bool, error) {

	policyMandatory := [2]string{"PolicyName", "PolicyDocument"}

	polDocMandatory := [2]string{"Version", "Statement"}

	statementMandatory := [3]string{"Effect", "Action", "Resource"}

	resourceList := map[string]bool{}

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

		to_add := to_check["Resource"]
		switch to_add := to_add.(type) {
		case string:
			resourceList[to_add] = true
		case []interface{}:
			for _, toAddFromT := range to_add {
				switch v := toAddFromT.(type) {
				case string:
					resourceList[v] = true
				default:
					return resourceList, fmt.Errorf(`incompatible type in "Resource"`)
				}

			}
		default:
			return resourceList, fmt.Errorf(`incompatible format of "Resource"`)
		}

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
