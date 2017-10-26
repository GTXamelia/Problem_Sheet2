package main

import (
	"html/template"
	"net/http"
	"math/rand"
	"time"
	"strconv"
)
	
type myMsg struct {
	Output string
	YourGuess int
}

func guessHandler(w http.ResponseWriter, r *http.Request) {

	output :="initialise"
	
	rand.Seed(time.Now().UTC().UnixNano())

	randNum:=rand.Intn(20-1)

	var cookie, err = r.Cookie("randNum")

	if err == nil{
		randNum, _ = strconv.Atoi(cookie.Value)
	}

	yourGuess,_ := strconv.Atoi(r.FormValue("guess"))

	if yourGuess== randNum{
		output ="You guessed correctly!"
	}else if yourGuess == 0{
		output ="Input your guess from 1 - 20"
	}else if yourGuess > 20{
		output ="Input too high, please guess from 1 - 20"
	}else if yourGuess < 0{
		output ="Input too low, please guess from 1 - 20"
	}else if yourGuess < randNum{
		output="Incorrect Try Again. (Guessed too low)"
	}else {
		output="Incorrect Try Again. (Guessed too high)"
	}

	cookie = &http.Cookie{ Name: "randNum",Value: strconv.Itoa(randNum), Expires: time.Now().Add(72 * time.Hour)}
	
	http.SetCookie(w,cookie)
	
	t, _ := template.ParseFiles("guess.tmpl")
	
	t.Execute(w, &myMsg{Output:output})
}

func main() {
	
	http.HandleFunc("/", guessHandler)

	http.ListenAndServe(":8080", nil)
}