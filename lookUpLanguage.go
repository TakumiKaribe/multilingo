package main

import "fmt"

func lookUpLanguage(request *APIGateWayRequest) (string, error) {
	switch request.ApiAppID {
	case "AG6LQER0B":
		return "swift", nil
	default:
		return "", fmt.Errorf("No language corresponding to %s was found", request.ApiAppID)
	}
}
