package services

// Common Ports -> Service Names
var Services = map[int]string{
	22:   "SSH",
	80:   "HTTP",
	443:  "HTTPS",
	3306: "MySQL",
	5432: "PostgreSQL",
}
