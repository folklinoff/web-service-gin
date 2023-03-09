package main

type Game struct {
	name      string     `json:"name"`
	genre     string     `json:"genre"`
	year      int16      `json:"year"`
	publisher *Publisher `json:"Publisher"`
}

type Publisher struct {
	name string `json:"name"`
	year int16  `json:"year"`
}

func main() {

}
