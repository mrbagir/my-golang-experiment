package model

type App struct {
	Name     string    `yaml:"name" json:"name"`
	Settings []Setting `json:"settings"`
	Terminal *Terminal `json:"terminal"`
	Active   bool      `json:"active"`
	Hide     bool      `json:"-"`
}

type Setting struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
	Value  any    `json:"value"`
}

type Terminal struct {
	Log        []string `json:"log"`
	MaxLines   int      `json:"maxLines"`
	AutoScroll bool     `json:"autoScroll"`
}

type Config struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
