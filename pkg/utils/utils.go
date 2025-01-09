package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

/*
***************************************************************************************************************************
r *http.Request: Il parametro r rappresenta una richiesta HTTP. r.Body contiene il corpo della richiesta che vogliamo leggere.
x interface{}: Il parametro x è un'interfaccia generica (vuota), che consente di passare qualsiasi tipo.
L'idea è che x sarà una struttura (o un altro tipo) in cui verranno deserializzati i dati JSON.
****************************************************************************************************************************
*/
func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}

	}
	/*Se tutto funziona correttamente, la funzione popola la variabile x con i dati deserializzati. Non restituisce nulla, ma modifica direttamente
	  il valore puntato da x (poiché gli oggetti in Go sono passati per riferimento quando sono puntatori o strutture contenute in un'interfaccia).*/

}
