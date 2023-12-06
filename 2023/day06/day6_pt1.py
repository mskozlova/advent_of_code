def parse_races(filename):
    with open(filename, "r") as file:
        lines = file.readlines()
        assert len(lines) == 2

    times = list(map(int, lines[0][5:].strip().split()))
    records = list(map(int, lines[1][9:].strip().split()))

    return list(zip(times, records))


def get_possible_distances(target_time):
    distances = []
    for hold_time in range(target_time + 1):
        distances.append((target_time - hold_time) * hold_time)  # speed = hold_time
    return distances


races = parse_races("input.txt")
total_ways_to_beat_record = 1

for target_time, record_distance in races:
    possible_distances = get_possible_distances(target_time)
    n_ways_to_beat_record = sum(
        [int(distance > record_distance) for distance in possible_distances]
    )
    total_ways_to_beat_record *= n_ways_to_beat_record

print(total_ways_to_beat_record)
