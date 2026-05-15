package main

import (
	"fmt"

)

type Estado struct {
    texto string
    cursor int
}
type Editor struct {
	texto  string
	cursor int
    undo   []Estado
	redo   []Estado
}

func(e *Editor) SalvarEstado() {
    e.undo = append(e.undo, Estado{e.texto,e.cursor})
    e.redo = nil
}

func(e *Editor) Desfazer(){
    if len(e.undo) == 0 {
        return
    }

    e.redo = append(e.redo, Estado{e.texto,e.cursor})
    ultimo := e.undo[len(e.undo)-1]
    e.texto = ultimo.texto
    e.cursor = ultimo.cursor
    e.undo = e.undo[:len(e.undo)-1]
}


func(e *Editor) Refazer() {
    if len(e.redo) == 0{
        return
    }

    e.undo = append(e.undo, Estado{e.texto,e.cursor})
    ultimo := e.redo[len(e.redo)-1]
    e.texto = ultimo.texto
    e.cursor = ultimo.cursor
    e.redo = e.redo[:len(e.redo)-1]
}

func (e *Editor) Right() {
	if e.cursor < len(e.texto){

        e.cursor++
    }
}

func (e *Editor) Left() {
	if e.cursor > 0 {
        
        e.cursor--
    }
}

func (e *Editor) Enter() {
    e.SalvarEstado()
	antes := e.texto[:e.cursor]
	depois := e.texto[e.cursor:]

	e.texto = antes + "\n" + depois
    e.cursor++
}

/////

func (e *Editor) Backspace() {
    e.SalvarEstado()
	if e.cursor > 0 {
		antes := e.texto[:e.cursor-1]
		depois := e.texto[e.cursor:]

		e.texto = antes + depois
		e.cursor--
	}
}

func (e *Editor) delete() {
    e.SalvarEstado()
    if e.cursor < len(e.texto) {
        antes := e.texto[:e.cursor]
        depois := e.texto[e.cursor+1:]

        e.texto = antes + depois
      
    }
   
}


func(e *Editor) Show() string {
    antes := e.texto[:e.cursor]
	depois := e.texto[e.cursor:]

    return antes + "|" + depois
}

func (e *Editor) Inserir(c rune){
    e.SalvarEstado()
     antes := e.texto[:e.cursor]
    depois := e.texto[e.cursor:]

    e.texto = antes + string(c) + depois
    e.cursor++
}

func main() {
    entrada := ""
    fmt.Scan(&entrada)

    editor := Editor{
        texto: "",
        cursor: 0,
        undo:   []Estado{},
		redo:   []Estado{},
    }


    for _, c := range entrada{
       
        switch c {
        case  '>':
            editor.Right()
        case '<':
            editor.Left()
        case 'D':
            editor.delete()
        case 'B':
            editor.Backspace()
        case 'R':
            editor.Enter()
        case 'Z':
            editor.Desfazer()
        case 'Y':
            editor.Refazer()
            default :
            editor.Inserir(c)
        }
    }
    fmt.Println(editor.Show())
}
