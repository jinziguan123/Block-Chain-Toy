/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-01-20 00:54:47
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-01-20 23:33:19
 * @FilePath: /goBlockChain/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"goblockchain/cli"
	"os"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
