
include "HttpProxy.thrift"

service PhotoWall
{
	/*
	 * @func        ���Խӿ�ping.
	 * @return      �ɹ����� 0 ,ʧ�ܷ��ط�0.
	 */
	i32 ping();
	
	
	/*
	 * @func        �����ӽӿ�.
	 * @return      ����HTTP�ذ�
	 */
	HttpProxy.RespHttpPackage doRequest(1:HttpProxy.ReqHttpPackage  req, 2:map<string, string> extInfo);
}
