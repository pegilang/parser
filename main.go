package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/pegilang/parser/Go"
	"io/ioutil"
	"log"
	"os"
)
type pegiListener struct {
	*parser.BasePegiParserListener
}
func main () {
	path := os.Args[1]
	log.Println("parse", path)
	file, err :=ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	lexer := parser.NewPegiLexer(antlr.NewInputStream(string(file)))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPegiParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(&pegiListener{}, p.SourceFile())
}

func (s *pegiListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}