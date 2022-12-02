with open ('input.txt', 'r') as file:
    inputs_list = [str(i.replace('\n','')) for i in file.readlines()]
sum_list = []

n = 0

for line in inputs_list:
    if line == "":
        sum_list.append(n)
        n = 0
    else:
        n += int(line)

#Part 1
print(f"Part 1 answer: {max(sum_list)}")

#Part 2
sum_list.sort()
print(f"Part 2 answer: {sum(sum_list[len(sum_list)-3:])}")
