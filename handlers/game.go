package handlers

import (
	//"PSNapi/handlers/auth"
	"encoding/json"
	"fmt"

	//	"io/ioutil"
	"net/http"
	//	"net/url"
)

type trophy_titles struct {
	Apps []struct {
		NpTitleId    string `json:"npTitleId"`
		TrophyTitles []struct {
			NpCommunicationId   string `json:"npCommunicationId"`
			TrophyTitleName     string `json:"trophyTitleName"`
			TrophyTitleIconUrls []struct {
				Size               string `json:"size"`
				TrophyTitleIconUrl string `json:"trophyTitleIconUrl"`
			} `json:"trophyTitleIconUrls"`
			Users []interface{} `json:"users"`
		} `json:"trophyTitles"`
	} `json:"apps"`
}
type all_trophies struct {
	Trophies []struct {
		TrophyId      int    `json:"trophyId"`
		TrophyHidden  bool   `json:"trophyHidden"`
		TrophyType    string `json:"trophyType"`
		TrophyName    string `json:"trophyName"`
		TrophyDetail  string `json:"trophyDetail"`
		TrophyIconUrl string `json:"trophyIconUrl"`
		TrophyRare    int    `json:"TrophyRare"`
		ComparedUser  struct {
			OnlineId   string `json:"onlineId"`
			Earned     bool   `json:"earned"`
			EarnedDate string `json:"earnedDate"`
		} `json:"comparedUser"`
	} `json:"trophies"`
}
type all_trophies_data struct {
	Trophies []struct {
		TrophyId         int    `json:"trophyId"`
		TrophyHidden     bool   `json:"trophyHidden"`
		TrophyType       string `json:"trophyType"`
		TrophyName       string `json:"trophyName"`
		TrophyDetail     string `json:"trophyDetail"`
		TrophyIconUrl    string `json:"trophyIconUrl"`
		TrophyRare       int    `json:"TrophyRare"`
		TrophyEarnedRate string `json:"trophyEarnedRate"`
	} `json:"trophies"`
}
type trophie_groups struct {
	TrophyTitleName     string `json:"trophyTitleName"`
	TrophyTitleDetail   string `json:"trophyTitleDetail"`
	TrophyTitleIconUrl  string `json:"trophyTitleIconUrl"`
	TrophyTitlePlatfrom string `json:"trophyTitlePlatfrom"`
	DefinedTrophies     struct {
		Bronze   int `json:"bronze"`
		Silver   int `json:"silver"`
		Gold     int `json:"gold"`
		Platinum int `json:"platinum"`
	} `json:"definedTrophies"`
	TrophyGroups []struct {
		TrophyGroupId      string `json:"trophyGroupId"`
		TrophyGroupName    string `json:"trophyGroupName"`
		TrophyGroupDetail  string `json:"trophyType"`
		TrophyGroupIconUrl string `json:"trophyGroupDetail"`
		DefinedTrophies    struct {
			Bronze   int `json:"bronze"`
			Silver   int `json:"silver"`
			Gold     int `json:"gold"`
			Platinum int `json:"platinum"`
		} `json:"definedTrophies"`
	} `json:"trophyGroups"`
}

func GetGameTrophieTitles(oauth oauth_response, TitleId string) (trophy_titles, error) {
	client := &http.Client{}
	url := TROPHY_ENDPOINT + "apps/trophyTitles" + "?npTitleIds=" + TitleId + "&fields=%40default&npLanguage=en"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		return trophy_titles{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		var api_error user_error
		err := json.NewDecoder(resp.Body).Decode(&api_error)
		if err != nil {
			return trophy_titles{}, err
		}
		return trophy_titles{}, fmt.Errorf(api_error.Error.Message)
	}

	var ttitles trophy_titles
	err = json.NewDecoder(resp.Body).Decode(&ttitles)

	if err != nil {
		return trophy_titles{}, err
	}
	return ttitles, nil
}
func GetGameTrophies(oauth oauth_response, npCommunicationId string, comparedUser string) (all_trophies, error) {
	client := &http.Client{}
	var url string
	//var Url *url.URL
	//Url, err := url.Parse(TROPHY_ENDPOINT)
	if comparedUser == "" {
		url = TROPHY_ENDPOINT + "trophyTitles/" + npCommunicationId + "/trophyGroups/all/trophies?fileds=%40default,trophyRare,trophyEarnedRate,hasTrophyGroups,trophySmallIconUrl&iconSize=m&visibleType=1&npLanguage=en"
	} else {
		url = TROPHY_ENDPOINT + "trophyTitles/" + npCommunicationId + "/trophyGroups/all/trophies?fileds=%40default,trophyRare,trophyEarnedRate,hasTrophyGroups,trophySmallIconUrl&iconSize=m&visibleType=1&npLanguage=en&comparedUser=" + comparedUser
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		return all_trophies{}, err
	}
	defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	if resp.StatusCode != http.StatusOK {
		var api_error user_error
		err := json.NewDecoder(resp.Body).Decode(&api_error)
		if err != nil {
			return all_trophies{}, err
		}
		return all_trophies{}, fmt.Errorf(api_error.Error.Message)
	}

	var alltrophies all_trophies
	err = json.NewDecoder(resp.Body).Decode(&alltrophies)

	if err != nil {
		return all_trophies{}, err
	}
	return alltrophies, nil
}
func GetGameTrophieGroups(oauth oauth_response, npCommunicationId string) (trophie_groups, error) {
	client := &http.Client{}
	var url string
	//var Url *url.URL
	//Url, err := url.Parse(TROPHY_ENDPOINT)
	url = TROPHY_ENDPOINT + "trophyTitles/" + npCommunicationId + "/trophyGroups/?npLanguage=en"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		return trophie_groups{}, err
	}
	defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	if resp.StatusCode != http.StatusOK {
		var api_error user_error
		err := json.NewDecoder(resp.Body).Decode(&api_error)
		if err != nil {
			return trophie_groups{}, err
		}
		return trophie_groups{}, fmt.Errorf(api_error.Error.Message)
	}

	var groups trophie_groups
	err = json.NewDecoder(resp.Body).Decode(&groups)

	if err != nil {
		return trophie_groups{}, err
	}
	return groups, nil
}
func GetGameTrophieData(oauth oauth_response, npCommunicationId string, groupId string) (all_trophies_data, error) {
	client := &http.Client{}
	var url string
	//var Url *url.URL
	//Url, err := url.Parse(TROPHY_ENDPOINT)
	url = TROPHY_ENDPOINT + "trophyTitles/" + npCommunicationId + "/trophyGroups/all/trophies?fields=%40default,trophyRare,trophyEarnedRate&npLanguage=en"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		return all_trophies_data{}, err
	}
	defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(body))
	if resp.StatusCode != http.StatusOK {
		var api_error user_error
		err := json.NewDecoder(resp.Body).Decode(&api_error)
		if err != nil {
			return all_trophies_data{}, err
		}
		return all_trophies_data{}, fmt.Errorf(api_error.Error.Message)
	}

	var trophies all_trophies_data
	err = json.NewDecoder(resp.Body).Decode(&trophies)

	if err != nil {
		return all_trophies_data{}, err
	}
	return trophies, nil
}
