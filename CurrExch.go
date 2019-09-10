package exchange

import(
        "fmt" 
        "math"
      )
func exchange() 
{

  //declarations
  //stores options

  var options, amount, euro, dollar, euro2, gbd2, dtotal, ptotal int
  var ans, choice int

  //core program

  //welcome

  fmt.Println("Welcome to the currency exchange program")
  fmt.Println("Enter 1 if you want to convert euro to dollar")
  fmt.Println("Enter 2 if you want to convert euro to pounds")

  fmt.Println("Enter your choice:",&choice)

  //scanf("%d", &choice)
  get char()
  
  fmt.Println("Enter the amount of euro:",&amount)
  //scanf("%f", &amount)
  
  if (choice ==1)
    {
      get char()
  
      fmt.Println("%.1f euro(s)is equivalent to %.2f dollar(s)", amount, amount*0.693)
    }
  getchar()
  
  for(amount=1; amount<=10;amount++)
    {
      fmt.Println("%.1f euro(s) is equivalent to %.2f gbp and %.2f usd\n", amount, amount*1.293")
  
      fmt.Println("Goodbye\n")
    }
  }
  
  
