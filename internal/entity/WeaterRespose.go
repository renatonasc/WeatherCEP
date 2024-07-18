package entity

type WeaterRespose struct {
	Location              string  `json:"location"`
	Temparatue_celcius    float32 `json:"temparatue_celcius"`
	Temparatue_fahrenheit float32 `json:"temparatue_fahrenheit"`
	Temperature_kelvin    float32 `json:"temperature_kelvin"`
}
