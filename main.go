package main
import (
  "fmt"
)
type Product interface{
Income()float64
 IncomeMonth()float64
  
}
type Deposit struct{
  Amount float64
  Term int
  Annual float64
}
type SavingAccount struct{
  DailyBalance bool
  Amount float64
  Annual float64
}
type DigitalAssets struct{
  Amount float64
  Term int
  Annual float64
  Tax float64
}

func(s SavingAccount)Income()float64{
    return s.Amount*((s.Annual/100)/12)
}


func(d Deposit)Income()float64{
  d.Amount*= (d.Annual/100)/12
  d.Amount*=float64(d.Term)
  return d.Amount
}


func(d DigitalAssets)Income()float64{
  d.Amount*= (d.Annual/100)/12
  d.Amount*=float64(d.Term)
  d.Amount-= d.Amount * (d.Tax/100)
  return d.Amount
}

func(s SavingAccount)IncomeMonth()float64{
  return s.Amount*((s.Annual/100)/12)
}


func(d Deposit)IncomeMonth()float64{
  return d.Income()/float64(d.Term)
}


func(d DigitalAssets)IncomeMonth()float64{
return d.Income()/float64(d.Term)
}


func Description (a Product){
  fmt.Println(a)
  fmt.Println("Доход за весь период:", a.Income())
fmt.Println("Доодность за месяц:", a.IncomeMonth())
}

func main(){
var myProducts []Product
  var choice int
  var amount float64
fmt.Printf("Что хотите сделать?\n 1 - получить новый продукт\n")
  fmt.Scanln(&choice)
  switch choice{
  case 1:
    fmt.Printf("Цифровой актив - 1\n Вклад - 2\n Накопительный счет - 3\n")
    // Можно добавить еще вариант с кредитами но не буду усложнять
    fmt.Scanln(&choice)
    fmt.Println("Введите сумму")
    fmt.Scanln(&amount)
    
    myProducts = append(myProducts, DigitalAssets{
        Amount: amount,
    Term: choice,
    Annual: 15.2,
    Tax: 13,
      })

        asset := myProducts[len(myProducts)-1]

    
    switch choice{  
    case 1:
    
      fmt.Println("На какой срок? Доступно 1, 3, 6, 9 или 12 месяцев")
      fmt.Scanln(&choice)
      switch choice{
      case 1:
        

      case 3:
        myProducts = append(myProducts, DigitalAssets{
        Amount: amount,
    Term: choice,
    Annual: 15.7,
    Tax: 13,
      })
  
      case 6:
       myProducts = append(myProducts, DigitalAssets{
        Amount: amount,
    Term: choice,
    Annual: 17,
    Tax: 13,
      })
  
      case 9:
        myProducts = append(myProducts, DigitalAssets{
        Amount: amount,
    Term: choice,
    Annual: 17.5,
    Tax: 13,
      })

        
      case 12:
      myProducts = append(myProducts, DigitalAssets{
        Amount: amount,
    Term: choice,
    Annual: 18,
    Tax: 13,
      })

        }

  Description(asset)





      
    case 2:
          fmt.Println("На какой срок? Доступно  от 1 до 12, или 24 месяца")
      fmt.Scanln(&choice)
      myProducts = append(myProducts, Deposit{
        
      })

    case 3:
      myProducts = append(myProducts, SavingAccount{
        
      })

    }
  }
}