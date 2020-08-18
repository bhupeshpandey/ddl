package models

import "time"


type Configurations struct {
	Server       ServerConfigurations
	Database     DatabaseConfigurations
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBName     string `json:"dbName"`
	DBUser     string `json:"dbuserName"`
	DBPassword string `json:"dbpassword"`
}

type Patient struct {
	Id int `json:"Id"`
	First string `json:"firstName"`
	Last string `json:"lastName"`
	Middle string `json:"Middle"`
	Suffix string `json:"suffix"`
	SSN string `json:"ssn"`
	DOB time.Time `json:"DOB"`
	Gender string `json:"gender"`
	Age int `json:"age"`
	Visit string `json:"visit"`
	Alias string `json:"alias"`
	Action string `json:"action"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
	CreatedBy string `json:"createdBy"`
	UpdatedBy string `json:"updatedBy"`
}