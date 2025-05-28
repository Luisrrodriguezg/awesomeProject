package main

import (
	"fmt"

	"awesomeProject/evaluator"
	"awesomeProject/lexer"
	"awesomeProject/object"
	"awesomeProject/parser"
)

func run(src string, env *object.Environment) (object.Object, []string) {
	l := lexer.New(src)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		return nil, p.Errors()
	}

	return evaluator.Eval(program, env), nil
}

func main() {
	snippets := []string{
		// Fibonacci recursivo (n = 20)
		`let fibo = fn(n) {
            if (n == 1) { 1 } else {
                if (n == 2) { 1 } else { fibo(n-1) + fibo(n-2) }
            }
        };
        fibo(20);`,

		// dobleSuma usando recursión (evita while)
		`let suma = fn(a, b) { a + b };
        let dobleSuma = fn(x, y, n) {
            if (n == 0) { 0 } else { suma(x, y) * 2 + dobleSuma(x, y, n-1) }
        };
        dobleSuma(3, 7, 7);`,

		// factorial recursivo
		`let factorial = fn(x) {
            if (x == 0) { 1 } else { x * factorial(x-1) }
        };
        factorial(5);`,

		// 45 negaciones
		`!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!false;`,

		// Ejemplo con error de sintaxis
		`let errorFn = fn(a b) { a + b }; errorFn(1);`,
	}

	env := object.NewEnvironment()

	for i, code := range snippets {
		result, errs := run(code, env)
		if len(errs) > 0 {
			fmt.Printf("Errores en snippet #%d:\n", i+1)
			for _, e := range errs {
				fmt.Println("\t", e)
			}
			continue
		}
		fmt.Printf("Resultado #%d → %s\n", i+1, result.Inspect())
	}
}
