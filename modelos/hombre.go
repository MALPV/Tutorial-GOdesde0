package modelos

type Hombre struct {
	Persona
}

func (h *Hombre) Escribir() { h.EstaEscribiendo = true }
func (h *Hombre) Leer() { h.EstaLeyendo = true }
func (h *Hombre) Dibujar() { h.EstaDibujando = true }


func (h *Hombre) EstaVivo() bool { return true }
func (h *Hombre) Sexo() string { return "Masculino" }