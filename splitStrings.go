package kata

import (
  "fmt"
 )

func Solution(str string) []string {
  
    c := []rune(str)
    cCount:= len(c)
   
    var stringsArray []string = make([]string, 0)
    
  
  if(cCount <= 1){
                stringsArray = append(stringsArray, string(c[0]) + "_")

  }
   
    
   for index := 0; index < len(c)-1; index++ {
          
          firstTwoValue := string(c[index]) + string(c[index+1])
          stringsArray = append(stringsArray, firstTwoValue)
          index++
          
      if(cCount % 2 != 0 && index + 1 == cCount - 1 ){
        fmt.Println(index == cCount - 1)
          lastValue := string(c[cCount - 1]) + "_"
          stringsArray = append(stringsArray, string(lastValue))

      }
    }
   
  
  return stringsArray

}