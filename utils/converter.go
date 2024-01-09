package utils

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//func LocationToKM(c *gin.Context, userLatitude, userLongitude, campaignLatitude, campaignLongitude string) string {
//	baseURL := "https://maps.googleapis.com/maps/api/distancematrix/json?"
//	params := url.Values{}
//	params.Set("origins", userLatitude+","+userLongitude)
//	params.Set("destinations", campaignLatitude+","+campaignLongitude)
//	params.Set("mode", "driving")
//	params.Set("key", os.Getenv("GOOGLE_API_KEY"))
//
//	requestURL := baseURL + params.Encode()
//
//	resp, err := http.Get(requestURL)
//	if err != nil {
//		fmt.Println("Error sending request:", err)
//		return "0"
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("Error reading response:", err)
//		return "0"
//	}
//
//	var distanceMatrixResponse model.DistanceMatrixResponse
//	err = json.Unmarshal(body, &distanceMatrixResponse)
//	if err != nil {
//		fmt.Println("Error parsing JSON:", err)
//		return "0"
//	}
//
//	distanceStr := distanceMatrixResponse.Rows[0].Elements[0].Distance.Text
//
//	return distanceStr
//}

func StringToInteger(input string, c *gin.Context) int {
	converted, err := strconv.Atoi(input)
	if err != nil {
		HttpRespFailed(c, http.StatusBadRequest, "Invalid input")
		return 0
	}
	return converted
}

func StringToFloat(input string, c *gin.Context) float64 {
	converted, err := strconv.ParseFloat(input, 64)
	if err != nil {
		HttpRespFailed(c, http.StatusBadRequest, "Invalid input")
		return 0
	}
	return converted
}

func StringToUint(input string, c *gin.Context) uint {
	converted, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		HttpRespFailed(c, http.StatusBadRequest, "Invalid input")
		return 0
	}
	return uint(converted)
}

func Float64ToInt(input float64, c *gin.Context) int {
	return int(math.Round(input))
}

func GetOrdinalSuffix(n int) string {
	switch n % 100 {
	case 11, 12, 13:
		return "th"
	default:
		switch n % 10 {
		case 1:
			return "st"
		case 2:
			return "nd"
		case 3:
			return "rd"
		default:
			return "th"
		}
	}
}

func TimeToString(currentTime time.Time) string {
	// Get the day of the month
	day := currentTime.Day()

	// Define the desired layout with ordinal suffix
	layout := fmt.Sprintf("%d%s Jan 2006 | 15.04 WIB", day, GetOrdinalSuffix(day))

	// Convert the time to the desired layout
	formattedTime := currentTime.Format(layout)

	return formattedTime
}

func IntToRupiah(value int64) string {
	printer := message.NewPrinter(language.Indonesian)
	return printer.Sprintf("%d", value)
}
