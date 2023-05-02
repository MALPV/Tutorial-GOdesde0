package modelos

type Persona struct {
	Edad            int
	Altura          float32
	Peso            float32
	EstaEscribiendo bool
	EstaLeyendo     bool
	EstaDibujando   bool
}

func (p *Persona) Escribir() { p.EstaEscribiendo = true }
func (p *Persona) Leer()     { p.EstaLeyendo = true }
func (p *Persona) Dibujar()  { p.EstaDibujando = true }
