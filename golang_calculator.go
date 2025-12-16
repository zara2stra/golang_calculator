package main

import "fmt"
import "time"
import "strconv"
import "flag"

func main() {
	var num1 float64
	var num2 float64
	var ans string
	var operator string
	var result float64
	var input1 string
	var input2 string
seconds := flag.Int("s", 3, "number of seconds to wait")
flag.Parse()
for {
    //Here we input a string and then convert to float to avoid invalid inputs
    for {
        fmt.Println("Enter number 1")
	    _, err := fmt.Scanln(&input1)
	    if err != nil {
	        fmt.Println("Invalid number. Please enter a valid number")
	        continue
	    }
	    num1, err = strconv.ParseFloat(input1, 64)
        if err != nil {
        fmt.Println("Invalid number. Please enter a valid number.")
        continue
        }
	break
        }
    for {
        fmt.Println("Enter number 2")
	    _, err := fmt.Scanln(&input2)
	    if err != nil {
	        fmt.Println("Invalid number. Please enter a valid number")
	        continue
	    }
	    num2, err = strconv.ParseFloat(input2, 64)
        if err != nil {
        fmt.Println("Invalid number. Please enter a valid number.")
        continue
        }
	break
    }
for {
	fmt.Println("\nEnter the operator (+, - , *, /)")
	fmt.Scanf("%s", &operator)
/*	if operator == "+" {
	result = num1 + num2
    } else if operator == "-" {
        result = num1 - num2
        } else if operator == "*" {
            result = num1 * num2
        } else if operator == "/" {
            result = num1 / num2
        } else {
            fmt.Println("Invalid operator")
            continue
        }
*/
    switch {
        case operator == "+":
        result = num1 + num2
        case operator == "-":
        result = num1 - num2
        case operator == "*":
        result = num1 * num2
        case operator == "/":
        result = num1 / num2
        default:
        fmt.Println("Invalid Operator")
        continue
    }

	fmt.Print("\nThinking")
	for i := 0; i < *seconds; i++ {
		time.Sleep(1 * time.Second)
		fmt.Print(".")
	}
    fmt.Println("\n\nHere is the result")
	fmt.Printf("%.2f %s %.2f = %.3f", num1, operator, num2, result)
	break
}
for {
	fmt.Println("\n\nDo you want to continue? (y/n)")
    fmt.Scanf("%s", &ans)
if ans == "y" || ans == "Y" {
         break
    } else { 
         if ans == "n" || ans == "N"{
             fmt.Println("Bye")
             return
                } else {
                        fmt.Println("\nInvalid answer")
                        }
            }
    }
}
}
