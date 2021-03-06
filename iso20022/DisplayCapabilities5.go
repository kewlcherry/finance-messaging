package iso20022

// Capabilities of the display components performing the transaction.
type DisplayCapabilities5 struct {

	// Destination of the message to present.
	Destination []*UserInterface5Code `xml:"Dstn"`

	// Available message format.
	AvailableFormat []*OutputFormat1Code `xml:"AvlblFrmt,omitempty"`

	// Number of lines of the display.
	NumberOfLines *Number `xml:"NbOfLines,omitempty"`

	// Number of columns of the display or printer.
	LineWidth *Number `xml:"LineWidth,omitempty"`

	// Available language for the message. Reference ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	AvailableLanguage []*LanguageCode `xml:"AvlblLang,omitempty"`
}

func (d *DisplayCapabilities5) AddDestination(value string) {
	d.Destination = append(d.Destination, (*UserInterface5Code)(&value))
}

func (d *DisplayCapabilities5) AddAvailableFormat(value string) {
	d.AvailableFormat = append(d.AvailableFormat, (*OutputFormat1Code)(&value))
}

func (d *DisplayCapabilities5) SetNumberOfLines(value string) {
	d.NumberOfLines = (*Number)(&value)
}

func (d *DisplayCapabilities5) SetLineWidth(value string) {
	d.LineWidth = (*Number)(&value)
}

func (d *DisplayCapabilities5) AddAvailableLanguage(value string) {
	d.AvailableLanguage = append(d.AvailableLanguage, (*LanguageCode)(&value))
}
