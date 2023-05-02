package defer_panic

import "log"

func PanicExample(){

	defer func ()  {
		rec := recover()
		if rec != nil {
			log.Fatalf("Ocurrio un error que genero Panic\n%v\n", rec)
		}
	}()

	input := 7
	if input == 7 {
		panic("Se ha encontrado el valor 7")
	}

}
