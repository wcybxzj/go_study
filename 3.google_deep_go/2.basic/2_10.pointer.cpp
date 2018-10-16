#include <iostream>
using namespace std;

//c++值传递
void pass_by_val(int a)
{
	a++;
}

//c++引用传递
void pass_by_ref(int& a)
{
	a++;
}

int main(int argc, const char *argv[])
{
	int a=3;
	pass_by_val(a);
	cout<<a<<endl;//3


	pass_by_ref(a);
	cout<<a<<endl;//4

	return 0;
}
