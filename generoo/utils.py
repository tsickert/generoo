import os
import regex
from pystache import Renderer

yes_no = ['y', 'n']


def render_template_to_directory(destination: str, template: str, parameters: dict):
    try:
        template_text = open(template, 'r').read()
        permissions = os.stat(template).st_mode
        overwrite_file(destination, renderer.render(template_text, parameters), permissions)
    except UnicodeDecodeError:
        print(f'Unable to load: {template}, may be binary file')


def render_destination_path(destination: str, parameters: dict) -> str:
    return renderer.render(destination, parameters)


def overwrite_file(file, content, permissions=None):
    directory_name = os.path.dirname(file)
    if directory_name != '':
        os.makedirs(directory_name, exist_ok=True)
    f = open(file, 'w')
    f.write(content)
    f.close()
    if permissions:
        os.chmod(file, permissions)
    

def is_valid_input(input_response: str, validations: list) -> bool:
    """
    Will compare for valid input against a user response.

    At the moment, using the greater than or less than comparator is only support for integer types.

    :param input_response:
    :param validations:
    :return:
    """
    valid = True
    if validations:
        for validation in validations:
            evaluation = validation['evaluation']
            value = validation['value']
            if evaluation:
                if equals_ignore_case(evaluation, 'REGEX'):
                    valid = regex.match(value, input_response)
                elif equals_ignore_case(evaluation, 'GREATER_THAN'):
                    valid = int(input_response) > value
                elif equals_ignore_case(evaluation, 'LESS_THAN'):
                    valid = int(input_response) < value
                elif equals_ignore_case(evaluation, 'BOOL'):
                    valid = input_response == value
                else:
                    raise AttributeError(f'Invalid evaluation type for validations: {evaluation}')
                if not valid:
                    return valid
    return valid


def equals_ignore_case(candidate: str, target: str):
    return candidate.lower() == target.lower()


renderer = Renderer()
