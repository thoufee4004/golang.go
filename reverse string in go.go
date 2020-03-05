package main

import (
"fmt"
)

func main()  {
	var str ,str1 string
	str1 = ""
	fmt.Scanf("%s",&str)
	fmt.Println(str)
	for _, r:= range str{
		/*defer fmt.Printf(string(r))*/  //if u use this cmd next line trick and cmd str1
		str1=string(r)+str1  // if u use this cmd above line
	}
	fmt.Println(str1)
}
