def parse_seeds(line):
    seed_ranges = list(map(int, line[7:].strip().split()))
    return list(zip(seed_ranges[::2], seed_ranges[1::2]))


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


def intersect_segments(segment, rule_segments):
    rule_points = []
    for rule_start, rule_length in rule_segments:
        rule_points.append(rule_start)
        rule_points.append(rule_start + rule_length)

    rule_points.append(segment[0])
    rule_points.append(segment[0] + segment[1])

    rule_points = sorted(list(set(rule_points)))

    result_segments = []
    current_point = segment[0]
    segment_end = segment[0] + segment[1]
    for rule_point in rule_points:
        if rule_point > segment_end:
            break
        if rule_point > current_point and rule_point <= segment_end:
            result_segments.append((current_point, rule_point - current_point))
            current_point = rule_point

    assert segment[1] == sum(
        l for _, l in result_segments
    ), f"segment: {segment},  rule_segments: {rule_segments}, result_segments: {result_segments}"

    return result_segments


def find_destination_segments(segments, mapping):
    destination_segments = []
    for segment_start, segment_length in segments:
        is_mapping_found = False
        for destination, source, length in mapping:
            if segment_start >= source and segment_start < source + length:
                destination_segments.append(
                    [segment_start + destination - source, segment_length]
                )
                is_mapping_found = True
                break
        if not is_mapping_found:
            destination_segments.append([segment_start, segment_length])

    assert sum(l for _, l in segments) == sum(
        l for _, l in destination_segments
    ), f"segments: {segments}, mapping: {mapping}, destination_segments: {destination_segments}"
    return destination_segments


def unify_segments(segments):
    segments = sorted(segments, key=lambda x: x[0])

    unified_segments = [segments[0]]

    for segment in segments[1:]:
        if segment[0] < unified_segments[-1][0] + unified_segments[-1][1]:
            unified_segments[-1][1] = max(
                segment[0] + segment[1] - unified_segments[-1][0], segment[1]
            )
        else:
            unified_segments.append(segment)

    return unified_segments


print("Parsing the file...")
current_segments, mappings = parse_file("input.txt")
print("...parsed\n")

print(current_segments)

for i, mapping in enumerate(mappings):
    print(f"Processing mapping {i + 1} out of {len(mappings)}")
    print("current_segments:", current_segments)
    next_segments = []
    for segment_start, segment_length in current_segments:
        print("segment:", segment_start, segment_length)

        intersected_segments = intersect_segments(
            [segment_start, segment_length],
            [[source, range_length] for _, source, range_length in mapping],
        )
        print("intersected_segments:", intersected_segments)

        destination_segments = find_destination_segments(intersected_segments, mapping)
        print("destination_segments:", destination_segments)

        unified_segments = unify_segments(destination_segments)
        print("unified_segments:", unified_segments)

        next_segments.extend(unified_segments)

    current_segments = next_segments
    print("\t...done!\n")

print(current_segments)
print(min([start for start, _ in current_segments]))
