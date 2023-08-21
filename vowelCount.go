package kata

func GetCount(str string) (count int) {
  // Enter solution here
  
    vowels := []string{"a", "e", "i", "o", "u"}
    for _, char := range(str){
      for _, vowel := range vowels{
        if(string(char) == string(vowel)) {
          count++
        }
      }
    }
  
  return count
}