package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	//	"io/ioutil"
	"net/http"
	"net/url"
)

var login_request = map[string]string{
	"authentication_type": "password",
	"username":            "",
	"password":            "",
	"client_id":           "71a7beb8-f21a-47d9-a604-2e71bee24fe0",
}

var two_factor_login_request = map[string]string{
	"authentication_type": "two_step",
	"ticket_uuid":         "d92c734d-f861-4726-8efc-00d57f5ae143",
	"code":                "",
	"client_id":           "ebee17ac-99fd-487c-9b1e-18ef50c39ab5",
}

var oauth_request = map[string]string{
	"app_context":   "inapp_ios",
	"client_id":     "ebee17ac-99fd-487c-9b1e-18ef50c39ab5",
	"client_secret": "e4Ru_s*LrL4_B2BD",
	"refresh_token": "0d0fca9f-4bc1-4bd9-b488-2876984a6ad3",
	"duid":          "0000000d00040080027BC1C3FBB84112BFC9A4300A78E96A",
	"grant_type":    "refresh_token",
	"scope":         "kamaji:get_players_met kamaji:get_account_hash kamaji:activity_feed_submit_feed_story kamaji:activity_feed_internal_feed_submit_story kamaji:activity_feed_get_news_feed kamaji:communities kamaji:game_list kamaji:ugc:distributor oauth:manage_device_usercodes psn:sceapp user:account.profile.get user:account.attributes.validate user:account.settings.privacy.get kamaji:activity_feed_set_feed_privacy kamaji:satchel kamaji:satchel_delete user:account.profile.update",
}

var code_request = map[string]string{
	"state":         "06d7AuZpOmJAwYYOWmVU63OMY",
	"duid":          "0000000d000400808F4B3AA3301B4945B2E3636E38C0DDFC",
	"app_context":   "inapp_ios",
	"client_id":     "b7cbf451-6bb6-4a5a-8913-71e61f462787",
	"scope":         "capone:report_submission,psn:sceapp,user:account.get,user:account.settings.privacy.get,user:account.settings.privacy.update,user:account.realName.get,user:account.realName.update,kamaji:get_account_hash,kamaji:ugc:distributor,oauth:manage_device_usercodes",
	"response_type": "code",
}

type login_response struct {
	Npsso string `json:"npsso"`
}

type oauth_response struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

type login_response_fail struct {
	Error            string        `json:"error"`
	ErrorDescription string        `json:"error_description"`
	ErrorCode        int           `json:"error_code"`
	Docs             string        `json:"docs"`
	Parameters       []interface{} `json:"parameters"`
}

//Takes a map of strings with the request parameters and returns a request string
//Implements the http_build_query from PHP into GO
//Written by Tustin
func http_build_query(data map[string]string) string {
	var res bytes.Buffer
	for k, v := range data {
		res.WriteString(k)
		res.WriteByte('=')
		res.WriteString(url.QueryEscape(v))
		res.WriteByte('&')
	}
	s := res.String()
	return s[0 : len(s)-1]
}

func Login(refresh_token string) (oauth_response, error) {
	oauth, err := a2fact(refresh_token)
	if err != nil {
		return oauth_response{}, fmt.Errorf("could not obtain oauth token: %v", err)
	}
	//	fmt.Println(npsso)
	/*if err != nil {
		return oauth_response{}, fmt.Errorf("could not obtain npsso: %v", err)
	}

	grant_code, err := GrabCode(npsso)

	if err != nil {
		return oauth_response{}, fmt.Errorf("could not obtain grant code: %v", err)
	}

	oauth, err := GrabOAuth(npsso, grant_code)

	if err != nil {
		return oauth_response{}, fmt.Errorf("could not obtain oauth token: %v", err)
	}*/

	return oauth, err
}

func GrabOAuth(npsso string, grant_code string) (oauth_response, error) {
	cookie := http.Cookie{Name: "npsso", Value: npsso}
	oauth_request["code"] = grant_code

	req, err := http.NewRequest("POST", OAUTH_URL, bytes.NewBufferString(http_build_query(oauth_request)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(&cookie)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return oauth_response{}, err
	}

	var oa oauth_response
	err = json.NewDecoder(res.Body).Decode(&oa)
	defer res.Body.Close()

	if err != nil {
		return oauth_response{}, err
	}

	return oa, nil
}
func GrabCode(npsso string) (string, error) {

	cookie := http.Cookie{Name: "npsso", Value: npsso}
	url := fmt.Sprintf("%s?%s", CODE_URL, http_build_query(code_request))
	req, err := http.NewRequest("GET", url, nil)
	req.AddCookie(&cookie)

	//Need to use the RoundTripper for this request because the response returns a 304 code and the http.Client automatically follows it
	//We don't want this to happen because we need the X-NP-GRANT-CODE from the response header
	var DefaultTransport http.RoundTripper = &http.Transport{}
	resp, err := DefaultTransport.RoundTrip(req)

	if err != nil {
		return "", err
	}

	header := resp.Header
	grant_code := header.Get("X-NP-GRANT-CODE")

	if grant_code == "" {
		return "", fmt.Errorf("unable to fetch X-NP-GRANT-CODE header")
	}

	return grant_code, nil
}
func a2fact(refresh_token string) (oauth_response, error) {
	oauth_request["refresh_token"] = refresh_token
	req, err := http.NewRequest("POST", OAUTH_URL, bytes.NewBufferString(http_build_query(oauth_request)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return oauth_response{}, err
	}

	var oa oauth_response
	err = json.NewDecoder(resp.Body).Decode(&oa)
	defer resp.Body.Close()

	if err != nil {
		return oauth_response{}, err
	}

	return oa, nil
}
func GrabNPSSO() (string, error) {
	req, err := http.NewRequest("POST", SSO_URL, bytes.NewBufferString(http_build_query(login_request)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	//API didn't give an OK, so handle it as an error
	if resp.StatusCode != http.StatusOK {
		var api_error login_response_fail
		err = json.NewDecoder(resp.Body).Decode(&api_error)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf(api_error.ErrorDescription)
	}

	var res login_response
	err = json.NewDecoder(resp.Body).Decode(&res)

	if err != nil {
		return "", err
	}

	return res.Npsso, nil
}
