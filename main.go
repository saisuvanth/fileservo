package main

import "github.com/saisuvanth/fileservo/app"

func main() {
	err := app.SetApp()
	if err != nil {
		panic(err)
	}
}
