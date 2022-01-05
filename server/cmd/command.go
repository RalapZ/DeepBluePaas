package cmd

import (
	"fmt"
	"github.com/RalapZ/DeepBluePaas/server/config"
	"github.com/RalapZ/DeepBluePaas/server/router"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type CobraFunc func(command *cobra.Command, args []string)

//var test  CobraFunc
var (
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
			router.Serve()
		},
	}
)

var myzoneCmd = &cobra.Command{
	Use:   "myzone",
	Long:  "myzone long 信息",
	Short: "myzone short 信息",
}

func init() {
	rootCmd.Execute()
	//Logger = logger.InitLogger()
	//rootCmd.Flags().StringP(version,"a","ralap","作者名称")
	//rootCmd.Flags().BoolP("start","y",true,"是否开启")
	config.ParserConfig()
}

func Execute() {
	rootCmd.Execute()
}
