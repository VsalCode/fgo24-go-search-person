package utils

import "strings" 

func SearchPerson(users []string , input string) []string {
	var result []string

	for x := range users {
		true := strings.Contains(strings.ToLower(users[x]) , strings.ToLower(input))
		if(true){
			result = append(result, users[x])
		}
	}
	
	return result 
}
