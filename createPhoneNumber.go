package kata

import (
  "strconv"
)

func CreatePhoneNumber(numbers [10]uint) string {

   phoneString := "("
   for index, number := range numbers {
     
     if(index <= 2){
       phoneString += strconv.FormatUint(uint64(number), 10)
     }
     if(index == 2){
       phoneString += ") "
     }
     
     if(index >= 3 && index <= 5){
        phoneString += strconv.FormatUint(uint64(number), 10)
     }
     if(index == 5){
              phoneString += "-"
     }
     if(index > 5){
               phoneString += strconv.FormatUint(uint64(number), 10)
     }
     
       
        
   }
  
  return phoneString
  
}