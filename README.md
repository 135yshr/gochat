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

## Lesson 3

[net](https://golang.org/pkg/net/)を使って、ネットワーク通信を実装しましょう

1. `fmt.Println("Hello", name)` の次の行に `listen, _ := net.Listen("tcp", "127.0.0.1:8888")` と書きます
2. `listen, _ := net.Listen("tcp", "127.0.0.1:8888")` の次の行に `fmt.Println("Listen 127.0.0.1:8888")` と書きます
3. `fmt.Println("Listen 127.0.0.1:8888")` の次の行に `conn, _ := listen.Accept()` と書きます
4. `go run cmd/main.go` を実行します
5. `conn, _ := listen.Accept()` の次の行に `buf := make([]byte, 1024)` と書きます
6. `buf := make([]byte, 1024)` の次の行に `n, _ := conn.Read(buf)` と書きます
7. `n, _ := conn.Read(buf)` の次の行に `fmt.Printf("[Message]\n%s", string(buf[:n]))` と書きます
8. 最後の行に `conn.Close()`

## Lesson 4

エラー処理を追加しましょう

1. `listen, _ := net.Listen("tcp", "127.0.0.1:8888")` を `listen, err := net.Listen("tcp", "127.0.0.1:8888")` に書き換えます
2. `listen, err := net.Listen("tcp", "127.0.0.1:8888")` の次の行にエラー処理を追加します
3. `conn, _ := listen.Accept()` を `conn, err := listen.Accept()` に書き換えます
4. `conn, err := listen.Accept()` の次の行にエラー処理を追加します
5. `n, _ := conn.Read(buf)` を `n, err := conn.Read(buf)` に書き換えます
6. `n, err := conn.Read(buf)` の次の行にエラー処理を追加します
