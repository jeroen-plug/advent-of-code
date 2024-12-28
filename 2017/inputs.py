import os

def string(day: int) -> str:
    input = os.path.join("input", f"{day}.txt")
    if not os.path.exists(input):
        raise FileNotFoundError(f"Input file for Day {day} not found at {input}.")

    with open(input, "r") as f:
        content = f.read().rstrip("\n ")

    return content

def lines(day: int) -> list[str]:
    input = os.path.join("input", f"{day}.txt")
    if not os.path.exists(input):
        raise FileNotFoundError(f"Input file for Day {day} not found at {input}.")

    with open(input, "r") as f:
        content = f.readlines()

    return content
