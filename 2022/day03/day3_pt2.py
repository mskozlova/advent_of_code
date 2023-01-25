priorities = 0


def find_badge_symbol(sack1, sack2, sack3):
    return set(sack1).intersection(set(sack2)).intersection(set(sack3)).pop()


sym_2_priority = dict(
    [
        (chr(ord_num + ord("a")), priority)
        for ord_num, priority in zip(range(26), range(1, 27))
    ]
    + [
        (chr(ord_num + ord("A")), priority + 26)
        for ord_num, priority in zip(range(26), range(1, 27))
    ],
)

with open("input.txt", "r") as file:
    three_sacks = []

    for line in file.readlines():
        new_line = line.strip()
        three_sacks.append(new_line)

        if len(three_sacks) == 3:
            sym = find_badge_symbol(*three_sacks)
            three_sacks = []

            priorities += sym_2_priority[sym]

print(priorities)
