package main

import "fmt"

type NodoLista struct {
	sig   *NodoLista
	valor int
}
type Lista struct {
	inicio   *NodoLista
	longitud int
}

// crear una lista vacia
func ListaVacia() *Lista {
	l := new(Lista)
	l.inicio = nil
	l.longitud = 0
	return l
}

// crear nodo y añadir elemento
func fNodoLista(elem int) *NodoLista {
	nl := new(NodoLista)
	nl.sig = nil
	nl.valor = elem
	return nl
}

//verificar si es una lista vacia
func (l *Lista) EsListaVacia() bool {
	return l.inicio == nil
}

//agregar elemento a la lista (inserta por la izquierda)
func (l *Lista) agregarNodo(elem int) {
	nl := fNodoLista(elem)
	if l.EsListaVacia() {
		l.inicio = nl
	} else {
		nl.sig = l.inicio
		l.inicio = nl
	}
	l.longitud += 1
}

func main() {
	l1 := ListaVacia()
	if !l1.EsListaVacia() {
		panic("error de lista vacia\n")
	}
	l1.agregarNodo(2)
	fmt.Printf("L1 tiene el elemento: %v\n", l1.inicio.valor)
	if l1.inicio.valor != 2 {
		panic("error de valor del primer elemento, primer elemento es 2\n")
	}
	fmt.Printf("All test are ok\n")
}
