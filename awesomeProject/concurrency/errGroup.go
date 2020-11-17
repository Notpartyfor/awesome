package main

import (
	"fmt"
	xContext "golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	ctx, cancel := xContext.WithCancel(xContext.Background())
	group, errCtx := errgroup.WithContext(ctx)

	for index := 0; index < 3; index++ {
		indexTemp := index // 子协程中若直接访问index，则可能是同一个变量，所以要用临时变量

		// 新建子协程
		group.Go(func() error {
			fmt.Printf("indexTemp=%d \n", indexTemp)
			if indexTemp == 0 {
				fmt.Println("indexTemp == 0 start ")
				fmt.Println("indexTemp == 0 end")
			} else if indexTemp == 1 {
				fmt.Println("indexTemp == 1 start")
				//这里一般都是某个协程发生异常之后，调用cancel()
				//这样别的协程就可以通过errCtx获取到err信息，以便决定是否需要取消后续操作
				cancel()
				fmt.Println("indexTemp == 1 err ")
			} else if indexTemp == 2 {
				fmt.Println("indexTemp == 2 begin")

				// 休眠1秒，用于捕获子协程2的出错
				time.Sleep(1 * time.Second)

				//检查 其他协程已经发生错误，如果已经发生异常，则不再执行下面的代码
				err := CheckGoroutineErr(errCtx)
				if err != nil {
					return err
				}
				fmt.Println("indexTemp == 2 end ")
			}
			return nil
		})
	}

	// 捕获err
	err := group.Wait()
	if err == nil {
		fmt.Println("都完成了")
	} else {
		fmt.Printf("get error:%v", err)
	}
}

//校验是否有协程已发生错误
func CheckGoroutineErr(errContext xContext.Context) error {
	select {
	case <-errContext.Done():
		return errContext.Err()
	default:
		return nil
	}
}
