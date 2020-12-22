---
title: "Pointer Definition"
date: 2020-12-18T09:57:01+07:00
author: "Pham Minh Toan"
draft: false
tags: 
  - go
---
# Pointers

## Variable
Variable (hay còn gọi là biến) là một ô nhớ đơn lẻ hoặc một vùng nhớ được hệ điều hành cấp phát cho chương trình Go nhằm để lưu trữ giá trị vào bên trong vùng nhớ đó. Để truy xuất đến giá trị mà biến đang nắm giữ, chương trình cần tìm đến vùng nhớ (địa chỉ) của biến để đọc giá trị bên trong vùng nhớ đó, cũng như bạn muốn lấy món đồ bên trong cái hộp, bạn cần biết cái hộp được đặt ở đâu.
Khi thao tác với các biến thông thường, chúng ta không cần quan tâm đến địa chỉ vùng nhớ của biến. Khi cần truy xuất giá trị của biến, chúng ta chỉ cần gọi định danh (hay thường gọi là tên biến).

Ví dụ:
int16 money;
Khi dòng lệnh này được CPU thực thi, một vùng nhớ có kích thước 2 bytes sẽ được cấp phát. Lấy ví dụ biến money này được đặt tại ô nhớ 1224 (trong địa chỉ ảo của máy tính).
![Image 1](https://raw.githubusercontent.com/sendo-research-team/sendo-research-team.github.io/gh-pages/static/images/1.JPG)
￼
Bất cứ khi nào chương trình thấy các bạn sử dụng biến money trong câu lệnh, chương trình hiểu rằng cần tìm đến ô nhớ 1224 để lấy giá trị đó ra.

## Virtual memory & Physical memory
Việc truy xuất dữ liệu trên bộ nhớ máy tính cần phải thông qua một số bước trung gian, người dùng không thể trực tiếp truy xuất vào các ô nhớ trên các thiết bị lưu trữ. Chúng ta chỉ có thể trỏ đến vùng nhớ ảo (virtual memory) trên máy tính, còn việc truy xuất đến bộ nhớ vật lý (physical memory) từ bộ nhớ ảo phải được thực hiện bởi thiết bị phần cứng có tên là Memory management unit (MMU) và một chương trình định vị địa chỉ bộ nhớ gọi là Virtual address space.
![Image 2](https://raw.githubusercontent.com/sendo-research-team/sendo-research-team.github.io/gh-pages/static/images/2.JPG)
￼

Virtual memory làm che giấu sự phân mảnh của bộ nhớ vật lý, khiến chúng ta có cảm giác đang thao tác với các vùng nhớ liên tục. Trong hình trên, từ phía Virtual memory cho đến Physical memory thuộc về phần quản lý của hệ điều hành, lập trình viên và người dùng chúng ta không thể can thiệp trực tiếp đến trong quá trình máy tính đang hoạt động.

## Variable address & address-of operator
Địa chỉ của biến mà chúng ta nhìn thấy thật ra chỉ là những giá trị đã được đánh số thứ tự đặt trên Virtual memory. Để lấy được địa chỉ ảo của biến trong chương trình, chúng ta sử dụng toán tử '&' đặt trước tên biến.
```go
int x = 5;
fmt.Println(x) // print the value of variable x
fmt.Println(&x) / print the memory address of variable x
```
Trên máy tính của mình, kết quả của đoạn chương trình trên được in ra như sau:
```go
5
0027FEA0
```
Dòng đầu tiên là kết quả của việc truy xuất giá trị của biến thông qua định danh (tên biến). Dòng thứ hai là kết quả của việc truy xuất đến địa chỉ ảo của biến.

Dereference operator
Toán tử trỏ đến (dereference operator) hay còn gọi là indirection operator (toán tử điều hành gián tiếp) được kí hiệu bằng dấu sao " * " cho phép chúng ta lấy ra giá trị của vùng nhớ có địa chỉ cụ thể.
Ví dụ:
```go
int n = 5;

fmt.Println(n)   //print the value of variable n
fmt.Println(&n)   //print the virtual memory address of variable n
fmt.Println(*(&n)) //print the value at the virtual memory address of variable n
```
* Dòng lệnh đầu tiên khá dễ hiểu, nó thực hiện in ra giá trị của biến n bằng cách gọi định danh n, còn lại phần truy xuất đến địa chỉ ảo của biến n sẽ do chương trình đảm nhiệm.
* Dòng lệnh thứ hai không dùng để lấy ra giá trị bên trong vùng nhớ mà biến n đang nắm giữ, mà nó lấy ra địa chỉ ảo của biến n.
* Dòng lệnh thứ ba chúng ta sử dụng toán tử trỏ đến " * " đặt trước toán tử address-of. Khi đó, (&n) sẽ lấy ra địa chỉ ảo của biến n, và toán tử * sẽ truy xuất giá trị bên trong địa chỉ đó.
Kết quả của đoạn chương trình trên là:
```go
5
0xBFD181AC
5
```
Ngoài việc truy xuất giá trị trong vùng nhớ của một địa chỉ cụ thể, toán tử trỏ đến (dereference operator) còn có thể dùng để thay đổi giá trị bên trong vùng nhớ đó.
```go
int n = 5;
fmt.Println(n);
*n = 10;
fmt.Println(n);
```
Kết quả đoạn chương trình này là:
```go
5
10
```
Như vậy, dereference operator cho phép chúng ta thao tác trực tiếp trên Virtual memory mà không cần thông qua định danh (tên biến).
![Image 3](https://raw.githubusercontent.com/sendo-research-team/sendo-research-team.github.io/gh-pages/static/images/3.JPG)

Mặc dù dereference operator có kí hiệu giống multiplication operator, nhưng các bạn có thể phân biệt được vì dereference operator là toán tử một ngôi, trong khi đó, multiplication operator là toán tử hai ngôi.
Khác với tham chiếu (reference), toán tử trỏ đến (dereference operator) không tạo ra một tên biến khác, mà nó truy xuất trực tiếp đến vùng nhớ có địa chỉ cụ thể trên Virtual memory.

## Con trỏ (Pointer)
Với những khái niệm mình trình bày ở trên (một số khái niệm các bạn đã được học), bây giờ chúng ta có thể nói đến con trỏ (pointer).
Một con trỏ (a pointer) là một biến được dùng để lưu trữ địa chỉ của biến khác.
Khác với tham chiếu, con trỏ là một biến có địa chỉ độc lập so với vùng nhớ mà nó trỏ đến, nhưng giá trị bên trong vùng nhớ của con trỏ chính là địa chỉ của biến (hoặc địa chỉ ảo) mà nó trỏ tới.
![Image 4](https://raw.githubusercontent.com/sendo-research-team/sendo-research-team.github.io/gh-pages/static/images/4.JPG)
￼

Trong ví dụ trên, một con trỏ sau khi khai báo đã được cấp phát vùng nhớ tại địa chỉ 3255, và nó trỏ đến địa chỉ 1224, do đó, giá trị bên trong vùng nhớ của con trỏ là 1224.

## Khai báo con trỏ
Cũng giống như biến thông thường, biến con trỏ cần được khai báo trước khi sử dụng. Con trỏ yêu cầu cú pháp khai báo mới hơn một chút so với biến thông thường.
var <name_of_pointer> *<data_type>;
Khác với biến thông thường, chúng ta cần đặt thêm dấu sao giữa tên biến và kiểu dữ liệu của con trỏ.
Ví dụ:
```go
var p *int
var pt *string
```
Lưu ý: Dấu sao trong khai báo con trỏ không phải là toán tử trỏ đến (dereference operator), nó chỉ là cú pháp được ngôn ngữ Golang quy định.

* Kiểu dữ liệu của con trỏ không mô tả giá trị địa chỉ được lưu trữ bên trong con trỏ, mà kiểu dữ liệu của con trỏ dùng để xác định kiểu dữ liệu của biến mà nó trỏ đến trên bộ nhớ ảo.

Hãy xem đoạn code sau:   
```go
func main() {
	// declare an int value and an int pointer
	var ival int = 1

	var iptr *int = &ival

	// declare a float value and a float pointer
	var fval float32 = 1.0

	var fptr *float32 = &fval

	// declare a char value and a char pointer
	var sval string = ""

	var sptr *string = &sval

	// can't do this, doesn't make sense
	// iptr = &fval;
	// fptr = &ival;
	// iptr = &sval;
}
```
    
Khi chúng ta định nghĩa một pointer kiểu int, chúng ta định nghĩa biến đó là 1 pointer, biến đó giữ địa chỉ tới một biến khác, và giá trị ở tại địa chỉ đó là 1 số nguyên int. Tương tự đối với float pointer, char pointer, hay bất cứ kiểu nào khác. Định nghĩa một pointer thuộc một kiểu xác định sẽ giúp cho trình biên dịch biết rằng khi chúng ta tham chiếu ngược tới một pointer đó thì nó sẽ trỏ đến giá trị thuộc kiểu nào.
Bạn sẽ thấy rằng trong ví dụ trên, chúng ta định nghĩa ra pointer thuộc một kiểu nào đó và gán địa chỉ của một giá trị thuộc cùng kiểu. Nếu bạn bỏ comment mấy dòng cuối và thử compile nó thì sẽ bị lỗi “assignment from incompatible pointer type” (gán giá trị sai kiểu) và code không thể compile được. Bạn chỉ có thể gán địa chỉ của một giá trị cho pointer cùng kiểu với nó.
Toán tử & trả về một pointer thuộc kiểu của biến mà nó đứng trước. Trong đoạn code trên &ival trả về một pointer thuộc kiểu int, fval trả về một pointer thuộc kiểu float và &sval trả về pointer thuộc kiểu string. Những chỗ nào mà bạn có thể dùng pointer thì cũng có thể sử dụng một biến &val tương ứng.
* Phép gán của con trỏ chỉ thực hiện được khi kiểu dữ liệu của con trỏ phù hợp kiểu dữ liệu của biến mà nó sẽ trỏ tới.

## Gán giá trị cho con trỏ
Giá trị mà biến con trỏ lưu trữ là địa chỉ của biến khác có cùng kiểu dữ liệu với biến con trỏ.
```go
var ptr *int;
var value int = 5;

ptr = &value;
```
Do đó, chúng ta cần sử dụng address-of operator để lấy ra địa chỉ ảo của biến rồi mới gán cho con trỏ được. Lúc này, biến ptr sẽ lưu trữ địa chỉ ảo của biến value.
![Image 5](https://raw.githubusercontent.com/sendo-research-team/sendo-research-team.github.io/gh-pages/static/images/5.JPG)
￼

Chúng ta có thể nói rằng con trỏ ptr đang nắm giữ địa chỉ của biến value, cũng có thể nói con trỏ ptr trỏ đến biến value.
Đoạn chương trình sau sẽ in ra địa chỉ của biến value và giá trị được lưu bởi con trỏ ptr sau khi trỏ đến biến value:
```go
func main()
{
	var value int = 5;
	var ptr *int = &value;
	
	fmt.Println(&value)
	fmt.Println(ptr)

}
```
Kết quả thu được trên màn hình console:
```go
0012FF7C
0012FF7C
```
Lý do mà chúng ta gán được địa chỉ của biến value cho con trỏ kiểu int (int *) là vì address-of operator của một biến kiểu int trả về giá trị kiểu con trỏ kiểu int (int *).
Do đó, chúng ta có thể gán &value cho con trỏ kiểu int (int *).

Bên cạnh đó, khi có hai con trỏ cùng kiểu thì chúng ta có thể gán trực tiếp mà không cần sử dụng address-of operator.
```go
int main()
{
	int value = 5
	var ptr1, ptr2 int

	ptr1 = &value; //ptr1 point to value
	ptr2 = ptr1;   //assign value of ptr1 to ptr2

	fmt.Println(ptr1)
	fmt.Println(ptr2)

}
```
Lúc này, ptr1 và ptr2 cùng giữ địa chỉ của biến value.

Truy xuất giá trị bên trong vùng nhớ mà con trỏ trỏ đến
Khi chúng ta có một con trỏ đã được trỏ đến địa chỉ nào đó trong bộ nhớ ảo, chúng ta có thể truy xuất giá trị tại địa chỉ đó bằng dereference operator. Dereference operator sẽ đánh giá nội dung địa chỉ được trỏ đến.
```go
var ptr *int *ptr //declare an int pointer
var value int = 5

ptr = &value //ptr point to value

fmt.Println(value) //print the address of value
fmt.Println(ptr)   //print the address of value which is held in ptr

fmt.Println(value) //print the content of value
fmt.Println(*(&value)) //print the content of value
fmt.Println(*ptr)	//print the content of value
```
Kết quả của đoạn chương trình trên như sau:

![Image 6](https://raw.githubusercontent.com/sendo-research-team/sendo-research-team.github.io/gh-pages/static/images/6.JPG)
￼
Toán tử trỏ đến (dereference operator) được dùng để truy cập trực tiếp vào vùng nhớ có địa chỉ cụ thể trên bộ nhớ ảo (virtual memory), vì biến con trỏ ptr đang giữ địa chỉ của biến value nên khi đặt toán tử trỏ đến (dereference operator) trước con trỏ ptr, nó sẽ truy xuất giá trị tại địa chỉ mà con trỏ ptr đang giữ.
Vì ptr có kiểu dữ liệu con trỏ int (int *), ptr chỉ có thể trỏ đến biến kiểu int. Lúc này, compiler hiểu rằng cần phân tích 4 bytes (đúng bằng kích thước kiểu int) trên bộ nhớ ảo tại địa chỉ mà ptr đang lưu trữ.
![Image 7](https://raw.githubusercontent.com/sendo-research-team/sendo-research-team.github.io/gh-pages/static/images/7.JPG)
￼

Đây là lý do tại sao chúng ta cần khai báo kiểu dữ liệu của con trỏ. Nếu không khai báo kiểu dữ liệu cho con trỏ, toán tử trỏ đến (dereference operator) sẽ không biết phải phân tích bao nhiêu bytes tại địa chỉ con trỏ trỏ đến để tính toán được giá trị của vùng nhớ đó. Không những thế, đây còn là lý do kiểu dữ liệu của biến phải tương xứng với kiểu dữ liệu được khai báo cho con trỏ.
Vì chúng ta có thể gán lại địa chỉ mới cho một con trỏ, nên chúng ta có thể truy xuất được giá trị của nhiều vùng nhớ khác nhau chỉ với một con trỏ:
```go
var value1 int = 1
var value2 int = 2

var ptr *int = &value1
fmt.Println(*ptr)

ptr = &value2
fmt.Println(*ptr)
```
Với khả năng truy cập đến vùng nhớ có địa chỉ cụ thể và thay đổi giá trị bên trong vùng nhớ của toán tử trỏ đến (dereference operator), chúng ta có thể sử dụng như sau:
```go
var value int = 5
var ptr *int = &value

*ptr = 10
fmt.Println(*ptr)
```
Đoạn chương trình này sẽ in ra giá trị 10.
Có thể giải thích dòng lệnh *ptr = 10; như sau:
Biến con trỏ ptr sau khi khai báo đã được khởi tạo bằng cách gán địa chỉ của biến value. Sử dụng dereference operator cho con trỏ ptr để truy cập đến địa chỉ ảo mà ptr đang nắm giữ, gán giá trị 10 vào vùng nhớ tại vị trí đó.

## Con trỏ chưa được gán địa chỉ
Con trỏ trong ngôn ngữ Golang vốn có rủi ro. Nếu sử dụng con trỏ không hợp lý có thể gây panic chương trình.
Khác với tham chiếu (reference), biến con trỏ có thể không cần khởi tạo giá trị ngay khi khai báo. Nhưng thực hiện truy xuất giá trị của con trỏ bằng dereference operator khi chưa gán địa chỉ cụ thể cho con trỏ, chương trình có thể bị đóng bởi hệ điều hành. Nguyên nhân là cố gắng đọc một con trỏ nil.
Đối với con trỏ, Nil là một giá trị đặc biệt (Zero Value), khi gán Nil cho con trỏ, điều đó có nghĩa là con trỏ đó chưa trỏ đến địa chỉ nào cả. Con trỏ đang giữ giá trị Nil được gọi là con trỏ Nil (Nil pointer).

## Tổng kết
Trong bài học này, các bạn đã được tìm hiểu khái niệm con trỏ và một số khái niệm có liên quan. Việc sử dụng con trỏ thường có một số hoạt động chủ yếu: 
(a) khai báo một con trỏ, 
(b) gán địa chỉ cho con trỏ,
(c) truy cập đến địa chỉ mà con trỏ đang nắm giữ bằng dereference operator.

Thử liên hệ một chút với cuộc sống thực tế, tưởng tượng rằng con đường nhà bạn (street) là bộ nhớ ảo, trên con đường đó có rất nhiều ngôi nhà (house), mỗi ngôi nhà đều được đánh số thứ tự gọi là địa chỉ nhà (house's address). Chúng ta tạm hình dung số người ở trong mỗi ngôi nhà (content) tương đương với nội dung của mỗi ô trên bộ nhớ ảo. Như vậy, address-of operator (&house) sẽ trả về địa chỉ của ngôi nhà, dereference operator (*&house) sẽ lấy ra số lượng người bên trong ngôi nhà có địa chỉ được xác định. Để sử dụng con trỏ trỏ đến mỗi ngôi nhà, chúng ta phải sử dụng một con trỏ kiểu House (giống với kiểu của từng ngôi nhà), giả sử con trỏ kiểu House được khai báo là House *h_ptr; thì con trỏ h_ptr có thể trỏ đến bất kì ngôi nhà nào trên con đường, và nó còn có thể thay đổi nội dung bên trong từng ngôi nhà mà nó trỏ đến.
Con trỏ (Pointer) là một công cụ mạnh mẽ đặc trưng của ngôn ngữ Golang. Con trỏ cho phép chúng ta chia sẻ data của chương trình trên bộ nhớ ảo.
