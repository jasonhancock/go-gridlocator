package gridlocator

import (
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Convert converts the specified decimanl longitude and latitude into the six
// digit Maidenhead grid locator.
func Convert(latitude, longitude float64) (string, error) {

	lat := latitude + 90
	lng := longitude + 180

	// Field
	lat = (lat / 10) // + 0.0000001;
	lng = (lng / 20) // + 0.0000001;
	val, err := upperN2L(int(math.Floor(lng)))
	if err != nil {
		return "", errors.Wrap(err, "field longitude")
	}
	locator := val
	val, err = upperN2L(int(math.Floor(lat)))
	if err != nil {
		return "", errors.Wrap(err, "field latitude")
	}
	locator += val

	// Square
	lat = 10 * (lat - math.Floor(lat))
	lng = 10 * (lng - math.Floor(lng))
	locator += strconv.Itoa(int(math.Floor(lng)))
	locator += strconv.Itoa(int(math.Floor(lat)))

	// Subsquare
	lat = 24 * (lat - math.Floor(lat))
	lng = 24 * (lng - math.Floor(lng))
	val, err = n2l(int(math.Floor(lng)))
	if err != nil {
		return "", errors.Wrap(err, "subsquare longitude")
	}
	locator += val
	val, err = n2l(int(math.Floor(lat)))
	if err != nil {
		return "", errors.Wrap(err, "subsquare latitude")
	}
	locator += val

	return locator, nil
}

// ConvertGridLocation converts a string grid location into latitude and longitude.
func ConvertGridLocation(location string) (float64, float64, error) {
	if len(location) != 4 && len(location) != 6 {
		return 0, 0, errors.New("grid location must be either 4 or 6 digits")
	}

	location = strings.ToLower(location)

	//lng = (($l[0] * 20) + ($l[2] * 2) + ($l[4]/12)  - 180);
	l := make([]int, 6)

	// Field
	var err error
	l[0], err = l2n(string(location[0]))
	if err != nil {
		return 0, 0, errors.Wrap(err, "longitude field value")
	}
	l[1], err = l2n(string(location[1]))
	if err != nil {
		return 0, 0, errors.Wrap(err, "latitude field value")
	}

	// Square
	val, err := strconv.ParseInt(string(location[2]), 10, 64)
	if err != nil {
		return 0, 0, errors.Wrap(err, "longitude sqare value")
	}
	l[2] = int(val)

	val, err = strconv.ParseInt(string(location[3]), 10, 64)
	if err != nil {
		return 0, 0, errors.Wrap(err, "latitude sqare value")
	}
	l[3] = int(val)

	if len(location) == 6 {
		// Subsquare
		l[4], err = l2n(string(location[4]))
		if err != nil {
			return 0, 0, errors.Wrap(err, "longitude subsquare value")
		}
		l[5], err = l2n(string(location[5]))
		if err != nil {
			return 0, 0, errors.Wrap(err, "latitude subsquare value")
		}
	}

	long := (float64(l[0]) * 20) + (float64(l[2]) * 2) + (float64(l[4]) / 12) - 180
	lat := (float64(l[1]) * 10) + float64(l[3]) + (float64(l[5]) / 24) - 90

	return lat, long, nil
}

var num2let = []string{
	`a`,
	`b`,
	`c`,
	`d`,
	`e`,
	`f`,
	`g`,
	`h`,
	`i`,
	`j`,
	`k`,
	`l`,
	`m`,
	`n`,
	`o`,
	`p`,
	`q`,
	`r`,
	`s`,
	`t`,
	`u`,
	`v`,
	`w`,
	`x`,
}

func n2l(number int) (string, error) {
	if number > (len(num2let) - 1) {
		return "", errors.New("number out of bounds")
	}

	return num2let[number], nil
}

func upperN2L(number int) (string, error) {
	val, err := n2l(number)
	return strings.ToUpper(val), err
}

var let2num = map[string]int{
	`a`: 0,
	`b`: 1,
	`c`: 2,
	`d`: 3,
	`e`: 4,
	`f`: 5,
	`g`: 6,
	`h`: 7,
	`i`: 8,
	`j`: 9,
	`k`: 10,
	`l`: 11,
	`m`: 12,
	`n`: 13,
	`o`: 14,
	`p`: 15,
	`q`: 16,
	`r`: 17,
	`s`: 18,
	`t`: 19,
	`u`: 20,
	`v`: 21,
	`w`: 22,
	`x`: 23,
}

func l2n(letter string) (int, error) {
	val, ok := let2num[letter]
	if !ok {
		return 0, errors.New("illegal character")
	}
	return val, nil
}
