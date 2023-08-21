package kata
import (
    "strings"
)
func SpinWords(str string) string {
  res1 := strings.Split(str, " ")
  
  
  for index, word := range res1 {
    if(len(word) >= 5){
        res1[index] = Reverse(word)
       

    }
    
  }
  
  
  
  return strings.Join(res1, " ")
}// SpinWords

func Reverse(input string) (output string) {
	for _, r := range input {
		output = string(r) + output
	}

	return output
}