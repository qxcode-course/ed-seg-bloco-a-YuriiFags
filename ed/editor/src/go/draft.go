package main

import (
	"fmt"
)


type Editor struct {
    texto string
    cursor int
}


func (e *Editor) Inserir(c rune){
    antes := e.texto[:e.cursor]
    depois := e.texto[e.cursor:]

    e.texto = antes + string(c) + depois
    e.cursor++
}

func (e *Editor) Enter(){
    antes := e.texto[:e.cursor]
    depois := e.texto[e.cursor:]

    e.texto = antes + "\n" + depois
    e.cursor++
}

func (e *Editor) Backspace(){
    if e.cursor > 0 {
        antes := e.texto[:e.cursor -1]
        depois := e.texto[e.cursor:]

        e.texto = antes + depois
        e.cursor--
    }
}

func (e *Editor) Delete() {
    if e.cursor < len(e.texto) {
        antes := e.texto[:e.cursor]
        depois := e.texto[e.cursor+1:]

        e.texto = antes + depois
    }
}

func (e *Editor) Direita(){
    if e.cursor < len(e.texto){
        e.cursor++
    }
}

func (e *Editor) Esquerda() {
    if e.cursor > 0 {
        e.cursor--
    }
}
func (e *Editor) Show() string{
    antes := e.texto[:e.cursor]
    depois := e.texto[e.cursor:]

   return  antes + "|" + depois
}

func main() {
   entrada :=""
   fmt.Scan(&entrada)
   
   editor := Editor{
    texto: "",
    cursor: 0,
   }

   for _, c := range entrada{

        switch c {
        case '>':
            editor.Direita()
        case '<':
            editor.Esquerda()

        case 'D':
            editor.Delete()

        case 'B':
            editor.Backspace()

        case 'R':
            editor.Enter()

        default:
            editor.Inserir(c)
        }
   }
   fmt.Println(editor.Show())
}
