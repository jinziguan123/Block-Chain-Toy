/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-01-20 14:38:16
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-01-20 16:46:34
 * @FilePath: /goBlockChain/transaction/inoutput.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package transaction

import "bytes"

type TxOutput struct {
	Value     int
	ToAddress []byte
}

type TxInput struct {
	TxID        []byte
	OutIdx      int
	FromAddress []byte
}

func (in *TxInput) FromAddressRight(address []byte) bool {
	return bytes.Equal(address, in.FromAddress)
}

func (out *TxOutput) ToAddressRight(address []byte) bool {
	return bytes.Equal(address, out.ToAddress)
}
