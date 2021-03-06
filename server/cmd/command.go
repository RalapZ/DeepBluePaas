package cmd

import (
	"fmt"
	"github.com/RalapZ/DeepBluePaas/server/config"
	"github.com/RalapZ/DeepBluePaas/server/grpc"
	"github.com/RalapZ/DeepBluePaas/server/router"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

type CobraFunc func(command *cobra.Command, args []string)

//var test  CobraFunc
var (
	ProDone chan int

	version = "autor"

	Logger *zap.Logger

	rootCmd = &cobra.Command{
		Use:   "init",
		Short: "信息",
		Long:  "root信息",
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Printf("参数 %##v",args)

			fmt.Println(cmd.Flags().GetString(version))
			//logger.InitLogger()
			//router.Serve()
			manager()
		},
	}
	myzoneCmd = &cobra.Command{
		Use:   "myzone",
		Long:  "myzone long 信息",
		Short: "myzone short 信息",
	}
)


//

func init() {
	//rootCmd.Execute()
	//Logger = logger.InitLogger()
	//rootCmd.Flags().StringP(version,"a","ralap","作者名称")
	//rootCmd.Flags().BoolP("start","y",true,"是否开启")

}

func Execute() {
	rootCmd.Execute()
}

func exitHandle(exitChan chan os.Signal) {
	for {
		select {
		case sig := <-exitChan:
			fmt.Println("接受到来自系统的信号：", sig)
			ProDone <- 1
			//os.Exit(1) //如果ctrl+c 关不掉程序，使用os.Exit强行关掉
		}
	}

}

func SignalNotify(){
	processChan:=make(chan os.Signal)
	signal.Notify(processChan,os.Interrupt,os.Kill,syscall.SIGTERM)
	exitHandle(processChan)
}

func manager(){
	go SignalNotify()
	config.ParserConfig()
	go dbgrpc.Startkratos()
	go func(){
		for{
			select {
			case <- ProDone:
				os.Exit(1)
			}
		}
	}()
	router.Serve()
}