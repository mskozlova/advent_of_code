n_intervals = 0

with open("input.txt", "r") as file:
    for line in file.readlines():
        int1, int2 = line.strip().split(",")
        int1_start, int1_end = list(map(int, int1.split("-")))
        int2_start, int2_end = list(map(int, int2.split("-")))

        if (int1_start >= int2_start and int1_end <= int2_end) or (
            int1_start <= int2_start and int1_end >= int2_end
        ):
            n_intervals += 1

print(n_intervals)
