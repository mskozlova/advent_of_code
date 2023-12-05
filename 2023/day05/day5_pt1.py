def parse_seeds(line):
    return list(map(int, line[7:].strip().split()))


def parse_maps(lines):
    maps = []

    for line in lines:
        if line.strip() == "":
            continue

        if line.strip().endswith("map:"):
            maps.append([])
            continue

        maps[-1].append(list(map(int, line.strip().split())))

    for i, mp in enumerate(maps):
        maps[i] = sorted(mp, key=lambda x: x[1])  # source range start

    return maps


def parse_file(file_name):
    with open(file_name, "r") as file:
        seeds = parse_seeds(file.readline())
        maps = parse_maps(file.readlines())
    return seeds, maps


def get_location(seed, maps):
    current_idx = seed
    for mp in maps:
        for dest, source, range_len in mp:
            if (
                source <= current_idx and current_idx < source + range_len
            ):  # found mapping!
                idx = current_idx - source
                current_idx = dest + idx
                break
    return current_idx


seeds, maps = parse_file("input.txt")
locations = [get_location(seed, maps) for seed in seeds]
print(min(locations))
