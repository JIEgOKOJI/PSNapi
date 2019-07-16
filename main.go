// PSNapi project main.go
package main

import (
	"PSNapi/handlers"
	"fmt"

	//"strconv"
	//"net/http"
	"time"
)

func main() {
	oauth, err := handlers.Login("f00c5319-2325-4eb7-b0a5-fe15a09fd44d")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(oauth)
	profile, err := handlers.UserInfo(oauth, "ledokol322")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(profile)
	//	fmt.Println(profile.Profile.)*/
	//handlers.UserAddFriend(oauth, "ledokol322", "gorcheque", "test msg")
	/*usernam := "jiegokoji"
	threads, _ := handlers.MessageThreads(oauth, usernam)
	for _, title := range threads.ThreadIds {
		threadInfo, _ := handlers.MessageThreadInfo(oauth, title.ThreadId)
		for id, events := range threadInfo.ThreadEvents {
			fmt.Println(events.MessageEventDetail.AttachedMediaPath)
			handlers.MessageAttachment(oauth, events.MessageEventDetail.AttachedMediaPath, usernam+"_"+strconv.Itoa(id))
		}
		fmt.Println(threadInfo.MaxEventIndexCursor)
	}
	//threadInfo, err := handlers.MessageThreadInfo(oauth, "~1492BCCAACF017DE.3003C1AB75ABBEF0")
	//fmt.Println(threadInfo, err)
	//handlers.MessageAttachment(oauth, "https://us-gmsgp.np.community.playstation.net/groupMessaging/resources/eae1492BCCAACF017DE.3003C1AB75ABBEF0/event/397041907835503/1550944952482", "newAttachment2")
	//fmt.Println(threads)*/
	//games, err := handlers.UserGames(oauth, "jiegokoji", "1", "0")
	//total := games.TotalResults
	//offset := 0
	//fmt.Println(games)
	/*for total > 100 {
		fmt.Println("request for :", 100, " offset: ", offset)
		handlers.UserGames(oauth, "jiegokoji", "100", strconv.Itoa(offset))
		total = total - 100
		offset = offset + 100
		fmt.Println("total left:", total, " offset: ", offset)
	}
	if total != 0 {
		games, err = handlers.UserGames(oauth, "jiegokoji", "100", strconv.Itoa(offset))
	}
	for _, title := range games.TrophyTitles {
		fmt.Println(title.NpCommunicationId)
		fmt.Println(title.TrophyTitleName)
		fmt.Println(title.TrophyTitleIconUrl)
		fmt.Println(title.TrophyTitlePlatfrom)
	}
	/*trophyTitles, err := handlers.GetGameTrophieTitles(oauth, "CUSA12448_00")
	fmt.Println(trophyTitles.Apps[0].TrophyTitles[0].NpCommunicationId)*/
	/*var t2 time.Time
	timeNow := time.Now().UTC().Add(time.Minute * -time.Duration(4880)).Format(time.RFC3339)
	t1, err := time.Parse(time.RFC3339, timeNow)
	if err != nil {
		fmt.Println(err)
		return
	}*/

	alltrophies, err := handlers.GetGameTrophies(oauth, "NPWR01580_00", "jiegokoji")
	fmt.Println(alltrophies)

	//for _, trophie := range alltrophies.Trophies {
	//		if len(trophie.ComparedUser.EarnedDate) > 2 {
	//		t2, err = time.Parse(time.RFC3339, trophie.ComparedUser.EarnedDate)
	//	if err != nil {
	//	fmt.Println(err)
	//return
	//}
	//}
	//if trophie.ComparedUser.OnlineId != "" && trophie.ComparedUser.Earned == true && inTimeSpan(t1, t2) {
	//fmt.Println(trophie.TrophyName)
	//fmt.Println(trophie.ComparedUser.EarnedDate)
	//fmt.Println(trophie.TrophyType)
	//fmt.Println(trophie.TrophyDetail)
	//.Println(trophie)
	//}
	//}
	//groups, err := handlers.GetGameTrophieGroups(oauth, "NPWR16409_00")
	//fmt.Println(groups)
	//	alltrophies, err := handlers.GetGameTrophieData(oauth, "NPWR16409_00", "default")*/
	//fmt.Println(alltrophies)

}
func inTimeSpan(start, check time.Time) bool {
	//fmt.Println(check.After(start), "  ", check.Before(end))
	return check.After(start)
}
