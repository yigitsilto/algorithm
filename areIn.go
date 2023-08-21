package kata

import (
	"strings"
  "sort"
)

func InArray(array1 []string, array2 []string) []string {
    // your code
  
      result := make([]string, 0)
  
      for _, first := range array1 {
        found := false
        
        for _, second := range array2{
          if(strings.Contains(second, first)){
            found = true
          }
        }
        
        if(found){
            
          if(!contains(result, string(first))){
            result = append(result, string(first))
          }
          
           
          
        }
        
      }
  
     // sort the string
   	  sort.Strings(result)

  
      return result
  
  
}

func contains(array []string, keyword string) bool {
	for _, value := range array {
		if value == keyword {
			return true
		}
	}
	return false
}