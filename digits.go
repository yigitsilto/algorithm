package kata
import (
	"strconv"
)

func DigitalRoot(n int) int {

  stringValueOfN := strconv.Itoa(n)
   var resultCount int
   count := 0
   for _, char := range stringValueOfN{
    		digit, _ := strconv.Atoi(string(char))
        count += int(digit)
        resultCount = int(digit) + int(resultCount)
   }
  
  if(int(resultCount) >= 10){
     return DigitalRoot(resultCount)
  }
  
  
  
  
  return resultCount
    
  
  
  
}
