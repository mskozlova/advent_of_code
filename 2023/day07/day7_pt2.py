from collections import Counter

card_labels = ["J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"]


def parse_hand(line):
    data = line.strip().split()
    assert len(data) == 2

    return data[0], int(data[1])


def parse_input(file_name):
    hands = []
    with open(file_name, "r") as file:
        for line in file.readlines():
            hands.append(parse_hand(line))
    return hands


def get_hand_type(hand):
    card_counter = Counter(hand[0])
    if "J" in hand[0]:
        j_number = card_counter.pop("J")
    else:
        j_number = 0

    # from best to worst:
    # five of a kind - 6
    # four of a kind - 5
    # full house - 4
    # three of a kind - 3
    # two pair - 2
    # one pair - 1
    # high card - 0

    if len(card_counter) <= 1:  # five of a kind
        return 6

    if len(card_counter) == 5:  # high card
        return 0

    if card_counter.most_common()[0][1] == 1:
        # j_number == 4 - then len(card_counter) == 1, already covered

        if j_number == 3:  # four of a kind
            return 5

        if j_number == 2:  # three of a kind
            return 3

        if j_number == 1:  # one pair
            return 1

        # j_number == 0 - then len(card_counter) == 5, already covered

    if card_counter.most_common()[0][1] == 2:
        # j_number == 3 - then len(card_counter) == 1, already covered

        if j_number == 2:  # four of a kind
            return 5

        if j_number == 1:
            if card_counter.most_common()[1][1] == 2:  # full house
                return 4
            else:  # three of a kind
                return 3

        if j_number == 0:
            if card_counter.most_common()[1][1] == 2:  # two pair
                return 2
            else:  # one pair
                return 1

    if card_counter.most_common()[0][1] == 3:
        # j_number == 2 - then len(card_counter) == 1, already covered

        if j_number == 1:  # four of a kind
            return 5

        if j_number == 0:
            if card_counter.most_common()[1][1] == 2:  # full house
                return 4
            else:  # three of a kind
                return 3

    if card_counter.most_common()[0][1] == 4:  # four of a kind
        # j_number == 1 - then len(card_counter) == 1, already covered
        return 5

    raise Exception(f"type not found for hand: {hand[0]}")


def get_hand_code(hand):
    code = []
    for symbol in hand[0]:
        code.append(hex(card_labels.index(symbol)))

    return "".join(code)


hands = parse_input("input.txt")

total_score = 0
for i, hand in enumerate(
    sorted(hands, key=lambda hand: (get_hand_type(hand), get_hand_code(hand)))
):
    total_score += (i + 1) * hand[1]

print(total_score)
