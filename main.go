package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Gamemode struct {
	Id int
	Name string
	Image string
	Players string
	Label string
}

type Contributor struct {
	Battletag string
	Avatar string
}

type Arcade struct {
	Created_at string
	Is_today bool
	User Contributor
	Modes struct {
		Tile_1 Gamemode
		Tile_2 Gamemode
		Tile_3 Gamemode
		Tile_4 Gamemode
		Tile_5 Gamemode
		Tile_6 Gamemode
		Tile_7 Gamemode
	}
}

var reqClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	req, err := reqClient.Get(url)
	req.Header.Set("User-Agent", "Pandaaa/OverwatchArcadeCLI")
	if err != nil {
		return err
	}
	defer req.Body.Close()
	return json.NewDecoder(req.Body).Decode(target)
}

func main() {

	arcade := new (Arcade)

	err := getJson("https://overwatcharcade.today/api/overwatch/today", arcade)

	if err != nil {
		fmt.Println("error while fetching arcade modes")
		return
	}

	if err != nil {
		fmt.Println("error while parsing arcade modes")
		return
	}

	if arcade.Is_today {

		time, err := time.Parse(time.RFC3339Nano, arcade.Created_at)

		if err != nil {
			fmt.Println("error while printing arcade modes")
			return
		}

		fmt.Printf("     Current Overwatch Arcade Gamemodes     \n")
		fmt.Printf("-------------------------------------------\n")
		fmt.Printf("%s [%s] \n", arcade.Modes.Tile_1.Name, arcade.Modes.Tile_1.Players)
		fmt.Printf("%s [%s] \n", arcade.Modes.Tile_2.Name, arcade.Modes.Tile_2.Players)
		fmt.Printf("%s [%s] \n", arcade.Modes.Tile_3.Name, arcade.Modes.Tile_3.Players)
		fmt.Printf("%s [%s] \n", arcade.Modes.Tile_4.Name, arcade.Modes.Tile_4.Players)
		fmt.Printf("%s [%s] \n", arcade.Modes.Tile_5.Name, arcade.Modes.Tile_5.Players)
		fmt.Printf("%s [%s] \n", arcade.Modes.Tile_6.Name, arcade.Modes.Tile_6.Players)
		fmt.Printf("%s [%s] \n", arcade.Modes.Tile_7.Name, arcade.Modes.Tile_7.Players)
		fmt.Printf("-------------------------------------------\n")
		fmt.Printf("Last Updated %s\n", time.Format("02 Jan 2006 03:04:05 PM"))
		fmt.Printf("     by %s", arcade.User.Battletag)

	} else {
		fmt.Printf("Current Overwatch Arcade Gamemodes (Last Updated %s)", arcade.Created_at)
	}

}
