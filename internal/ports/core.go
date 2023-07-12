package ports

type ArithmeticPort interface {
	Addition(a, b int32) (int32, error)
	Subtraction(a, b int32) (int32, error)
	Multiplication(a, b int32) (int32, error)
	Divisions(a, b int32) (int32, error)
}
