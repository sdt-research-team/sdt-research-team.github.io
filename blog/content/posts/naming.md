---
title: "Naming"
date: 2020-12-19T09:57:01+07:00
draft: false
tags: 
  - go
---
# Naming

### Table of content

[Package](#package)\
[Variable](#variable)\
[Interface](#interface)\
[Semicolons](#semicolons)


### Package
Packages được sử dụng để tổ chức code sao cho việc đọc và tái sử dụng code dễ dàng hơn. Packages giúp phân chia code thành nhiều phần, do đó việc update ứng dụng cũng thuận tiện hơn.

Tên một package nên được đặt một cách ngắn gọn và rõ nghĩa. Theo quy định chung của ngôn ngữ Golang thì nên dùng chuẩn Lower-case để đặt tên khi viết code trong Golang (bao gồm package, biến, hàm, ...)
Lấy ví dụ những package trong thư viện golang thường là những danh từ đơn giản như:
    - time (hỗ trợ các thao tác liên quan đến xử lý thời gian)
    - list (dùng để dựng chuỗi liên kết đơn hoặc đôi)

Đánh dấu `var`s và `const`s bằng token `_` để thể hiện đây là các biến được dùng global trong code.

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
// foo.go

const (
  defaultPort = 8080
  defaultUser = "user"
)

// bar.go

func Bar() {
  defaultPort := 9090
  ...
  fmt.Println("Default port", defaultPort)

  // We will not see a compile error if the first line of
  // Bar() is deleted.
}
```

</td><td>

```go
// foo.go

const (
  _defaultPort = 8080
  _defaultUser = "user"
)
```

</td></tr>
</tbody></table>

### Variable Naming Convention
Tên biến trong Golang được viết theo chuẩn CamelCase, tuy nhiên đối với một vài từ khóa đặc biệt thì nên đặt tên theo các mẫu sau:

Đây là danh sách những từ khóa nên được viết theo mẫu đặc biệt

```go
// A GonicMapper that contains a list of common initialisms taken from golang/lint
var LintGonicMapper = GonicMapper{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XSRF":  true,
	"XSS":   true,
}
```

Đối với kiểu dữ liệu `bool` thì tên biến nên được bắt đầu bằng những từ khóa `Has`, `Is`, `Can` hoặc `Allow`, etc.

Lấy ví dụ định nghĩa một struct theo quy tắc đặt tên biến

```go
	// Webhook represents a web hook object.
	type Webhook struct {
		ID           int64 `xorm:"pk autoincr"`
		RepoID       int64
		OrgID        int64
		URL          string `xorm:"url TEXT"`
		ContentType  HookContentType
		Secret       string `xorm:"TEXT"`
		Events       string `xorm:"TEXT"`
		*HookEvent   `xorm:"-"`
		IsSSL        bool `xorm:"is_ssl"`
		IsActive     bool
		HookTaskType HookTaskType
		Meta         string     `xorm:"TEXT"` // store hook-specific attributes
		LastStatus   HookStatus // Last delivery status
		Created      time.Time  `xorm:"CREATED"`
		Updated      time.Time  `xorm:"UPDATED"`
	}
```

### Functions and Methods

Tuân thủ theo nguyên tắc CamelCase, tuy nhiên nếu hàm trả về kiểu `bool` thì tên hàm nên được bắt đầu với những từ khóa `Has`, `Is`, `Can` hoặc `Allow`, etc.

```go
	func HasPrefix(name string, prefixes []string) bool { ... }
	func IsEntry(name string, entries []string) bool { ... }
	func CanManage(name string) bool { ... }
	func AllowGitHook() bool { ... }
	```

### Constants
Constants nên được đặt tất cả các ký tự bằng chữ viết hoa, và nên dùng ký tự `_` để phân biệt các từ.

```go
	const APP_VER = "0.7.0.1110 Beta"
```

### Interface
Trong Golang, một Interface sẽ được đặt tên theo 3 cách
- Khai báo Interface bằng tên phương thức muốn dùng kết hợp với "er" ở sau cùng.
- Khai báo Interface bằng tên phương thức muốn dùng kết hợp với prefix "I".
- Khai báo Interface bằng tên phương thức muốn dùng kết hợp với "er", nhưng tùy biến với từng trường hợp. Ví dụ: có thể dùng `Reader` làm tên Interface mà không cần kết hợp thêm er ở sau cùng.

### Semicolons
Tương tự như ngôn ngữ C, Golang dùng ký tự `;` để kết thúc một dòng code, tuy nhiên điểm khác biệt là ký tự này không bắt buộc phải xuất hiện lúc viết code. Thay vào đó bộ compile của ngôn ngữ Go sẽ tự động thêm ký tự này vào dòng code.

Nguyên tắc hoạt động của bộ compile như sau: nếu ký tự cuối cùng của dòng code là biến, là con số, ... hoặc là một trong những ký tự sau đây.

```shell
  break continue fallthrough return ++ -- ) }
```

Một trong những nguyên tắc là không được thêm ký tự `;` vào sau những từ khóa của câu điều kiện. Cấu trúc điều kiện nên được viết như thế này:

```shell
  if i < f() {
    g()
  }
```

Không nên như thế này

```shell
  if i < f()  // wrong!
  {           // wrong!
      g()
  }
```


*<p style="text-align: end;">- Pham Quoc Dat -</p>*