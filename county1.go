package main

import(
	"fmt"
	"os"
	"bufio"
	//"io"
	"strconv"
	//"io/ioutil"
	"math"
)
func haversine (theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
func distance(county1 vertex, county2 vertex) float64 {
	var lat1, long1, lat2, long2, r float64

	lat1 = county1.lat * math.Pi / 180
	long1 = county1.long * math.Pi / 180
	lat2 = county2.lat * math.Pi / 180
	long2 = county2.long * math.Pi / 180

	r = 3959 // radius of earth in miles

	h := haversine(lat2-lat1) + math.Cos(lat1)*math.Cos(lat2)*haversine(long2-long1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func user_input(county_info map[string]vertex) float64 {
	var county1 vertex
	var county2 vertex
	var distance_sum float64

	distance_sum = 0.0

	fmt.Println("Type county name: ")

	input_scan := bufio.NewScanner(os.Stdin)
	var input_line string 
	input_line = "sample"

	county1.lat = 200
	county2.lat = 200

for input_scan.Scan(){
	input_line = input_scan.Text()

	//fmt.Println(input_line)
	

	if input_line == "" {
		return distance_sum
	}

	if value, ok := county_info[input_line]; ok {
		if county1.lat == 200 {
		county1 = value
		}else { if county2.lat == 200 {
			county2 = value
	}
	}
	fmt.Println(county1)
	fmt.Println(county2)
	
	}

	

		if county1.lat != 200 && county2.lat != 200 {
			//fmt.Println("Calc distance sum!")
			distance_sum = distance_sum + distance(county1, county2)

			county1 = county2
			county2.lat = 200

		}
	}
		return 1


}


type vertex struct {
	lat float64
	long float64
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main(){
	var county_info map[string]vertex
	var coords vertex
	var name string
	
	county_info = make(map[string]vertex)

	file, err := os.Open("latlong.txt")
	check(err)

	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)

	var lat_d string
	var long_d string

	i := 1

	for scan.Scan() {
		if i % 3 == 1 {
			name = scan.Text()
		}

		if i % 3 == 2 {
			lat_d = scan.Text()
		}

		if i % 3 == 0 {
			long_d = scan.Text()
			coords.lat, _ = (strconv.ParseFloat(lat_d, 64))
			coords.long, _ = (strconv.ParseFloat(long_d, 64))
			county_info[name] = coords
		}
		i = i+1
	}

	for k, v := range county_info {
		fmt.Println(k, v)
	}
	dist := (user_input(county_info))
	fmt.Println("The final distance is:")
	fmt.Println(dist)
}	