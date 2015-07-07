#!/bin/env pypy
#coding:utf-8
#usage : file summary for certain key.
import sys
import time
import pdb 

def sample_read(file_path):
	with open(file_path,'r') as f:
		datas = f.readlines()
	return datas

def key_alphabet(file_path, key_list, val_list) :
	kv_dict={}

	file_name = 'k%s_v%s.txt' % (''.join(key_list), ''.join(val_list))

	key_list = [ int(idx) for idx in key_list ]
	val_list = [ int(idx) for idx in val_list ]


	for line in open(file_path,'r'):

		data = line.strip().split(',')
		key = ','.join( [ data[idx] for idx in key_list ] )

		try :
			kv_data = kv_dict[key]

			pdb.set_trace()
			for enum, idx in enumerate(val_idx) :
				kv_data[enum] += int(data[idx])

		except :
			kv_dict[key] = [ int(data[idx]) for idx in val_list ]

	f = open('/dev/shm/TEST/%s' % file_name,'w') 
	for k in kv_dict.keys() :
		f.write("%s : %s\n" %(k, ','.join([str(i) for i in kv_dict[k]])))
	f.close()
	sys.stdout.write('file://%s\n' % f.name)
	sys.stdout.flush()


def key_odd_eval(datas) :
	pass

key_list = sys.argv[1].split(',')
val_list = sys.argv[2].split(',')

test_input = sys.stdin.readline()
test_input = test_input.strip()

file_path = test_input.split("://")[1]

st_time = time.time()
key_alphabet(file_path, key_list, val_list)

sys.stderr.write(test_input + '\n')
sys.stderr.write('Processing Time : %.3f\n' % (time.time() - st_time))
sys.stderr.flush()
#key_odd_eval(datas)




time.sleep(888888888)
