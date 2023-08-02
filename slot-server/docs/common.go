package docs

func init() {
	//SwaggerInfo.InfoInstanceName = "XXXApi"
	//SwaggerInfo.Host = "127.0.0.1:8888"
	SwaggerInfo.Host = "api.bigwin.money"
	SwaggerInfo.Description =
		"## Seamless Wallet HTTP 服务\n" +
			"这是一个供 xxx Play 游戏平台连接到玩家钱包的简单 API。" +
			"所有请求与响应类型都应为 application/json POST。\n" +
			"### 哈希签名通过以下公式来计算：\n" +
			"从请求 POST 参数获取所有参数（预期哈希）并附加到字符串：\n" +
			"1. 对排除sign字段外的所有参数进行字典序排序。\n" +
			"2. 在 key1=value1&key2=value2 中附加它们。\n" +
			"3. 附加密钥，即：key1=value1&key2=value2{secret}。大括号省略\n" +
			"4. 使用 MD5 计算哈希。\n" +
			"5. 与哈希参数对比。如果失败，娱乐场运营商应发送错误代码 code 7。\n" +
			"### 运营商请求通用参数\n" +
			"运营商必须在每个请求体中包含以下通用参数：\n" +
			"1. agent - 运营商 ID\n" +
			"2. sign - 通过上述公式计算的哈希\n" +
			"3. timestamp - 毫秒级时间戳\n" +
			"### 供应商请求通用参数\n" +
			"供应商将在每个请求体中包含以下通用参数：\n" +
			"1. providerId - 供应商ID 可由娱乐场运营商提供。否则使用默认标识符 \n" +
			"2. token - 运营商在鉴权接口中提供的令牌\n" +
			"3. sign - 通过上述公式计算的哈希\n" +
			"4. timestamp - 毫秒级时间戳\n" +
			"### 通用响应\n" +
			"每个响应都应包含以下参数：\n" +
			"1. code - 7 失败 0 成功 \n" +
			"2. data - 数据\n" +
			"3. msg - 消息\n"
	// 读取docs/swagger.json文件
	//bytes, err := os.ReadFile("./docs/swagger.json")
	//if err == nil {
	//	sjson, err := simplejson.NewJson(bytes)
	//	if err == nil {
	//		sjson.Set("tags", map[string]interface{}{
	//			"name":        "Seamless-Provider",
	//			"description": "Seamless-Providerxxxabc",
	//		})
	//	}
	//	v, _ := sjson.MarshalJSON()
	//	os.WriteFile("./docs/swagger.json", v, 0666)
	//}
}
