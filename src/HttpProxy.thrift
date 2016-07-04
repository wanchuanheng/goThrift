namespace cpp zeus
namespace cocoa zeus
namespace java zeus

// �ͻ�������
enum E_CLIENT_TYPE {
	UNKNOW = 0, 
	ANDROID = 1,      
	IOS = 2,
	WEB = 3
	WEB_COSCHOOL = 4,   
	WEB_GROUP =5,
}        

// ��Ϣ����������
enum E_HTTP_COMMAND {
	ECMD_NULL = 0,

	
	//��Ƭǽ
	ECMD_PHOTO_WALL_ADD = 1000,
	ECMD_PHOTO_WALL_DELETE = 1001,
	ECMD_PHOTO_WALL_QUERY =1002,
}        	

// ���� HTTPЭ���������
struct ReqHttpPackage
{
	1:string       sVer ;                 // �汾��
	2:i32          nClientType ;          //�ͻ�������
	3:string       uid ;		          // �û�id
	4:i32          nCMDID ;               // ������
	5:i64          sequence ;             // ����������
	6:binary       buf;                   // �����������
	7:string       sToken="";             //��Ҫ��֤���ʱ����
}

// ������Ϣ     
struct RespHttpPackage
{
	1:i32          nCMDID;         //�����֣�������Ķ�Ӧ
	2:i64          sequence ;     //���������Ӧ
	3:i32          result ;       //״̬  0�ɹ�������ʧ��
	4:string       sMessage ;      //��ʾ
	5:binary       buf;	           // ���ص�������
}



