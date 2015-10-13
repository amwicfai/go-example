include "Blog.thrift"

namespace csharp Idl

service BlogService {  #  注释3 
    i32 testCase1(1:i32 num1, 2:i32 num2, 3:string  num3)  #  注释4

    list<string> testCase2(1:map<string,string>  num1)

    void testCase3()

    void testCase4(1:list<Blog.Blog> blog)  
}