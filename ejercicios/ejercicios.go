package ejercicios

import (
	"guia2/queue"
	"guia2/stack"
)

func InvertirCadena(cadena string) string {
	//TODO
	stack := stack.NewStack[string](len(cadena))
	for _, value := range cadena {
		stack.Push(string(value))
	}

	stringInvertido := ""

	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		stringInvertido += string(val)
	}
	return stringInvertido

}

func Palindromo(cadena string) bool {
	//TODO
	stack := stack.NewStack[string](len(cadena))
	queue := queue.NewQueue[string](len(cadena))
	for _, value := range cadena {
		stack.Push(string(value))
		queue.Enqueue(string(value))
	}

	esPalindromo := false

	for i := 0; !queue.IsEmpty() && !stack.IsEmpty(); i++ {
		valueQueue, _ := queue.Dequeue()
		valueStack, _ := stack.Pop()

		if valueQueue == valueStack {
			esPalindromo = true
		} else {
			esPalindromo = false
		}
	}
	return esPalindromo
}

func Balanceada(cadena string) bool {
	stack := stack.NewStack[rune](0)

	for _, valor := range cadena {
		//caso de ejemplo: [{}]
		switch valor {
		case '[', '{', '(': // entra [{
			stack.Push(valor)
		case ']', '}', ')':
			if stack.IsEmpty() {
				return false
			}
			ultimoValor, _ := stack.Pop() //Saco el ultimo valor, en este caso "{" y me lo quedo

			//comparo si el simbolo del string es el que le corresponde a la apertura que saque de la pila
			if (valor == ']' && ultimoValor != '[') || (valor == '}' && ultimoValor != '{') || (valor == ')' && ultimoValor != '(') {
				return false
			}
		}
	}
	//en el caso de que este balanceado se deberian haber eliminado todos los elementos de la pila
	return stack.IsEmpty()
}

func UnirColas(q1, q2 *queue.Queue[int]) queue.Queue[int] {
	// TODO
	qNew := queue.NewQueue[int](5)

	for !q1.IsEmpty() {
		dequeValue, _ := q1.Dequeue()
		qNew.Enqueue(dequeValue)
	}
	for !q2.IsEmpty() {
		dequeValue, _ := q2.Dequeue()
		qNew.Enqueue(dequeValue)
	}
	return *qNew
}
