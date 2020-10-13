#include <stdio.h>


class Base
{
	public:
	virtual void print()
	{
		printf("Base class print() \n");
	}
};

class Derived1: public Base
{
	public:
	void print()
	{
		printf("Derived1 class print() \n");
	}
};

main()
{

	printf("============= C++ class \n");
	Base b1;
	b1.print();
	Derived1 d1;
	d1.print();
	
	printf("============= Polimorphism: \n");
	Base* p;
	p = &b1;
	p->print();
	p = &d1;
	p->print();
	printf("\n\n");
}

