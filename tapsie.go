package tapsie

import "time"

// breweries have an address
type Address struct {
	Street string
	City   string
	State  string
	Zip    int
}

// hours of operation for an entity (food truck, brewery)
type Hours struct {
	Start time.Time
	End   time.Time
}

// hours of operation by day of week
type Open struct {
	Sunday    Hours
	Monday    Hours
	Tuesday   Hours
	Wednesday Hours
	Thursday  Hours
	Friday    Hours
	Saturday  Hours
}

// email and phone
type Contact struct {
	Email string
	Phone int
}

// amenities for a brewery, does it have x ?
type Amenities struct {
	Pets   bool
	Kids   bool
	Wifi   bool
	Food   bool
	Events bool
}

// a brewery
type Brewery struct {
	ID        int
	Name      string
	Address   Address
	Open      Open
	Contact   Contact
	Amenities Amenities
}

// a beer
// a few of these could be arrays
// for example, a beer might belong to
// multiple styles or have multiple flavors
// right now, keeping at one per attribute
// for simplicities sake
type Beer struct {
	ID           int
	ABV          float64
	IBU          int
	Color        string
	Name         string
	Fermentation string
	Origin       string
	Style        string
	Era          string
	Flavor       string
	Appearance   string
	Mouthfeel    string
}
