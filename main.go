package main

// struct je kao objekat u JS-u (key + value parovi)
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// * zvezdica je pointer

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
