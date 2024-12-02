from collections import Counter

list1 = []
list2 = []
similarity_score = 0

with open("2024/day01/input.txt", "r") as file:
    for line in file.readlines():
        l, r = line.strip().split()
        list1.append(int(l))
        list2.append(int(r))
        
l_counter = Counter(list1)
r_counter = Counter(list2)

for id, count in l_counter.items():
    similarity_score += id * count * r_counter[id]
    
print(similarity_score)