antlr4 -o ./Go -listener -visitor -Dlanguage=Go -lib ./ ./PegiParser.g4 ./PegiLexer.g4
antlr4 -o ./Java -listener -visitor -Dlanguage=Java -lib ./ ./PegiParser.g4 ./PegiLexer.g4
