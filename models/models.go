package models

import "gorm.io/gorm"

type Port struct {
	gorm.Model
	Port      int
	Receiveds []Received `gorm:"many2many:reciber_ports;"`
}

type Received struct {
	gorm.Model
	Ip       string
	Hostname string
	Os       string
	Ports    []Port `gorm:"many2many:reciber_ports;"`
}
