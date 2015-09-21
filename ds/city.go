package ds

import (
	"github.com/boltdb/bolt"
	"strconv"
	"strings"
)

type City struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	CountryCode string   `json:"-"`
	Population  uint32   `json:"population"`
	Latitude    string   `json:"latitude"`
	Longitude   string   `json:"longitude"`
	Timezone    string   `json:"timezone"`
	Country     *Country `json:"country,omitempty"`
}

func cityFromString(id string, cityString string) (*City, error) {
	var city City
	var err error

	cityData := strings.Split(cityString, "\t")

	if len(cityData) == 6 {
		var population int64
		population, err = strconv.ParseInt(cityData[2], 0, 64)

		city.Id = id
		city.Name = cityData[0]
		city.CountryCode = cityData[1]
		city.Population = uint32(population)
		city.Latitude = cityData[3]
		city.Longitude = cityData[4]
		city.Timezone = cityData[5]
	} else {
		err = InvalidDataError{CitiesBucketName, id, cityString}
	}

	return &city, err
}

func (city *City) toString() string {
	return city.Name + "\t" + city.CountryCode + "\t" +
		strconv.Itoa(int(city.Population)) + "\t" + city.Latitude +
		"\t" + city.Longitude + "\t" + city.Timezone
}

func FindCity(db *bolt.DB, id string, includeCountry bool) (*City, error) {
	var city *City = nil

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(CitiesBucketName)
		val := bucket.Get([]byte(id))
		var err error

		if val != nil {
			city, err = cityFromString(id, string(val))
			if err == nil && includeCountry == true {
				city.Country, err = FindCountryByCode(db, city.CountryCode)
			}
		}
		return err
	})

	return city, err
}
