directories = dict()
stack = []

with open("input.txt", "r") as file:
    for line in file.readlines():
        if line.startswith("$ cd"):
            dir_name = line.strip()[5:]
            if dir_name == "..":
                stack.pop()
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

total_disk = 70000000
needed_disk = 30000000

occupied_disk = directories["/"]
need_freed = needed_disk - (total_disk - occupied_disk)

print("occupied: {}, need_freed: {}".format(occupied_disk, need_freed))

candidates = [(key, size) for key, size in directories.items() if size >= need_freed]
print(sorted(candidates, key=lambda x: x[1])[0])
