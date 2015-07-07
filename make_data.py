#/bin/env python
#usage : compare golang in eventflow

import random

READLINE = 100000

def sample_make_write():
	f = open('/dev/shm/TEST/sample_input.txt','w')
	sample_string='abcde'

	for i in xrange(READLINE):
		f.write("%s,%s,%s,%s\n" %(random.choice(sample_string), random.choice(sample_string), random.randint(1,5), random.randint(5,10)))
	f.write('END')
	f.close()


sample_make_write()
