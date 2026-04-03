def count_words(text: str) -> int:
    """Counts the words in the text

    Args:
        text (str): The string of words to be counted

    Returns:
        int: The number of words in the text string
    """
    return len(text.split())


def count_chars(text: str) -> dict:
    """Counts the number of each char in the text string

    Args:
        text (str): The text string to be analyzed

    Returns:
        dict: The dictionary of counted chars {char: frequency}
    """
    chars = {}
    for char in text:
        char = char.lower()
        if char in chars:
            chars[char] += 1
        else:
            chars[char] = 1

    return chars


def sort_chars(chars: dict) -> dict:
    """Sorts the given char dictionary in order of frequency (highest to largest)

    Args:
        chars (dict): The dict of char frequency {char: frequency}

    Returns:
        dict: The sorted chart dict
    """
    sorted_chars = {}
    for value in sorted(chars, key=chars.get, reverse=True):
        sorted_chars[value] = chars[value]

    return sorted_chars
