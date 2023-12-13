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

type Service struct {
	gorm.Model
	Name     string
	Received []Received `gorm:"many2many:reciber_services;"`
}

type Received struct {
	gorm.Model
	Ip           string
	Hostname     string
	Os           string
	Ports        []Port    `gorm:"many2many:reciber_ports;"`
	Apps         []App     `gorm:"many2many:reciber_apps;"`
	Services     []Service `gorm:"many2many:reciber_services;"`
	AgentVersion string
}
