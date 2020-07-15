[总后台新版接口列表](http://git.ddddian.com:3000/ddd/document/src/dev/interface/dddPublicInterface.md)  

# 申请消息--选择小店
> 活动列表接口文档，接口返回head和middle两个内容。具体参考如下。


#### 1. 接口URL
`{{url}}/api/Message/Shop`
#### 2. 开发人
郭志强
#### 3. 请求方式
`POST`
#### 4. 请求`Content-Type`
`multipart/form-data`
#### 5. 请求参数
| 参数 | 参数描述 | 类型 | 是否必填 | 示例值 |
| --- | --- | --- | --- | --- |
| Mobile | 手机号 | string | 是 | 13617675661 |
| Token | token | string | 是 | 889270BC-36EA-E36D-1443-FC7FF3DA6C46 |
| keyword | 关键字 | string | 否 |  |
| level | 等级id | string | 否 |  |
| type | 状态 | string | 否 | 1 合作 2 未合作 3 合作与为合作 |


#### 6. 成功响应示例
```
{
	"code": 200,
	"msg": "查询成功",
	"data": [
		{
			"Shop_Id": "SH102400000000000012",
			"Shop_Nm": "合师附小副食",
			"Mobl_Num": "18725732263",
			"Vip_Id": null
		}
	]
}

```
#### 7. 其他说明
暂无
