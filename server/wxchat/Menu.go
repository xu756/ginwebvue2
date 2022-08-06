/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/8/3 19:36
*  @To:
 */

package wxchat

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type GetMenu struct {
	Is_menu_open  int         `json:"is_menu_open"`
	Selfmenu_info interface{} `json:"selfmenu_info"`
}

func GetMenuRouter(c *gin.Context) {
	url := "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=" + GetAccessToken()
	method := "GET"
	req, _ := http.NewRequest(method, url, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var getMenu GetMenu
	json.Unmarshal(body, &getMenu)
	c.JSON(200, gin.H{
		"type": "success",
		"msg":  "获取菜单成功",
		"data": getMenu.Selfmenu_info,
	})
}
