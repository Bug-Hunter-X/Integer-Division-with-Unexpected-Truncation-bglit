# Go Integer Division Bug

This repository demonstrates a common error in Go related to integer division and its unexpected behavior with negative numbers and zero values.

The `bug.go` file contains a function `MyFunc` that performs integer division. The function is flawed as it doesn't handle edge cases like division by zero, and negative number division.

The `bugSolution.go` file presents a corrected version that addresses these issues and ensures accurate and robust division.

## How to Run

1. Clone the repository.
2. Navigate to the repository directory.
3. Run the `bug.go` and `bugSolution.go` files using the `go run` command.