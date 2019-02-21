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
type thread_Info struct {
	ThreadMembers []struct {
		AccountId string `json:"accountId"`
		OnlineId  string `json:"onlineId"`
	} `json:"threadMembers"`
	ThreadNameDetail struct {
		Status     int    `json:"status"`
		ThreadName string `json:"threadName"`
	} `json:"threadNameDetail"`
	ThreadThumbnailDetail struct {
		Status int `json:"status"`
	} `json:"threadThumbnailDetail"`
	ThreadProperty struct {
		FavoriteDetail struct {
			FavoriteFlag bool `json:"favoriteFlag"`
		} `json:"favoriteDetail"`
		NotificationDetail struct {
			PushNotificationFlag bool `json:"pushNotificationFlag"`
		} `json:"notificationDetail"`
		KickoutFlag    bool   `json:"kickoutFlag"`
		ThreadJoinDate string `json:"threadJoinDate"`
	} `json:"threadProperty"`
	NewArrivalEventDetail struct {
		NewArrivalEventFlag bool   `json:"newArrivalEventFlag"`
		EventIndex          string `json:"eventIndex"`
	} `json:"newArrivalEventDetail"`
	ThreadEvents []struct {
		MessageEventDetail struct {
			EventIndex           string `json:"eventIndex"`
			PostDate             string `json:"postDate"`
			EventCategoryCode    int    `json:"eventCategoryCode"`
			AltEventCategoryCode int    `json:"altEventCategoryCode"`
			Sender               struct {
				AccountId string `json:"accountId"`
				OnlineId  string `json:"onlineId"`
			} `json:"sender"`
			AttachedMediaPath string `json:"attachedMediaPath"`
			MessageDetail     struct {
				MessageSubject     string `json:"messageSubject"`
				Body               string `json:"body"`
				TransferredFromPS3 bool   `json:"transferredFromPS3"`
			} `json:"messageDetail"`
		} `json:"messageEventDetail"`
	} `json:"threadEvents"`
	ThreadId              string `json:"threadId"`
	ThreadType            int    `json:"threadType"`
	ThreadModifiedDate    string `json:"threadModifiedDate"`
	ResultsCount          int    `json:"resultsCount"`
	MaxEventIndexCursor   string `json:"maxEventIndexCursor"`
	SinceEventIndexCursor string `json:"sinceEventIndexCursor"`
	LatestEventIndex      string `json:"latestEventIndex"`
	EndOfThreadEvent      bool   `json:"endOfThreadEvent"`
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
func MessageThreadInfo(oauth oauth_response, threadId string) (thread_Info, error) {
	client := &http.Client{}
	var url string
	url = MESSAGE_THREAD_ENDPOINT + "threads/" + threadId + "?fields=threadMembers,threadNameDetail,threadThumbnailDetail,threadProperty,latestTakedownEventDetail,newArrivalEventDetail,threadEvents&count=1"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+oauth.AccessToken)
	resp, err := client.Do(req)

	if err != nil {
		return thread_Info{}, err
	}
	defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	if resp.StatusCode != http.StatusOK {
		var api_error user_error
		err := json.NewDecoder(resp.Body).Decode(&api_error)
		if err != nil {
			return thread_Info{}, err
		}
		return thread_Info{}, fmt.Errorf(api_error.Error.Message)
	}

	var threadInfo thread_Info
	err = json.NewDecoder(resp.Body).Decode(&threadInfo)

	if err != nil {
		return thread_Info{}, err
	}
	return threadInfo, nil
}
