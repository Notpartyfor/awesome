package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	group, _ := errgroup.WithContext(ctx)

	for index := 0; index < 3; index++ {
		indexTemp := index // 子协程中若直接访问index，则可能是同一个变量，所以要用临时变量

		group.Go(func() error {
			fmt.Printf("indexTemp=%d \n", indexTemp)
			if indexTemp == 0 {
				fmt.Println("indexTemp == 0 start ")
				fmt.Println("indexTemp == 0 end")
			} else if indexTemp == 1 {
				fmt.Println("indexTemp == 1 start")
				// 如果出错了
				fmt.Println("indexTemp == 1 err ")
				return errors.New("indexTemp == 1时出错了")
			} else if indexTemp == 2 {
				fmt.Println("indexTemp == 2 begin")
				//当indexTemp == 1时出错，会使group.err != nil
				//但要停止当时其他goroutine的运作，还是得在goroutine中加入控制
				//
				//休眠1秒，用于捕获出错
				time.Sleep(1 * time.Second)
				//
				//检查 其他goroutine已经发生错误，如果已经发生异常，则不再执行下面的代码
				//err := CheckGoroutineErr(errCtx)
				//if err != nil {
				//	return err
				//}
				fmt.Println("indexTemp == 2 end ")
			}
			return nil
		})
	}

	// 捕获err
	if err := group.Wait(); err != nil {
		fmt.Printf("get error:%v", err)
	} else {
		fmt.Println("all finished")
	}
}

//校验组中是否goroutine已发生错误
func CheckGoroutineErr(errCtx context.Context) error {
	select {
	case <-errCtx.Done():
		return errCtx.Err()
	default:
		return nil
	}
}
