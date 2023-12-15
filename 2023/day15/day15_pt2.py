class Lens:
    def __init__(self, label, focal_length=None):
        self.label = label
        self.focal_length = focal_length
        self.box = self._hash(self.label)

    def _hash(self, label):
        current_value = 0
        for symbol in label:
            current_value += ord(symbol)
            current_value = (current_value * 17) % 256
        return current_value


class BoxSet:
    def __init__(self, n=256):
        self.boxes = [[] for _ in range(n)]

    def add_lens(self, lens):
        for i, box_lens in enumerate(self.boxes[lens.box]):
            if box_lens.label == lens.label:
                self.boxes[lens.box][i].focal_length = lens.focal_length
                return

        self.boxes[lens.box].append(lens)

    def remove_lens(self, lens):
        box_number, label = lens.box, lens.label

        for i, box_lens in enumerate(self.boxes[box_number]):
            if box_lens.label == label:
                self.boxes[box_number].pop(i)
                return

    def get_config(self):
        config_number = 0
        for box_number, box in enumerate(self.boxes):
            for lense_number, lense in enumerate(box):
                config_number += (
                    (1 + box_number) * (1 + lense_number) * lense.focal_length
                )
        return config_number


def parse_sequence(file_name):
    with open(file_name, "r") as file:
        sequence = file.readline().strip().split(",")

    commands = []
    for command in sequence:
        if "=" in command:
            label, focal_length = command.split("=")
            commands.append(("=", Lens(label, int(focal_length))))
        else:
            label = command[:-1]
            commands.append(("-", Lens(label)))
    return commands


sequence = parse_sequence("input.txt")
box_set = BoxSet()

for command, lens in sequence:
    if command == "=":
        box_set.add_lens(lens)
    else:
        box_set.remove_lens(lens)

print(box_set.get_config())
