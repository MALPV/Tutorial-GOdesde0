package modelos

type Mujer struct {
	Persona
}

func (m *Mujer) Escribir() { m.EstaEscribiendo = true }
func (m *Mujer) Leer()     { m.EstaLeyendo = true }
func (m *Mujer) Dibujar()  { m.EstaDibujando = true }

func (m *Mujer) EstaVivo() bool { return true }
func (m *Mujer) Sexo() string   { return "Femenino" }
