## x-ecutioner

This is just a script plus some code that will generate, from a linux ELF binary, a compressed version (so your Go linux binary could be smaller!)

Requirements:

	go > 1.13
	go-binddata
	
Usage:

	./generate.sh BINARYFILE
	
This wil generate a new file, called *BINARYFILE_mem*, that hopefully will be smaller than the original.


>For small binaries this doesn't work, it will actually generate a bigger binary! So try and see.

The new binary will be slightly slower at start. See below.

memrun.go is originally from here https://github.com/guitmz/memrun, modified for this usage.

#### My tests

**Compression:**

	13197468 binary1
	90451972 binary1_mem
	
	(from 13mb to 9mb)
	
	27190864 binary2
	19457229 binary2_mem
	
	(from 27mb to 19mb)
	
**Execution time:**

	time ./binary1
	...
	real    0m0.014s
	user    0m0.012s
	sys     0m0.004s

and the compressed one:

	time ./binary1_mem 
	...
	real    0m0.348s
	user    0m0.327s
	sys     0m0.047s
	
another example:

	time ./binary2
	...
	real    0m0.007s
	user    0m0.000s
	sys     0m0.008s
	
and the compressed version:

	time ./binary2_mem 
	...
	real    0m0.630s
	user    0m0.600s
	sys     0m0.064s


	real    0m0.630s
	user    0m0.600s
	sys     0m0.064s

