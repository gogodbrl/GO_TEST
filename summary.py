#/bin/env python
#coding:utf-8
#usage : file summary for certain key.

def sample_read():
	with open('sample_output.txt','r') as f:
		datas = f.readlines()
	return datas

def key_alphabet(datas) :
	kv_dict={}

	for line in datas :
		line = line.strip()
		data = line.split(',') #['a', 'b', '1', '5', '0']
		
		if line=='END' :
			break

		key = data[0]+data[1]
		if not kv_dict.has_key(key):
			kv_dict[key]=[0,0] # [odd_count, eval_count]
	
		#ab의 값중에 짝수는 x번 출현 홀수는 y번 출현했다.
		odd_count = int(kv_dict[key][0])
		eval_count = int(kv_dict[key][1])

		if int(data[4]) == 0 : #eval
			kv_dict[key]=[odd_count,eval_count+1]
		else : #odd
			kv_dict[key]=[odd_count+1,eval_count]

	f = open('sample_result.txt','w') 
	for k in kv_dict.keys() :
		f.write("%s %s\n" %(k, kv_dict[k]))
	f.close()


def key_odd_eval(datas) :
	pass


datas = sample_read()
key_alphabet(datas)
#key_odd_eval(datas)
