EXPANTION_COEFF = 1000000


def parse_universe(file_name):
    with open(file_name, "r") as file:
        return [line.strip() for line in file.readlines()]


def generate_default_galaxy_coordinates(universe):
    coordinates = list()
    for i, row in enumerate(universe):
        for j, sym in enumerate(row):
            if sym == "#":
                coordinates.append([i, j])
    return coordinates


def adjust_coordinates(galaxies, axis=0):
    galaxies = sorted(galaxies, key=lambda x: x[axis])
    empty_counter = 0
    prev_galaxy_coordinate = 0
    for i, galaxy in enumerate(galaxies):
        print(i, galaxy, prev_galaxy_coordinate, empty_counter)
        coordinate = galaxy[axis]
        if coordinate - prev_galaxy_coordinate > 1:
            empty_counter += (EXPANTION_COEFF - 1) * (
                coordinate - prev_galaxy_coordinate - 1
            )
        galaxies[i][axis] += empty_counter
        prev_galaxy_coordinate = coordinate
    return galaxies


def get_pair_distances(galaxies):
    total_distance = 0
    for i, rhs_galaxy in enumerate(galaxies):
        for lhs_galaxy in galaxies[:i]:
            total_distance += abs(lhs_galaxy[0] - rhs_galaxy[0])
            total_distance += abs(lhs_galaxy[1] - rhs_galaxy[1])
    return total_distance


universe = parse_universe("input.txt")
galaxies = generate_default_galaxy_coordinates(universe)
galaxies = adjust_coordinates(galaxies, axis=0)
galaxies = adjust_coordinates(galaxies, axis=1)
print(get_pair_distances(galaxies))
