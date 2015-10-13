/*
注释
*/
namespace csharp Idl

struct Blog {
    1: string topic 
    2: binary content 
    3: i64    createdTime 
    4: string id 
    5: string ipAddress 
    6: map<string,string> props 
	7: string note
	8: i32 notid
  }