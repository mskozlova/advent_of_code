import json
from itertools import zip_longest

def compare(lst1, lst2):
    for l1, l2 in zip_longest(lst1, lst2, fillvalue=None):
        if l1 is None:
            return True
        elif l2 is None:
            return False
        elif isinstance(l1, int) and isinstance(l2, int):
            if l1 > l2:
                return False
            elif l1 < l2:
                return True
        elif isinstance(l1, list) and isinstance(l2, list):
            if not compare(l1, l2):
                return False
        elif isinstance(l1, int) and isinstance(l2, list):
            if not compare([l1], l2):
                return False
        else:
            if not compare(l1, [l2]):
                return False

    return True

correct_cnt = 0


with open("input.txt", "r") as file:
    active_lines = []
    for line in file.readlines():
        if len(active_lines) < 2:
            active_lines.append(json.loads(line))
        else:
            for l in active_lines:
                print(l)
            if compare(*active_lines):
                correct_cnt += 1
            print(compare(*active_lines), "\n")
            active_lines = []

print(correct_cnt)
