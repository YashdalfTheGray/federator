package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/sts"

	"github.com/YashdalfTheGray/federator/models"
)

// PrintCredsFromSTSResponse prints out the credentials we got from the
// STS output in a way that the user can export them in to the shell
// as well as the expiration information about the session
func PrintCredsFromSTSResponse(out *sts.AssumeRoleOutput, outputJSON bool) {
	if os.Getenv("CI_MODE") == "true" {
		if outputJSON {
			fmt.Println("<Running in quiet mode because of CI but would print JSON>")
		} else {
			fmt.Println("<Running in quiet mode because of CI>")
		}
		return
	}

	credsDetails := models.NewCredsDetails(out)

	if outputJSON {
		jsonOutput, err := credsDetails.ToJSONString()
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Println(jsonOutput)
	} else {
		fmt.Println("Successfully authenticated with STS. Commands to use below.")
		fmt.Println(credsDetails.ToString())
	}
}

// PrintLoginURLDetails prints out the login URL as well as the expiration date
// of the session
func PrintLoginURLDetails(out *sts.AssumeRoleOutput, loginURL string, outputJSON bool) {
	if os.Getenv("CI_MODE") == "true" {
		if outputJSON {
			fmt.Println("<Running in quiet mode because of CI but would print JSON>")
		} else {
			fmt.Println("<Running in quiet mode because of CI>")
		}
		return
	}

	linkDetails := models.NewLinkDetails(out.Credentials.Expiration, loginURL)

	if outputJSON {
		jsonOutput, err := linkDetails.ToJSONString()
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Println(jsonOutput)
	} else {
		fmt.Println("Successfully authenticated with STS. Login URL below.")
		fmt.Println(linkDetails.ToString())
	}
}
