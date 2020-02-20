package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go/service/sts"

	"github.com/YashdalfTheGray/federator/constants"
)

// GetSessionString returns a JSON.stringified representation of
// the session object
func GetSessionString(creds *sts.AssumeRoleOutput) string {
	session := struct {
		SessionID    string `json:"sessionId"`
		SessionKey   string `json:"sessionKey"`
		SessionToken string `json:"sessionToken"`
	}{
		SessionID:    *creds.Credentials.AccessKeyId,
		SessionKey:   *creds.Credentials.SecretAccessKey,
		SessionToken: *creds.Credentials.SessionToken,
	}

	sessionStr, err := json.Marshal(session)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return string(sessionStr)
}

// GetSigninTokenURL builds a url.URL object using the particulars from the
// session string and the federation URL
func GetSigninTokenURL(creds *sts.AssumeRoleOutput) url.URL {
	u, err := url.Parse(constants.FederationEndpoint)
	if err != nil {
		log.Fatalln(err.Error())
	}

	q := u.Query()
	q.Set("Action", "getSigninToken")
	q.Set("SessionDuration", "3600")
	q.Set("Session", GetSessionString(creds))

	u.RawQuery = q.Encode()
	return *u
}

// GetSigninToken uses the signin URL and calls it to get the user a signin
// token
func GetSigninToken(signinURL url.URL) string {
	var signinResponse struct {
		SigninToken string `json:"SigninToken"`
	}
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, signinReqErr := client.Get(signinURL.String())
	if signinReqErr != nil {
		log.Fatalln(signinReqErr.Error())
	}
	defer resp.Body.Close()

	body, readBodyErr := ioutil.ReadAll(resp.Body)
	if readBodyErr != nil {
		log.Fatalln(readBodyErr.Error())
	}

	unmarshalErr := json.Unmarshal(body, &signinResponse)
	if unmarshalErr != nil {
		log.Fatalln(unmarshalErr.Error())
	}

	return signinResponse.SigninToken
}

// GetLoginURL builds the console login URL after all of the federation is
// done and returns the URL object
func GetLoginURL(signinToken string) url.URL {
	u, err := url.Parse(constants.FederationEndpoint)
	if err != nil {
		log.Fatalln(err.Error())
	}

	q := u.Query()
	q.Set("Action", "login")
	q.Set("Issuer", constants.DefaultIssuer)
	q.Set("Destination", constants.DefaultDestination)
	q.Set("SigninToken", signinToken)

	u.RawQuery = q.Encode()
	return *u
}