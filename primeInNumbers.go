import java.util.*;
public class PrimeDecomp {
   
    public static String factors(int n) {
      ArrayList<Integer> arr = new ArrayList<>();
        
        for (int asal = 2; asal * asal <= n; asal++) {
            while (n % asal == 0) {
                arr.add(asal);
                n /= asal;
            }
        }
        
        if (n > 1) {
            arr.add(n);
        }
      
        StringBuilder result = new StringBuilder();
        int currentFactor = arr.get(0);
        int currentCount = 1;
        for (int i = 1; i < arr.size(); i++) {
            if (arr.get(i) == currentFactor) {
                currentCount++;
            } else {
                if (currentCount > 1) {
                    result.append("(").append(currentFactor).append("**").append(currentCount).append(")");
                } else {
                    result.append("(").append(currentFactor).append(")");
                }
                currentFactor = arr.get(i);
                currentCount = 1;
            }
        }
        if (currentCount > 1) {
            result.append("(").append(currentFactor).append("**").append(currentCount).append(")");
        } else {
            result.append("(").append(currentFactor).append(")");
        }
        
        return result.toString();
      
    }
  public static boolean checkAsal(int value)
    {
      int sayac = 0;
        
        for(int i = 2; i < value; i++)
        {
            if(value % i == 0) {
                sayac++;
            }
        }
        if(sayac == 0) {
           return true;
        }
        else {
        return false;      
        }
    }
       
}