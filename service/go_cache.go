package service

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)
func cacheTest() {
  // 创建一个默认过期时间为 5 分钟的缓存，每 10 分钟清除一次过期项目
  c := cache.New(5*time.Minute, 10*time.Minute)

  // 使用默认过期时间，将 key "foo" 的 vaule 设置为 "bar"
  c.Set("foo", "bar", cache.DefaultExpiration)

  // 将 key "baz" 的 vaule 设置为 42，并设置为永不过期
  // (这个 item 将不会被移除，直到缓存被重置，或者通过 c.Delete("baz") 移除。)
  c.Set("baz", 42, cache.NoExpiration)

  // 从缓存中获取与 key "foo" 关联的字符串
  v1, found := c.Get("foo")
  if found {
	  fmt.Printf("v1: %s\n", v1)
  }

  // 因为 Go 是静态类型的，并且缓存值可以是任何东西，
  // 所以当值被传递给不采用任意类型的函数（即 interface{}）时，需要类型断言。
  // 对只会使用一次的值,执行此操作的最简单方法是：
  // 示例 传递一个值给函数
  v2, found := c.Get("foo")
  if found {
	  MyFunction(v2.(string))
  }

  // 如果在同一个函数中多次使用该值，这将变得乏味。
  // 可以改为执行以下任一操作：
  if x, found := c.Get("foo"); found {
	  v3 := x.(string)
	  fmt.Printf("v3: %s\n", v3)

	  // ...
  }
  // or
  var v4 string
  if x, found := c.Get("foo"); found {
	  v4 = x.(string)
  }
  fmt.Printf("v4: %s\n", v4)
  // ...
  // 然后 foo 可以作为字符串自由传递

  // Want performance? Store pointers!
  c.Set("foo", &MyStruct{Msg: "test"}, cache.DefaultExpiration)
  if x, found := c.Get("foo"); found {
	  v5 := x.(*MyStruct)
	  fmt.Printf("v5: %s\n", v5)
	  // ...
  }
}

type MyStruct struct {
  Msg string
}

func MyFunction(msg string) {
  fmt.Println(msg)
}