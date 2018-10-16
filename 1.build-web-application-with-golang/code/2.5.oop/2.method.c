#include <stdio.h>

struct box{
	int a;
	int b;
};

void func1(struct box *b)
{
	b->a=123;
	b->b=456;
	return;
}

int main(int argc, const char *argv[])
{
	struct box b1;

	func1(&b1);

	printf("%d\n",b1.a);

	return 0;
}
