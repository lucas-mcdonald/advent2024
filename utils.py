def get_input_lines(file_name: str):
    with open(file_name, "r") as f:
        return list(map(lambda line: line.strip().split(), f.readlines()))