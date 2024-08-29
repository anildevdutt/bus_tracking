package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type BusTraking struct {
	Status     int `json:"status"`
	EtaMapData []struct {
		GeofenceName     string      `json:"geofence_name"`
		ID               int         `json:"id"`
		ScheduledTime    string      `json:"scheduled_time"`
		ServicePlaceName string      `json:"service_place_name"`
		Color            string      `json:"color"`
		DelayTime        interface{} `json:"delay_time"`
		ExpectedTime     string      `json:"expected_time"`
		Skipped          bool        `json:"skipped"`
		RunningStatus    interface{} `json:"running_status"`
		IsPickUp         int         `json:"is_pick_up"`
		ArrivalTime      string      `json:"arrival_time"`
		DepartureTime    string      `json:"departure_time"`
	} `json:"eta_map_data"`
	EtaPickupData                []interface{} `json:"eta_pickup_data"`
	TravellerPickupServicePlaces interface{}   `json:"traveller_pickup_service_places"`
	CurrentSpID                  int           `json:"current_sp_id"`
	IsPassed                     bool          `json:"is_passed"`
	CurrentStatusDetails         struct {
		LatLong []float64 `json:"lat_long"`
		Details struct {
			Speed     int    `json:"speed"`
			Timestamp string `json:"timestamp"`
			Location  string `json:"location"`
			AstlID    int    `json:"astl_id"`
			ClassName string `json:"class_name"`
		} `json:"details"`
	} `json:"current_status_details"`
	DistanceDetails struct {
	} `json:"distance_details"`
	LastDropoffID        int `json:"last_dropoff_id"`
	DistCurAndSecondNext struct {
	} `json:"dist_cur_and_second_next"`
	LastCrossedPickup  interface{} `json:"last_crossed_pickup"`
	IsPassedPickup     interface{} `json:"is_passed_pickup"`
	SttPickupAssetName interface{} `json:"stt_pickup_asset_name"`
	LastBoardingID     int         `json:"last_boarding_id"`
}

func main() {

	trackurl := "https://bus.trackingo.in/api/live/eta_map?current_status=true&key=PZEYHJ"
	req, _ := http.NewRequest("GET", trackurl, nil)
	client := &http.Client{}
	var track BusTraking
	for {
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		data, err := io.ReadAll(resp.Body)
		f, err := os.OpenFile("data.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = f.Write(data)
		if err != nil {
			log.Fatalln(err)
		}
		f.Close()
		json.Unmarshal(data, &track)
		log.Println(track.CurrentStatusDetails.LatLong, track.CurrentStatusDetails.Details.Speed)
		resp.Body.Close()
		time.Sleep(30 * time.Second)
	}

}
