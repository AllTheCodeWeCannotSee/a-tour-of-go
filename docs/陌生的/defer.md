

我们将使用下面这个例子来制作时序图，这个例子同时展示了 **LIFO（后进先出）** 和 **参数立即求值** 这两个核心特性。

### 示例代码

```go
package main

import "fmt"

func flowExample() {
	fmt.Println("函数主体：开始执行")

	i := 10

	// 1. 注册第一个 defer，此时 i 的值是 10
	defer fmt.Println("defer 1 (后执行): 此处 i 的值为", i)

	i = 20 // 修改 i 的值

	// 2. 注册第二个 defer
	defer fmt.Println("defer 2 (先执行)")

	fmt.Println("函数主体：即将返回，此时 i 的值为", i)
}

func main() {
	flowExample()
}
```

#### 预期输出

在看时序图之前，我们先看预期输出，这样可以带着问题去理解流程：

```
函数主体：开始执行
函数主体：即将返回，此时 i 的值为 20
defer 2 (先执行)
defer 1 (后执行): 此处 i 的值为 10
```

-----

### Defer 流程时序图

下面是上述代码执行流程的可视化时序图。

```mermaid
sequenceDiagram
    participant main as main()
    participant flow as flowExample()
    participant runtime as Go 运行时 (Defer 栈)
    participant fmt as fmt.Println()

    main->>flow: 调用 flowExample()
    activate flow

    %% ---- 函数主体执行阶段 ----
    flow->>fmt: Println("函数主体：开始执行")
    note right of flow: i := 10

    note over flow,runtime: 遇到 defer 1<br/>立即对参数求值，捕获到 i = 10
    flow->>runtime: **推入**调用: fmt.Println("...", 10)

    note right of flow: i = 20

    note over flow,runtime: 遇到 defer 2
    flow->>runtime: **推入**调用: fmt.Println("defer 2...")

    flow->>fmt: Println("函数主体：即将返回，此时 i 的值为 20")

    %% ---- 函数返回前的 Defer 执行阶段 ----
    note over flow,runtime: flowExample() 即将返回，触发 Defer 栈 (LIFO)
    
    runtime-->>fmt: [后进先出 1] **执行**栈顶调用: fmt.Println("defer 2...")
    runtime-->>fmt: [后进先出 2] **执行**下一个调用: fmt.Println("...", 10)

    deactivate flow
    flow-->>main: 返回
```

