package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var robotNames = make(map[string]bool)

func (robot *Robot) Name() (string, error) {
	if robot.name == "" {
		robot.name = ""
		rand.Seed(time.Now().UnixNano())
		for inUse, ok := string(robot.name) == "", true; ok && inUse; inUse, ok = robotNames[robot.name] {
			robot.name = letter() + letter() + number()
		}
		robotNames[robot.name] = true
	}
	return string(robot.name), nil
}

func (robot *Robot) Reset() {
	robot.name = ""
}

func letter() string {
	return string(rand.Intn('Z'-'A') + 'A')
}

func number() string {
	return fmt.Sprintf("%03d", rand.Intn(1000))
}
