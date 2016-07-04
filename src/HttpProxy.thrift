namespace cpp zeus
namespace cocoa zeus
namespace java zeus

// 客户端类型
enum E_CLIENT_TYPE {
	UNKNOW = 0, 
	ANDROID = 1,      
	IOS = 2,
	WEB = 3
	WEB_COSCHOOL = 4,   
	WEB_GROUP =5,
}        

// 消息请求命令字
enum E_HTTP_COMMAND {
	ECMD_NULL = 0,

	
	//照片墙
	ECMD_PHOTO_WALL_ADD = 1000,
	ECMD_PHOTO_WALL_DELETE = 1001,
	ECMD_PHOTO_WALL_QUERY =1002,
}        	

// 请求 HTTP协议包体数据
struct ReqHttpPackage
{
	1:string       sVer ;                 // 版本号
	2:i32          nClientType ;          //客户端类型
	3:string       uid ;		          // 用户id
	4:i32          nCMDID ;               // 命令字
	5:i64          sequence ;             // 请求包的序号
	6:binary       buf;                   // 请求的数据体
	7:string       sToken="";             //需要验证身份时带上
}

// 返回消息     
struct RespHttpPackage
{
	1:i32          nCMDID;         //命令字，与请求的对应
	2:i64          sequence ;     //与请求包对应
	3:i32          result ;       //状态  0成功，非零失败
	4:string       sMessage ;      //提示
	5:binary       buf;	           // 返回的数据体
}



