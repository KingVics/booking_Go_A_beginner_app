package main

import( 
	"fmt" 
	"booking/helper"
	"time"
	)


var conferenceName = "GoCon"
const ticektPrice = 50
var remainTicket uint = 50
var booking = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numofTicket uint
}

func main() {

	// fmt.Printf("Conference name is %T, ticket price is %T, remaining ticket is %T\n", conferenceName, ticektPrice, remainTicket)
	// fmt.Println("The conference name is:", conferenceName)
	// fmt.Println("The ticket price is:", ticektPrice)
	// fmt.Printf("Welcome to %v where all tickets are bought at the same time. We have %v and %v tickets", conferenceName, ticektPrice, remainTicket )
	greetUser()

	for {
	
		if remainTicket != 0  {
			
			userName, lastName, email, userTicker := getUsersInput()

			isValidNames, isEmail, isValidTicket := helper.ValidateUserInput(userName,lastName, email, userTicker, remainTicket )
			
			if isValidNames && isEmail && isValidTicket {
		
				//Booke
				bookTicket(userName, lastName, userTicker, email)
				go sendTickets(userTicker, userName, lastName, email)

				// Print first name
				firstNames := getFirstname()
				fmt.Printf("%v booking: \n", firstNames)

				fmt.Printf(" Thank you %v %v  for purchasing %v tickets for %v. You will receiev an email confirmation at %v\n", userName, lastName, userTicker, conferenceName, email)
				fmt.Printf("%v tickets remaining \n", remainTicket)

			} else {
				if !isValidNames {
					fmt.Println("Please enter a valid names")
				}
				if !isEmail {
					fmt.Println("Please enter a valid email")
				}
				if !isValidTicket {
					fmt.Println("Please enter a valid ticket number")
				}
				
			}
		} else {
			fmt.Println("No more ticker")
			break
		}

	}
	// switch city {
	// 	case "Nigeria":
	// 		// do something
	// 	case "Ghana":
	// 		// do something
	// 	case "Ivory coast": 
	// 		// do something
	// 	case "Cameroon":
	// 		//do something
	// 	default: 
	// 		// do soemthing
	// }

}

func greetUser() {
	fmt.Printf("Welcome to %v where all tickets are bought at the same time. We have %v and %v tickets\n", conferenceName, ticektPrice, remainTicket )
}

func getFirstname() []string {
	firstNames := []string{}
	for _, booked := range booking {
		firstNames = append(firstNames, booked.firstName)
	}

	return firstNames
}



func getUsersInput() (string, string, string, uint) {
	var userName string
	var lastName string
	var email string
	var userTicker uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&userName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter your ticker number: ")
	fmt.Scan(&userTicker)

	return userName, lastName, email, userTicker
}

func bookTicket(userName string, lastName string, userTicker uint, email string ) {
	var userData  =  UserData {
		firstName: userName,
		lastName: lastName,
		email: email,
		numofTicket: userTicker,
	}

	fmt.Printf("%v\n", userData)

	booking = append(booking, userData)
	remainTicket = remainTicket - userTicker
}


func sendTickets(userTicker uint, userName string, lastName string, email string) {
	time.Sleep( 10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicker, userName, lastName)
	fmt.Println("****************")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("****************")
}