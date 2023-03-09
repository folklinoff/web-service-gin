package main

func main() {
	var publishers []Publisher = []Publisher{
		Publisher{name: "Valve", year: 1996},
		Publisher{name: "Mojang", year: 2009},
		Publisher{name: "Bethesda Game Studios", year: 2001},
	}
	var games []Game = []Game{
		Game{name: "CS:GO", genre: "FPS", year: 2001, publisher: &publishers[0]},
		Game{name: "Minecraft", genre: "Sandbox", year: 2009, publisher: &publishers[1]},
		Game{name: "Skyrim", genre: "Roleplay Fighting", year: 2011, publisher: &publishers[2]},
	}
}
