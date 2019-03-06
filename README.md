# gochat

## Lesson 1

Go を使って画面にメッセージを表示しましょう

1. cmd/main.go ファイルを作成します
2. `package main` と `func main()` を書きます
3. main 内に `fmt.Println("Hello")` と書きます
4. `go run cmd/main.go` を実行します

## Lesson 2

標準入力を受け取って、受け取った内容を画面に表示しましょう

1. `fmt.Println("Hello")` の前の行に `name := "My Name"` と書きます
2. `fmt.Println("Hello")` を `fmt.Println("Hello", name)` と書き換えます
3. `go run cmd/main.go` を実行します
4. `name := "My Name"` の前の行に `stdin := bufio.NewScanner(os.Stdin)` と書きます
5. `stdin := bufio.NewScanner(os.Stdin)` の次の行に `stdin.Scan()` と書きます
6. `go run cmd/main.go` を実行します
7. `name := "My Name"` を `name := stdin.Text()` と書き換えます
8. `go run cmd/main.go` を実行します
9. `func main()` の次の行に `fmt.Print("Hi. What are you name? ")` と書きます
10. `go run cmd/main.go` を実行します

**`name := "My Name"`は、`var name string; name = "My Name"` や `var name string = "My Name"`と書くこともできます**
