def parse_sketch(file_name):
    with open(file_name, "r") as file:
        return list(map(lambda x: x.strip(), file.readlines()))


def find_start(sketch):
    for i, row in enumerate(sketch):
        for j, sym in enumerate(row):
            if sym == "S":
                return i, j


# ["|", "-", "L", "J", "7", "F", "S"]
tile_relationships = {
    "|": {(-1, 0), (1, 0)},
    "-": {(0, 1), (0, -1)},
    "L": {(-1, 0), (0, 1)},
    "J": {(-1, 0), (0, -1)},
    "7": {(1, 0), (0, -1)},
    "F": {(1, 0), (0, 1)},
    "S": {(-1, 0), (1, 0), (0, -1), (0, 1)},
    ".": set(),
}


def revert_relationship(relationship):
    return tuple(-r for r in relationship)


def find_next_tile(current_row, current_col, visited_tiles, sketch):
    adjacent_tiles = [(-1, 0), (1, 0), (0, -1), (0, 1)]
    adjacent_tiles = list(
        filter(
            lambda coords: current_row + coords[0] >= 0
            and current_row + coords[0] < len(sketch)
            and current_col + coords[1] >= 0
            and current_col + coords[1] < len(sketch[0]),
            adjacent_tiles,
        )
    )

    current_tile = sketch[current_row][current_col]
    next_tiles = []
    for relationship in adjacent_tiles:
        next_row, next_col = (
            current_row + relationship[0],
            current_col + relationship[1],
        )
        next_tile = sketch[next_row][next_col]

        print(f"??? {current_tile} -> {next_tile}")
        print(
            f"from: {relationship}, {relationship in tile_relationships[current_tile]}"
        )
        print(
            f"to: {revert_relationship(relationship)}, {revert_relationship(relationship) in tile_relationships[next_tile]}"
        )
        print()

        if (
            relationship in tile_relationships[current_tile]
            and revert_relationship(relationship) in tile_relationships[next_tile]
        ):
            next_tiles.append((next_row, next_col))

    assert (
        len(next_tiles) == 2
    ), f"found {len(next_tiles)} neighbour tiles for {(current_row, current_col)}: {next_tiles}"

    next_tiles = list(filter(lambda x: x not in visited_tiles, next_tiles))

    assert (len(next_tiles) == 2 and current_tile == "S") or len(next_tiles) <= 1, (
        f"Unexpected number of neighbour tiles for tile {(current_row, current_col)}: {len(next_tiles)}, {next_tiles}"
        + f"\nVisited tiles: {visited_tiles}"
    )

    if len(next_tiles) == 0: # we reached the start
        return None, None
    return next_tiles[0]


sketch = parse_sketch("input.txt")
current_row, current_col = find_start(sketch)
cycle_length = 0
visited_tiles = set()
visited_tiles.add((current_row, current_col))

while True:
    print(f"-------- STEP {cycle_length} ---------")
    print(f"current tile: {(current_row, current_col)}\n")
    current_row, current_col = find_next_tile(
        current_row, current_col, visited_tiles, sketch
    )

    if current_row is None and current_col is None:
        break

    visited_tiles.add((current_row, current_col))
    cycle_length += 1

print((cycle_length + 1) // 2)
