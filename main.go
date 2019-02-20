// PSNapi project main.go
package main

import (
	"fmt"

	"PSNapi/handlers"
	//"net/http"
)

func main() {
	oauth, err := handlers.Login("0d0fca9f-4bc1-4bd9-b488-2876984a6ad3")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(oauth)
	profile, err := handlers.UserInfo(oauth, "jiegokoji")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(profile)
	/*games, err := handlers.UserGames(oauth, "jiegokoji")
	for _, title := range games.Titles {
		fmt.Println(title.TitleId)
		fmt.Println(title.Name)
		fmt.Println(title.Image)
	}*/
	//trophyTitles, err := handlers.GetGameTrophieTitles(oauth, "CUSA12448_00")
	//fmt.Println(trophyTitles.Apps[0].TrophyTitles[0].NpCommunicationId)
	alltrophies, err := handlers.GetGameTrophies(oauth, "NPWR16409_00", "jiegokoji")
	for _, trophie := range alltrophies.Trophies {
		if trophie.ComparedUser.OnlineId != "" && trophie.ComparedUser.Earned == true {
			fmt.Println(trophie.TrophyName)
		}
	}
	//fmt.Println(alltrophies)

}