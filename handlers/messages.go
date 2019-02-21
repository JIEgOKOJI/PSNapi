package handlers

import (
	//"PSNapi/handlers/auth"
	//	"bytes"
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
	//	"net/url"
)

type thread_Ids struct {
	ThreadIds []struct {
		ThreadId           string `json:"threadId"`
		ThreadModifiedDate string `json:"threadModifiedDate"`
	} `json:"threadIds"`
}

func MessageThreads(oauth oauth_response, OnlineId string) (thread_Ids, error) {
	client := &http.Client{}
	var url string
	url = MESSAGE_THREAD_ENDPOINT + "users/me/threadIds?withOnlineIds=" + OnlineId
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		return thread_Ids{}, err
	}
	defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	if resp.StatusCode != http.StatusOK {
		var api_error user_error
		err := json.NewDecoder(resp.Body).Decode(&api_error)
		if err != nil {
			return thread_Ids{}, err
		}
		return thread_Ids{}, fmt.Errorf(api_error.Error.Message)
	}

	var threadIds thread_Ids
	err = json.NewDecoder(resp.Body).Decode(&threadIds)

	if err != nil {
		return thread_Ids{}, err
	}
	return threadIds, nil
}
