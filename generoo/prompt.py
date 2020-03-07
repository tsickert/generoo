from .utils import is_valid_input, equals_ignore_case

yes_no = ['y', 'n']

def handle_prompt(prompt: dict):
    prompt_type = prompt.get('type')
    if not prompt_type:
        prompt_type = 'STRING'
    if prompt_type == 'BOOL':
        prompt['options'] = yes_no
    return prompt_user(prompt)


def prompt_user(prompt: dict):
    text = prompt['text']
    default = prompt.get('default')
    options = prompt.get('options')
    validations = prompt.get('validations')
    type = prompt.get('type')
    text = format_prompt_text(text, default, options)
    input_response = input(text)
    if input_response == "":
        if default is not None:
            input_response = default
    valid = is_valid_input(input_response, validations)
    while not valid:
        text = f'{text[0:-2]} ({get_validation_strings(validations)}): '
        input_response = input(text)
        valid = is_valid_input(input_response, validations)
    if type == 'BOOL':
        input_response = yes_no_to_bool(input_response)
    return input_response

def format_prompt_text(text, default, options):
    if default is not None and not "":
        text += f" ({default})"
    if options is not None and len(options) > 0:
        text += f" [{options}]"
    text += ": "
    return text


def get_validation_strings(validations: list) -> str:
    validation_string = "The following validations must be met to continue: "
    for validation in validations:
        evaluation = validation['evaluation']
        value = validation['value']
        if evaluation:
            if equals_ignore_case(evaluation, 'REGEX'):
                text = f'Must match regular expression: {value}. '
            elif equals_ignore_case(evaluation, 'GREATER_THAN'):
                text = f'Must be greater than: {value}. '
            elif equals_ignore_case(evaluation, 'LESS_THAN'):
                text = f'Must be less than: {value}. '
            else:
                raise AttributeError(f'Invalid evaluation type for validations: {evaluation}')
            validation_string += text
    return validation_string[:-1]


def yes_no_to_bool(response: str) -> bool:
    if len(response) > 0:
        return equals_ignore_case(response[0], 'y')
    return False
