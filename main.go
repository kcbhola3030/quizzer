package main

import (
	//"bytes"
	"encoding/json"
	"fmt"
	//"net/http"
	"time"
)


type Person struct {
    name string `json:"name"`
    score  int    `json:"score"`
}


func main() {
	url := "http://localhost:3000/questions"

	var name string
	fmt.Print("Enter your name : ")
	fmt.Scanf("%s",&name);


	data := Person{name: name}
    jsonData, _ := json.Marshal(data)

	fmt.Print(data); 
	fmt.Print(jsonData); 


	// _, err := http.Post("http://localhost:3000/user", "application/json", bytes.NewBuffer(jsonData));
	// if err!=nil{
	// 	fmt.Print("Something went wrong");
	// }


	
	problems := questionPuller(url)

	tobj := time.NewTimer(20*time.Duration(len(problems)) * time.Second) // Time for all the questions --> 1 question => 10 seconds

	correctAns := 0
ProblemLoop:

	for i, problem := range problems {
		var answer string
		fmt.Printf("\nProblem %d: %s", i+1, problem.Question)
		fmt.Printf("\n a. %s \n b. %s \n c. %s \n d. %s \n Select Option 'a','b','c','d' : ",
			problem.Options.A, problem.Options.B, problem.Options.C, problem.Options.D)
		ansC := make(chan string)

		go func() {
			fmt.Scanf("%s", &answer)
			ansC <- answer
		}()
		
		
		select {
		case <-tobj.C:
			fmt.Println("\nTime Over !!! Your Quiz has been Submitted")
			break ProblemLoop
		case iAns := <-ansC:
			if iAns == problem.Answer {
				correctAns++
			}
			if i == len(problems)-1 {
				fmt.Print("All Questions Submitted Successfully...:) ")
			}

		}

	}

	fmt.Printf("Correct %d out of %d", correctAns, len(problems))

}
