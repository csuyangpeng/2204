#include <iostream>
#include <string>

std::string printValue(int len, unsigned char* value)
{
//   std::ostringstream ostream;
//
//   ostream << "len(" << len  <<") value("; 
//
//   for(int i=0;i<len;i++)
//   {
//      cout << "test"<< hex << value[i]<< endl;
//      ostream << value[i];
//   }
//
//   ostream <<")";
//   
//   return ostream.str();

	std::string str("");  
	std::string str2("0123456789abcdef");   
	 for (int i=0;i<len;i++) {  
	   int b;  
	   b = 0x0f&(value[i]>>4);	
	   char s1 = str2.at(b);  
	   str.append(1,str2.at(b));			
	   b = 0x0f & value[i];  
	   str.append(1,str2.at(b));  
	   char s2 = str2.at(b);  
	 }	
	 return str;  
}


