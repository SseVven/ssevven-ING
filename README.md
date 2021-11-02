# ssevven-ING

1. `go env -w GOPROXY=https://goproxy.io,direct`

   创建 go.mod 文件

2. 加载所有的 html 模板

   ```go
   func loadTemplates() *template.Template {
       result := template.New("templates")
       template.Must(result.ParseGlob("templates/*.html"))
       return result
   }
   ```

3. https

   `go run GIROOT\crypto\tls\generate_cert.go -h`
