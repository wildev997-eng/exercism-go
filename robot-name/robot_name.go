package robotname

import (
	"fmt"
	"math/rand"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var availableName = []string{}

func AvailableName() {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for first := 0; first < len(alphabet); first++ {
		for second := 0; second < len(alphabet); second++ {
			for number := 0; number <= 999; number++ {
				name := string(alphabet[first]) +
					string(alphabet[second]) +
					fmt.Sprintf("%03d", number)

				availableName = append(availableName, name)
			}
		}
	}
}

func init() {
	AvailableName()
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(availableName) == 0 {
		return "", fmt.Errorf("no robot names available")
	}
	index := rand.Intn(len(availableName))
	candidate := availableName[index]
	last := len(availableName) - 1
	availableName[index] = availableName[last]
	availableName = availableName[:last]
	r.name = candidate
	return r.name, nil

}

func (r *Robot) Reset() {
	r.name = ""
}

//Name() function without the bonus test

// if r.name != "" {
// 	return r.name, nil
// }

// for attempts := 0; attempts < 676000; attempts++ {
// 	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	twoAlpha := string(alphabet[rand.Intn(len(alphabet))]) + string(alphabet[rand.Intn(len(alphabet))])
// 	randNum := rand.Intn(1000)
// 	candidate := string(twoAlpha) + fmt.Sprintf("%03d", randNum)
// 	if !usedNames[candidate] {
// 		usedNames[candidate] = true
// 		r.name = candidate
// 		return r.name, nil
// 	}

// }
// return "", errors.New("all robot names have been used")
