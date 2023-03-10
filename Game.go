package main

type Game struct {
	Name      string     `json:"name"`
	Genre     string     `json:"genre"`
	Year      int16      `json:"year"`
	Publisher *Publisher `json:"Publisher"`
}
