package main

func main() {
	action := &action{}

	eval := &evaluator{}
	eval.setNext(action)

	collect := &collector{}
	collect.setNext(eval)

	collect.execute(&order{
		id: 1,
	})
}
