/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/8/2 20:03
*  @To:
 */

package wxchat

import "encoding/xml"

// 接收的所有消息和事件
type res struct {
	ToUserName   string `xml:"ToUserName"`   // 开发者微信号		（都有）
	FromUserName string `xml:"FromUserName"` // 发送方帐号（一个OpenID）	（都有）
	CreateTime   int64  `xml:"CreateTime"`   // 消息创建时间 （整型）	（都有）
	MsgType      string `xml:"MsgType"`      // 消息类型			（都有）
	/*---------------------*/
	Content      string `xml:"Content"`      // 文本消息内容		(MsgType为text时有)
	MsgID        int64  `xml:"MsgId"`        // 消息id，64位整型							(消息都有)
	PicUrl       string `xml:"PicUrl"`       // 图片链接（由系统生成）  (MsgType为image时有)	图片消息专有
	ThumbMediaId string `xml:"ThumbMediaId"` //	视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据	(MsgType为 video  shortvideo时有)		视频消息专有
	Format       string `xml:"Format"`       // 语音格式，如amr，speex等	(MsgType为voice时有)		语音消息专有
	Recognition  string `xml:"Recognition"`  // 语音识别结果，UTF8编码		(MsgType为voice时有)		语音消息专有
	MediaId      string `xml:"MediaId"`      // 媒体消息id，可以调用多媒体文件下载接口拉取数据。	(MsgType为image时有)  媒体消息专有
	Location_X   string `xml:"Location_X"`   // 地理位置纬度
	Location_Y   string `xml:"Location_Y"`   // 地理位置经度
	Scale        string `xml:"Scale"`        // 地图缩放大小
	Label        string `xml:"Label"`        // 地理位置信息
	Title        string `xml:"Title"`        // 消息标题		(MsgType为link时有)		链接消息专有
	Description  string `xml:"Description"`  // 消息描述		(MsgType为link时有)		链接消息专有
	Url          string `xml:"Url"`          // 消息链接		(MsgType为link时有)		链接消息专有
	MsgDataId    string `xml:"MsgDataId"`    // 消息的数据ID（消息如果来自文章时才有）
	Idx          string `xml:"Idx"`          // 多图文时第几篇文章，从1开始（消息如果来自文章时才有）
	/*-------------------*/
	Event    string `xml:"Event"`    // 事件类型	 (MsgType为event时有) subscribe(订阅)、unsubscribe(取消订阅)
	EventKey string `xml:"EventKey"` // 事件KEY值		(MsgType为event时有)
	Ticket   string `xml:"Ticket"`   // 二维码的ticket，可用来换取二维码图片	(Event为subscribe)
}

// Text 回复文本消息
type Text struct {
	ToUserName   string   `xml:"ToUserName"`   // 开发者微信号		（都有）
	FromUserName string   `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64    `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      string   `xml:"MsgType"`      // 消息类型
	Content      string   `xml:"Content"`      // 文本消息内容
	XMLName      xml.Name `xml:"xml"`          // 必须要有
}

// News 回复图文消息
type News struct {
	ToUserName   string    `xml:"ToUserName"`   // 开发者微信号		（都有）
	FromUserName string    `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64     `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      string    `xml:"MsgType"`      // 消息类型
	ArticleCount int       `xml:"ArticleCount"` // 图文消息个数
	Articles     []Article `xml:"Articles>item"`
	XMLName      xml.Name  `xml:"xml"` // 必须要有
}

type Article struct {
	Title       string `xml:"Title"`       // 图文消息标题
	Description string `xml:"Description"` // 图文消息描述
	PicUrl      string `xml:"PicUrl"`      // 图片链接，支持JPG、PNG格式，较好的效果为大图360*200，小图200*200
	Url         string `xml:"Url"`         // 点击图文消息跳转链接
}
