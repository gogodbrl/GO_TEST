#/bin/env python
#usage : compare golang in eventflow

import time
import random
import sys

READLINE = 100000

def parser() :
	f = open('/dev/shm/TEST/sample_meta.txt','w') 
	for line in open('/dev/shm/TEST/sample_input.txt','r') :
		line = line.strip()
		data = line.split(',')

		if line=='END' :
			break

		num = int(data[2])+int(data[3])
		if num % 2 == 0 : #eval
			f.write("%s,%s\n" %(line, 0))
		else : #odd
			f.write("%s,%s\n" %(line, 1))
	f.close()

	sys.stdout.write("file://%s\n" % f.name)
	sys.stdout.flush()
#sample_make_write()

st_time = time.time()
parser()
sys.stderr.write('Processing Time : %.3f\n' % (time.time() - st_time))
sys.stderr.flush()

time.sleep(999999999)
