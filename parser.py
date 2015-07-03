#/bin/env python
#usage : compare golang in eventflow

import random

READLINE = 10

def sample_make_write():
	f = open('sample_raw.txt','w')
	sample_string='ab'

	for i in xrange(READLINE):
		f.write("%s,%s,%s,%s\n" %(random.choice(sample_string), random.choice(sample_string), random.randint(1,5), random.randint(5,10)))
	f.write('END')
	f.close()

def parser() :
	with open('sample_raw.txt','r') as f:
		datas = f.readlines()

	f = open('sample_output.txt','w') 
	for line in datas : 
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

#sample_make_write()
parser()
