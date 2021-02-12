- New chỉ đưa cho biến cái địa chỉ trong thanh nhớ thôi

- Make không chỉ đưa biến cho cái địa chỉ mà còn đặt cả giá trị zero cho biến

make and new are built-in functions golang uses to allocate memory and allocate memory on the heap. 

make allocates and initializes memory.New just clears memory and does not initialize it.

make returns the reference type itself; new returns a pointer to the type.

make can only be used to allocate and initialize data of slice, map, channel type; new can allocate any type of data