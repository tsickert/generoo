import re


def convert_to_snake(string):
    return re.sub(r'[-./|\s]', '_', cap_sanitized(string)).lower()


def convert_to_dashes(string):
    return re.sub(r'[._/|\s]', '-', cap_sanitized(string)).lower()


def convert_to_periods(string):
    return re.sub(r'[-_/|\s]', '.', cap_sanitized(string)).lower()


def convert_to_slashes(string):
    return re.sub(r'[-_.|\s]', '/', cap_sanitized(string)).lower()


def convert_to_lower_with_spaces(string):
    return convert_to_caps_with_spaces(string).lower()


def convert_to_camel(string):
    all_caps = convert_to_caps_no_spaces(string)
    return all_caps[0].lower() + all_caps[1:]


def convert_to_caps_with_spaces(string):
    capitalized_words = extract_words(string)
    return str.join(' ', capitalized_words)


def convert_to_caps_no_spaces(string):
    capitalized_words = extract_words(string)
    return str.join('', capitalized_words)


def extract_words(string):
    string = convert_to_dashes(string)
    words = string.split('-')
    capitalized_words = []
    for word in words:
        capitalized_words.append(word.capitalize())
    return capitalized_words


def cap_sanitized(string: str):
    index = 0
    for letter in string:
        if letter.isupper():
            if index > 0:
                string = string[:index] + '|' + letter.lower()
        index += 1
    return string
