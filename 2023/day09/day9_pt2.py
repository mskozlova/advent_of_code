def parse_file(file_name):
    with open(file_name, "r") as file:
        return [list(map(int, line.strip().split())) for line in file.readlines()]


def build_prediction_values(value_history):
    prediction_values = []
    pyramid_last_layer = value_history

    while not all(value == 0 for value in pyramid_last_layer):
        prediction_values.append(pyramid_last_layer[0])
        pyramid_last_layer = [
            value - prev_value
            for value, prev_value in zip(
                pyramid_last_layer[1:], pyramid_last_layer[:-1]
            )
        ]

    return prediction_values


def get_extrapolation(prediction_values):
    current_value = 0
    for value in prediction_values[::-1]:
        current_value = value - current_value
    return current_value


values_histories = parse_file("input.txt")
extrapolation_sum = 0

for value_history in values_histories:
    prediction_values = build_prediction_values(value_history)
    extrapolation_value = get_extrapolation(prediction_values)
    extrapolation_sum += extrapolation_value

print(extrapolation_sum)
