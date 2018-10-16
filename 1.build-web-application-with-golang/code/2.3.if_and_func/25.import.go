package main

//_操作其实是引入该包，而不直接使用包里面的函数，而是调用包里面的init
import (
	"database/sql" //引用 /usr/local/go/src/database/sql
	. "fmt"
	f "fmt"
	"github.com/google/uuid"   ///root/www/go_www/src/github.com/andyxning/shortme
	_ "github.com/google/uuid" ///root/www/go_www/src/github.com/andyxning/shortme
)

func main() {
	Println("yes")
	f.Println("yes")
}
