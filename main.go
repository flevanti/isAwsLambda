package isAwsLambda

import "os"

var isLambda bool
var isDocker bool
var isAWS bool
var isInitialised bool

func init() {
	initialise()
}

func IsItLambda() bool {
	return isLambda
}

func IsItDocker() bool {
	return isDocker
}

func IsItAWS() bool {
	return isAWS
}

func IsItInitialised() bool {
	return isInitialised
}

func initialise() {
	if isInitialised {
		return
	}

	//default value
	isLambda = false
	isDocker = false
	isAWS = false

	if len(os.Getenv("AWS_REGION")) != 0 {
		isLambda = true
	}
	//even if we are in an AWS/LAMBDA environment, it could be a docker container...
	//so let's use another env var to understand if docker
	//after comparing docker lambda and AWS lambda I noticed that AWS_SESSION_TOKEN env var is (for the moment) only available in AWS
	if isLambda && len(os.Getenv("AWS_SESSION_TOKEN")) == 0 {
		isDocker = true
	}
	//Try to understand if we are running in AWS
	isAWS = isLambda && !isDocker

	isInitialised = true

}
