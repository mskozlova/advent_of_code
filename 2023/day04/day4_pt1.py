def parse_card(line):
    number_lists = line[line.rindex(":") + 1 :].split("|")
    assert len(number_lists) == 2, f"Too many values after | split: {line}"

    winning_set = set(map(int, number_lists[0].strip().split()))
    our_set = set(map(int, number_lists[1].strip().split()))

    return winning_set, our_set


total_points = 0

with open("input.txt", "r") as file:
    for line in file.readlines():
        winning_set, our_set = parse_card(line)
        n_wins = len(winning_set.intersection(our_set))
        if n_wins > 0:
            total_points += 2 ** (n_wins - 1)

print(total_points)
