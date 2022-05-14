package main

import "simple-rest-client/restapi"

//Test url's
//https://www.postman.com/collections/bf7d9130241aaa7160d8

func main() {

	req := restapi.Request{
		Method: "GET",
		Url:    `https://services.odata.org/V4/TripPinService/Airports?$filter=Location/City/Region%20eq%20%27California%27`,
	}

	req.SendRequest()
}
