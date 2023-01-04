directories = dict()
stack = []

with open("input.txt", "r") as file:
    for line in file.readlines():
        if line.startswith("$ cd"):
            dir_name = line.strip()[5:]
            if dir_name == "..":
                stack.pop()
            elif dir_name == "/":
                stack = []
            else:
                stack.append(dir_name)
                directories["/".join(stack)] = 0

        elif line.startswith("$ ls"):
            continue

        elif line.startswith("dir"):
            continue

        elif line[0].isdigit():
            size, _ = line.split(" ")
            for i in range(len(stack)):
                directories["/".join(stack[:i + 1])] += int(size)

        else:
            print("unknown type of line: {}".format(line))

total_sum = 0
for size in directories.values():
    if size <= 100000:
        total_sum += size

print(total_sum)
