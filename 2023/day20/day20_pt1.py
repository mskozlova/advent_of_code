from collections import defaultdict, deque


class Node:
    def __init__(self, id, node_type, nodes_to):
        self.id = id
        self.node_type = node_type
        self.nodes_to = nodes_to
        self.nodes_from = None

        if node_type == "flip":
            self.active = False

    def add_nodes_from(self, nodes):
        self.nodes_from = nodes

        if self.node_type == "conj":
            self.states = dict(zip(self.nodes_from, ["low"] * len(self.nodes_from)))

    def process_signal(self, signal="low", signal_from=None):
        assert self.nodes_from is not None, f"nodes_from is None for node id {self.id}"

        if self.node_type == "flip":
            if signal == "high":
                return []
            self.active = not self.active
            output = "high" if self.active else "low"

        if self.node_type == "conj":
            assert signal_from is not None
            self.states[signal_from] = signal

            output = (
                "low"
                if all(map(lambda x: x == "high", self.states.values()))
                else "high"
            )

        if self.node_type == "broadcaster":
            output = signal

        return list(
            zip(
                self.nodes_to,
                [output] * len(self.nodes_to),
                [self.id] * len(self.nodes_to),
            )
        )


def parse_commands(input_file):
    nodes = dict()

    with open(input_file, "r") as file:
        for line in file.readlines():
            node_from, nodes_to = line.strip().split(" -> ")
            nodes_to = nodes_to.split(", ")

            if node_from == "broadcaster":
                nodes[node_from] = Node(node_from, node_from, nodes_to)
            elif node_from.startswith("%"):
                node_from = node_from[1:]
                nodes[node_from] = Node(node_from, "flip", nodes_to)
            elif node_from.startswith("&"):
                node_from = node_from[1:]
                nodes[node_from] = Node(node_from, "conj", nodes_to)
            else:
                print(f"Unknown node type: {node_from}")

    return nodes


def init_conjunctions(nodes):
    inputs = defaultdict(set)
    for node in nodes.values():
        for next_node in node.nodes_to:
            inputs[next_node].add(node.id)

    for node in nodes.values():
        node.add_nodes_from(inputs[node.id])


def run_commands(start_node, input_signal, nodes):
    low_signals, high_signals = 0, 0
    queue = deque()
    queue.append((start_node, input_signal, None))

    while len(queue) > 0:
        node_id, input_signal, signal_from = queue.popleft()

        if node_id in nodes:
            next_commands = nodes[node_id].process_signal(input_signal, signal_from)
            queue.extend(next_commands)

        if input_signal == "low":
            low_signals += 1
        else:
            high_signals += 1

    return low_signals, high_signals


nodes = parse_commands("input.txt")
init_conjunctions(nodes)

total_low_signals, total_high_signals = 0, 0

for _ in range(1000):
    low_signals, high_signals = run_commands("broadcaster", "low", nodes)
    total_low_signals += low_signals
    total_high_signals += high_signals

print(total_low_signals, total_high_signals, total_low_signals * total_high_signals)
