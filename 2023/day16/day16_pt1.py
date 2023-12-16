from collections import defaultdict, deque


def parse_contraption(file_name):
    with open(file_name, "r") as file:
        return [line.strip() for line in file.readlines()]


def get_next_moves(row, col, incoming_direction, contraption):
    total_rows, total_cols = len(contraption), len(contraption[0])
    current_symbol = contraption[row][col]
    next_moves = []

    if current_symbol == ".":
        if incoming_direction == "r":
            next_moves = [(row, col + 1, incoming_direction)]

        if incoming_direction == "l":
            next_moves = [(row, col - 1, incoming_direction)]

        if incoming_direction == "d":
            next_moves = [(row + 1, col, incoming_direction)]

        if incoming_direction == "u":
            next_moves = [(row - 1, col, incoming_direction)]

    if current_symbol == "-":
        if incoming_direction == "r":
            next_moves = [(row, col + 1, incoming_direction)]

        if incoming_direction == "l":
            next_moves = [(row, col - 1, incoming_direction)]

        if incoming_direction in ("u", "d"):
            next_moves = [(row, col + 1, "r"), (row, col - 1, "l")]

    if current_symbol == "|":
        if incoming_direction == "d":
            next_moves = [(row + 1, col, incoming_direction)]

        if incoming_direction == "u":
            next_moves = [(row - 1, col, incoming_direction)]

        if incoming_direction in ("r", "l"):
            next_moves = [(row + 1, col, "d"), (row - 1, col, "u")]

    if current_symbol == "/":
        if incoming_direction == "r":
            next_moves = [(row - 1, col, "u")]

        if incoming_direction == "l":
            next_moves = [(row + 1, col, "d")]

        if incoming_direction == "u":
            next_moves = [(row, col + 1, "r")]

        if incoming_direction == "d":
            next_moves = [(row, col - 1, "l")]

    if current_symbol == "\\":
        if incoming_direction == "r":
            next_moves = [(row + 1, col, "d")]

        if incoming_direction == "l":
            next_moves = [(row - 1, col, "u")]

        if incoming_direction == "u":
            next_moves = [(row, col - 1, "l")]

        if incoming_direction == "d":
            next_moves = [(row, col + 1, "r")]

    next_moves = list(
        filter(
            lambda x: x[0] >= 0
            and x[0] < total_rows
            and x[1] >= 0
            and x[1] < total_cols,
            next_moves,
        )
    )

    return next_moves


contraption = parse_contraption("input.txt")
visited_tiles = defaultdict(set)
visited_tiles[(0, 0)].add("r")  # direction of incoming light

queue = deque()
queue.append((0, 0, "r"))

while len(queue) > 0:
    current_row, current_col, current_direction = queue.popleft()
    next_moves = get_next_moves(
        current_row, current_col, current_direction, contraption
    )

    for next_row, next_col, next_direction in next_moves:
        if ((next_row, next_col) not in visited_tiles) or (
            next_direction not in visited_tiles[(next_row, next_col)]
        ):
            queue.append((next_row, next_col, next_direction))
            visited_tiles[(next_row, next_col)].add(next_direction)

print(len(visited_tiles))
