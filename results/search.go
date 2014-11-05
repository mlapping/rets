// search
package results

type Query string

type REData struct {
	MaxRows             int                   `xml:"MAXROWS"`
	ResidentialListings []ResidentialProperty `xml:"REProperties>ResidentialProperty"`
}
