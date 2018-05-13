package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func generateAst(outputDir, baseName string, defs []string) {

	for _, def := range defs {
		// a single definition is in the form of:
		//
		// ClassName : mname1 MType1, mname2 MType2, ...
		tokens := strings.Split(def, ":")
		className := strings.TrimSpace(tokens[0])
		membersDef := tokens[1]

		members := strings.Split(membersDef, ",")

		f, err := os.Create(outputDir + string(os.PathSeparator) + toSnake(className) + ".go")
		if err != nil {
			panic(err)
		}

		defer f.Close()

		buf := bufio.NewWriter(f)

		_, err = buf.WriteString("package main\n\n")
		if err != nil {
			panic(err)
		}

		_, err = buf.WriteString(fmt.Sprintf("type %s struct {\n", className))
		if err != nil {
			panic(err)
		}
		for _, member := range members {
			_, err = buf.WriteString(fmt.Sprintf("    %s\n", strings.TrimSpace(member)))
			if err != nil {
				panic(err)
			}
		}

		_, err = buf.WriteString("}\n\n")

		// now let's implement the constructor
		_, err = buf.WriteString(fmt.Sprintf("func New%s(%s) *%s {\n", className, membersDef, className))
		if err != nil {
			panic(err)
		}

		_, err = buf.WriteString(fmt.Sprintf("    return &%s{\n", className))
		if err != nil {
			panic(err)
		}

		// now let's pass all the required parameters
		for _, member := range members {
			parts := strings.Split(strings.TrimSpace(member), " ")
			_, err = buf.WriteString(fmt.Sprintf("        %s,\n", strings.TrimSpace(parts[0])))
		}
		_, err = buf.WriteString("    }\n}\n")
		if err != nil {
			panic(err)
		}
		buf.Flush()
	}
}

// freely copied from https://gist.github.com/elwinar/14e1e897fdbe4d3432e1
func toSnake(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "Expected output directory argument to be provided.\n")
		os.Exit(-1)
	}

	generateAst(os.Args[1], "Expression", []string{
		"Binary   : left *Expression, operator Token, right *Expression",
		"Grouping : expression *Expression",
		"Literal  : value interface{}",
		"Unary    : operator Token, right *Expression",
	})
}
