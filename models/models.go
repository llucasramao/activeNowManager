package models

import "gorm.io/gorm"

type Port struct {
	gorm.Model
	Port      int
	Receiveds []Received `gorm:"many2many:reciber_ports;"`
}

type App struct {
	gorm.Model
	Name      string
	Version   string
	Receiveds []Received `gorm:"many2many:reciber_apps;"`
}

type Received struct {
	gorm.Model
	Ip       string
	Hostname string
	Os       string
	Ports    []Port `gorm:"many2many:reciber_ports;"`
	Apps     []App  `gorm:"many2many:reciber_apps;"`
}
