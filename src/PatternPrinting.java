package test;
import java.util.*;
public class PatternPrinting {

	public static void main(String[] args) {
		
		for(int i=1;i<=5;i++){
			for(int j=1;j<=5;j++){
				if(i==1 || j==1 || i==5 || j==5){
					System.out.println("*");
				}
				else{
					System.out.println(" ");
				}
				
			}
			System.out.println();
		}
		

	}

}
