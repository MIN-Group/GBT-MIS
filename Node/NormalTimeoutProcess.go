package Node

import (
	"fmt"

	"GBT-MIS/MetaData"
	"GBT-MIS/utils"
)

func (node *Node) NormalTimeOutProcess() {
	height := node.mongo.GetHeight() + 1
	value, ok := node.BlockGroups.Load(height)
	if ok {
		item := value.(MetaData.BlockGroup)
		if item.ReceivedBlockGroupHeader {
			deVoteResult := utils.DecompressToIntSlice(item.VoteResult)
			for i := 0; i < len(deVoteResult); i++ {
				if deVoteResult[i] == 1 {
					if !item.Blocks[i].IsSet {
						header, msg := node.msgManager.CreateRequestBlockMsg(0, height, i)
						node.SendMessage(header, &msg)
						fmt.Println(node.network.MyNodeInfo.ID, "向节点0请求高度为", height, "区块号为", i, "的区块!")
					}
				}
			}
		} else {
			header, msg := node.msgManager.CreateRequestBlockGroupHeaderMsg(0, height)
			node.SendMessage(header, &msg)
			fmt.Println(node.network.MyNodeInfo.ID, "向节点0请求高度为", height, "的区块组头!")
		}
	}
}
