package values

type Value int

const (
	Rock Value = iota
	Paper
	Scissors
)

// String liefert den Wert als String.
func (v Value) String() string {
	// HINWEIS:
	// Verwenden Sie einen switch-case Block um den Wert v in einen String umzuwandeln.
	//solution:begin
	switch v % 3 {
	case Rock:
		return "Stein"
	case Paper:
		return "Papier"
	case Scissors:
		return "Schere"
	default:
		return "Unknown"
	}
	//solution:end
}

// Beats gibt an, ob der Wert v den Wert w schlägt.
func (v Value) Beats(w Value) bool {
	// HINWEIS:
	// Verwenden Sie einen if-else- oder switch-Block um zu prüfen, ob der Wert v den Wert w schlägt.
	// Alternativ können Sie auch die Werte in Integer umwandeln und rechnerisch prüfen.
	// solution:begin
	vi := int(v)
	wi := int(w)

	return vi == (wi+1)%3
	// solution:end
}
