package main

import (
	"fmt"

	"example.com/generic"
)

/**
一开始有 workspace 的概念还是在 [Rust Book](https://doc.rust-lang.org/stable/book/ch14-03-cargo-workspaces.html) 中了解到的。我对它印象并不深刻，因为没有它，一样可以写项目，并跑起来。
但随着学习的深入，对开源项目的更多关注，发现它被用在很多地方的同时也了解到使用它的更多好处，这不得不让我重新注意它、[了解它](https://github.com/suhanyujie/article-transfer-rs/blob/main/src/using_wasmer_for_plugins/part1.md)。

而随着 go 1.18.0 版本的发布， workspace 再一次出现在我的视野中，我在想，它的作用是否与 Rust 中的 workspace 概念类似？


*/
func main() {
	fmt.Println("hello ws...")
	// 虽然 goland 报红了，但是其实是可以编译并执行的。
	// https://go.dev/doc/tutorial/workspaces
	// https://go.dev/ref/mod#workspaces
	generic.Say()
}
