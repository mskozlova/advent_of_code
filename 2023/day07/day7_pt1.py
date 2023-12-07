from collections import Counter

card_labels = ["2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"]


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

    if len(card_counter) == 1:  # five of a kind
        return 6

    if len(card_counter) == 5:  # high card
        return 0

    if len(card_counter) == 2:
        if max(card_counter.values()) == 4:  # four of a kind
            return 5

        if (
            max(card_counter.values()) == 3 and min(card_counter.values()) == 2
        ):  # full house
            return 4

    if max(card_counter.values()) == 3:  # three of a kind
        return 3

    if (
        card_counter.most_common()[0][1] == 2 and card_counter.most_common()[1][1] == 2
    ):  # two pair
        return 2

    if (
        card_counter.most_common()[0][1] == 2 and card_counter.most_common()[1][1] == 1
    ):  # one pair
        return 1

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
