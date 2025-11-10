package player

import (
	"fmt"
	"rockpaperscissors/values"
)

type Human struct {
	Name  string
	Score int
}

// NewHuman erzeugt einen neuen Spieler mit dem Namen n.
func NewHuman(n string) *Human {
	return &Human{Name: n, Score: 0}
}

// GetName liefert den Namen des Spielers als String.
func (p Human) GetName() string {
	// HINWEIS:
	// Greifen Sie mit p.Name auf den Namen des Spielers zu.
	// solution:begin
	return p.Name
	// solution:end
}

// GetScore liefert den Punktestand des Spielers.
func (p Human) GetScore() int {
	// HINWEIS:
	// Greifen Sie mit p.Score auf den Punktestand des Spielers zu.
	// solution:begin
	return p.Score
	// solution:end
}

// IncrementScore erhöht den Punktestand des Spielers um eins.
func (p *Human) IncrementScore() {
	// HINWEIS:
	// Greifen Sie mit p.Score auf den Punktestand des Spielers zu.
	// Sie können p.Score wie eine reguläre Variable verwenden.
	// solution:begin
	p.Score++
	// solution:end
}

// GetMove liefert einen Zug des Spielers.
func (p Human) GetMove() values.Value {
	// HINWEIS:
	// Fordern Sie den Spieler mit fmt.Printf() auf, einen Zug zu wählen.
	// Geben Sie dem Spieler die Optionen "Stein", "Papier" und "Schere".
	// Die Optionen können z.B. mit Zahlen 1, 2 und 3 verknüpft werden.
	// Verwenden Sie fmt.Scanln() um die Eingabe des Spielers zu lesen.
	// solution:begin
	fmt.Printf("%s, was wählst du?\n", p.GetName())
	fmt.Println("  1: Stein")
	fmt.Println("  2: Papier")
	fmt.Println("  3: Schere")
	var v int
	fmt.Scanln(&v)
	return values.Value((v - 1) % 3)
	// solution:end
}
