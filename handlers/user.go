package handlers

import (
	//"PSNapi/handlers/auth"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type user_error struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
type user_games struct {
	TotalResults int `json:"totalResults"`
	Offset       int `json:"offset"`
	Limit        int `json:"limit"`
	TrophyTitles []struct {
		NpCommunicationId   string `json:"npCommunicationId"`
		TrophyTitleName     string `json:"trophyTitleName"`
		TrophyTitleDetail   string `json:"trophyTitleDetail"`
		TrophyTitleIconUrl  string `json:"trophyTitleIconUrl"`
		TrophyTitlePlatfrom string `json:"trophyTitlePlatfrom"`
		HasTrophyGroups     bool   `json:"hasTrophyGroups"`
		DefinedTrophies     struct {
			Platinum int `json:"platinum"`
			Gold     int `json:"gold"`
			Silver   int `json:"silver"`
			Bronze   int `json:"bronze"`
		} `json:"definedTrophies"`
		ComparedUser struct {
			OnlineId       string `json:"onlineId"`
			Progress       int    `json:"progress"`
			EarnedTrophies struct {
				Platinum int `json:"platinum"`
				Gold     int `json:"gold"`
				Silver   int `json:"silver"`
				Bronze   int `json:"bronze"`
			} `json:"earnedTrophies"`
			LastUpdateDate string `json:"lastUpdateDate"`
		} `json:"comparedUser"`
	} `json:"trophyTitles"`
}
type user_profile struct {
	Profile struct {
		OnlineID   string `json:"onlineId"`
		NpID       string `json:"npId"`
		AvatarUrls []struct {
			Size      string `json:"size"`
			AvatarURL string `json:"avatarUrl"`
		} `json:"avatarUrls"`
		Plus          int      `json:"plus"`
		AboutMe       string   `json:"aboutMe"`
		LanguagesUsed []string `json:"languagesUsed"`
		TrophySummary struct {
			Level          int `json:"level"`
			Progress       int `json:"progress"`
			EarnedTrophies struct {
				Platinum int `json:"platinum"`
				Gold     int `json:"gold"`
				Silver   int `json:"silver"`
				Bronze   int `json:"bronze"`
			} `json:"earnedTrophies"`
		} `json:"trophySummary"`
		IsOfficiallyVerified                    bool   `json:"isOfficiallyVerified"`
		PersonalDetailSharing                   string `json:"personalDetailSharing"`
		PersonalDetailSharingRequestMessageFlag bool   `json:"personalDetailSharingRequestMessageFlag"`
		PrimaryOnlineStatus                     string `json:"primaryOnlineStatus"`
		Presences                               []struct {
			OnlineStatus     string `json:"onlineStatus"`
			HasBroadcastData bool   `json:"hasBroadcastData"`
		} `json:"presences"`
		FriendRelation      string `json:"friendRelation"`
		RequestMessageFlag  bool   `json:"requestMessageFlag"`
		Blocking            bool   `json:"blocking"`
		FriendsCount        int    `json:"friendsCount"`
		MutualFriendsCount  int    `json:"mutualFriendsCount"`
		Following           bool   `json:"following"`
		FollowingUsersCount int    `json:"followingUsersCount"`
		FollowerCount       int    `json:"followerCount"`
	} `json:"profile"`
}

//Used for debugging API responses
//data, _ := ioutil.ReadAll(resp.Body)
//fmt.Printf("%s", string(data))

func UserInfo(oauth oauth_response, UserName string) (user_profile, error) {
	client := &http.Client{}
	url := USERS_URL + UserName + "/profile2?fields=npId,onlineId,avatarUrls,plus,aboutMe,languagesUsed,trophySummary(@default,progress,earnedTrophies),isOfficiallyVerified,personalDetail(@default,profilePictureUrls),personalDetailSharing,personalDetailSharingRequestMessageFlag,primaryOnlineStatus,presences(@titleInfo,hasBroadcastData),friendRelation,requestMessageFlag,blocking,mutualFriendsCount,following,followerCount,friendsCount,followingUsersCount&avatarSizes=m,xl&profilePictureSizes=m,xl&languagesUsedLanguageSet=set3&psVitaTitleIcon=circled&titleIconSize=s"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		return user_profile{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var api_error user_error
		err := json.NewDecoder(resp.Body).Decode(&api_error)
		if err != nil {
			return user_profile{}, err
		}
		return user_profile{}, fmt.Errorf(api_error.Error.Message)
	}

	var profile user_profile
	err = json.NewDecoder(resp.Body).Decode(&profile)

	if err != nil {
		return user_profile{}, err
	}

	return profile, nil
}
func UserGames(oauth oauth_response, UserName string, limit string, offset string) (user_games, error) {
	client := &http.Client{}
	var Url *url.URL
	Url, err := url.Parse(GAME_ENDPOINT)
	//Url.Path += "users/" + UserName + "/titles"
	parameters := url.Values{}
	//parameters.Add("type", "played")
	parameters.Add("fields", "@default")
	parameters.Add("app", "richProfile")
	parameters.Add("platform", "PS3,PS4,PSVITA")
	parameters.Add("npLanguage", "en")
	parameters.Add("sort", "-lastPlayedDate")
	parameters.Add("limit", limit)
	parameters.Add("offset", offset)
	parameters.Add("iw", "240")
	parameters.Add("ih", "240")
	parameters.Add("comparedUser", UserName)
	Url.RawQuery = parameters.Encode()
	fmt.Printf("Encoded URL is %q\n", Url.String())
	//uri := GAME_ENDPOINT + "users/" + UserName + "/titles?app=richProfile&ih=240&iw=240&limit=" + limit + "&offset=" + offset + "&platform=PS3%2CPSVITA%2CPS4&sort=-lastPlayedDate&type=played"
	//	fmt.Println(uri)
	req, _ := http.NewRequest("GET", Url.String(), nil)
	req.Header.Set("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		return user_games{}, err
	}
	defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	if resp.StatusCode != http.StatusOK {
		var api_error user_error
		err := json.NewDecoder(resp.Body).Decode(&api_error)
		if err != nil {
			return user_games{}, err
		}
		return user_games{}, fmt.Errorf(api_error.Error.Message)
	}

	var games user_games
	err = json.NewDecoder(resp.Body).Decode(&games)

	if err != nil {
		return user_games{}, err
	}
	return games, nil
}
func UserAddFriend(oauth oauth_response, LoggedPsn string, UserName string, msg string) error {
	client := &http.Client{}
	url := USERS_URL + LoggedPsn + "/friendList/" + UserName
	var jsonStr = []byte(`{"requestMessage":"` + msg + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+oauth.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return nil
}
