# my_dict={}
# my_dict['/']={}
# extra_dict={"extra_dict":[{"file":34564}]}
# my_dict['/']=[{"file":1522},{"fil4":152}]
# my_dict['/'].append(extra_dict)
# print(my_dict)

with open ('input.txt', 'r') as file:
    dict_list = [str(i.replace('\n','')) for i in file.readlines()]


#TODO remove
dict_list=dict_list[:10]

n=0
for line in dict_list:
    if n==0:
        previous_command = "not applicable"
    else: 
        previous_command=dict_list[n]
    
    if line.startswith('$'):
        line_type='command'
    elif line.startswith('dir'):
        line_type = "folder"
    else:
        line_type = "file"
    print("\nline:",line)
    print("line_type:",line_type)
    print("previous_command:",previous_command)

    n+=1

# print(dict_list)
