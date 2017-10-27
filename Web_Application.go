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
	UserGuess int
}

func guessHandler(w http.ResponseWriter, r *http.Request) {

	output :="initialise"
	
	rand.Seed(time.Now().UTC().UnixNano())
	randNum:=rand.Intn(20-1)

	guessTemplate, _ := template.ParseFiles("guess.tmpl")

	var cookie, err = r.Cookie("randNum")

	userGuess,_ := strconv.Atoi(r.FormValue("guess"))

	if err == nil{
		randNum, _ = strconv.Atoi(cookie.Value)
	}

	cookie = &http.Cookie{ Name: "randNum",Value: strconv.Itoa(randNum), Expires: time.Now().Add(72 * time.Hour)}
	http.SetCookie(w,cookie)

	if userGuess== randNum{
		output ="You guessed correctly!"
	}else if userGuess == 0{
		output ="Input your guess from 1 - 20"
	}else if userGuess > 20{
		output ="Input too high, please guess from 1 - 20"
	}else if userGuess < 0{
		output ="Input too low, please guess from 1 - 20"
	}else if userGuess < randNum{
		output="Incorrect Try Again. (Guessed too low)"
	}else {
		output="Incorrect Try Again. (Guessed too high)"
	}
	
	guessTemplate.Execute(w, &myMsg{Output:output,UserGuess:userGuess})
}

func main() {
	
	http.HandleFunc("/", guessHandler)

	http.ListenAndServe(":8080", nil)
}