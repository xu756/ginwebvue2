/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/8/1 23:34
*  @To: 获取access_token 缓存到redis
 */

package wxchat

import (
	"encoding/json"
	"example.com/mod/cache"
	"example.com/mod/config"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GetAccessTokenRouter(c *gin.Context) {
	c.JSON(200, gin.H{
		"access_token": GetAccessToken(),
	})
}

func GetAccessToken() string {
	if cache.Exists("access_token") {
		return cache.Get("access_token").(string)
	}
	//获取access_token
	accessToken := GetAccessTokenFromWx()
	//缓存access_token
	cache.Set("access_token", accessToken, 60*60*2)
	return accessToken
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetAccessTokenFromWx() string {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + config.InitData.Wxcgzh.AppID + "&secret=" + config.InitData.Wxcgzh.AppSecret
	method := "GET"
	req, _ := http.NewRequest(method, url, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var at AccessToken
	json.Unmarshal(body, &at)
	return at.AccessToken

}
