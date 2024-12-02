list1 = []
list2 = []
diff = 0

with open("2024/day01/input.txt", "r") as file:
    for line in file.readlines():
        l, r = line.strip().split()
        list1.append(int(l))
        list2.append(int(r))
        
for l, r in zip(sorted(list1), sorted(list2)):
    diff += abs(l - r)
    
print(diff)