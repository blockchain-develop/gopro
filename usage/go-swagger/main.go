package main

import "net/http"

func main()  {
	fs := http.FileServer(http.Dir("swagger"))
	http.Handle("/swagger/", http.StripPrefix("/swagger/", fs))
	http.ListenAndServe(":8000", nil)
}
