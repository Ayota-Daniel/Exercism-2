package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (t TemperatureUnit) String() string {
	temperature := []string{"°C", "°F"}
	return temperature[t]
}

func (t Temperature) String() string {
	return fmt.Sprintf("%v %v", t.degree, t.unit.String())
}

func (s SpeedUnit) String() string {
	speed := []string{"km/h", "mph"}
	return speed[s]
}

func (s Speed) String() string {
	return fmt.Sprintf("%v %v", s.magnitude, s.unit.String())
}

func (m MeteorologyData) String() string {
	return fmt.Sprintf("%s: %v, Wind %s at %v, %d%% Humidity", m.location, m.temperature.String(), m.windDirection, m.windSpeed.String(), m.humidity)
}
