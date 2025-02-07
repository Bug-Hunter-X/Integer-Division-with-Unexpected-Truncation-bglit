func MyFunc(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
} 
func handleError(err error) {
    if err != nil {
        fmt.Println("Error:", err)
    }
}