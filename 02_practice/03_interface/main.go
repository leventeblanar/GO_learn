package main


import "fmt"

// Interface -> azt mondja: bármi Notifier, aminek van Notify string metódusa
type Notifier interface {
	Notify() string
}

type Email struct {
	Address string
}

type SMS struct {
	PhoneNumber string
}


// Az alábbi functionök adott structhoz tartoznak és Notify string metódusok
// Így alapvetően teljesítik a Notifier interfacet
func (e Email) Notify() string {
	return "Sending email to " + e.Address 
}

func (s SMS) Notify() string {
	return "Sending SMS to " + s.PhoneNumber
}


func SendNotification(n Notifier) {
	fmt.Println(n.Notify())
}

func main() {

	// létrehozunk egy Notifier slice-ot, mivel mindegyiknek van Notify metódusa, így beleférnek egybe
	notifications := []Notifier {
		Email{Address: "test@example.com"},
		SMS{PhoneNumber: "+36701234567"},
		Email{Address: "admin@example.com"},
	}

	for _, notification := range notifications {
		SendNotification(notification)
	}

}
