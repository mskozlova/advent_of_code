def parse_card(line):
    number_lists = line[line.rindex(":") + 1 :].split("|")
    assert len(number_lists) == 2, f"Too many values after | split: {line}"

    winning_set = set(map(int, number_lists[0].strip().split()))
    our_set = set(map(int, number_lists[1].strip().split()))

    return winning_set, our_set


total_cards = 0
card_copies = dict()
is_start = True

with open("input.txt", "r") as file:
    for i, line in enumerate(file.readlines()):
        winning_set, our_set = parse_card(line)
        n_wins = len(winning_set.intersection(our_set))

        # first - how many copies of this card did we have?
        current_card = i + 1
        n_copies_current = card_copies.get(current_card, 1)

        total_cards += n_copies_current

        # second - how many copies this card generated further?
        for next_idx in range(current_card + 1, current_card + 1 + n_wins):
            card_copies[next_idx] = card_copies.get(next_idx, 1) + n_copies_current

print(total_cards)
