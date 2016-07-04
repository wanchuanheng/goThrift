
include "HttpProxy.thrift"

service PhotoWall
{
	/*
	 * @func        测试接口ping.
	 * @return      成功返回 0 ,失败返回非0.
	 */
	i32 ping();
	
	
	/*
	 * @func        短链接接口.
	 * @return      返回HTTP回包
	 */
	HttpProxy.RespHttpPackage doRequest(1:HttpProxy.ReqHttpPackage  req, 2:map<string, string> extInfo);
}
