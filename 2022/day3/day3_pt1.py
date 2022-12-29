priorities = 0

def find_error_symbol(half1, half2):
    return set(half1).intersection(set(half2)).pop()

sym_2_priority = dict(
    [(chr(ord_num + ord('a')), priority) for ord_num, priority in zip(range(26), range(1, 27))] +
    [(chr(ord_num + ord('A')), priority + 26) for ord_num, priority in zip(range(26), range(1, 27))],
)

print(sym_2_priority)

with open("input.txt", "r") as file:
    for line in file.readlines():
        new_line = line.strip()
        sym = find_error_symbol(
            new_line[:int(len(new_line) / 2)],
            new_line[int(len(new_line) / 2):]
        )
        priorities += sym_2_priority[sym]

print(priorities)
