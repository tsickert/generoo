import argparse
import os
import json
from pick import pick

from utils import handle_prompt, convert_to_hyphen_case, package_to_file, convert_to_period_case, \
    convert_to_caps_no_spaces, convert_to_caps_with_spaces, render_template_to_directory, render_destination_path, \
    is_valid_input


generate_options = ['generate', 'gen', 'g']
project_options = ['project', 'proj', 'pro', 'p']
excluded_archetypal_directories = ['common', '__pycache__']
archetype_default = 'archetypes'
template_configuration_filename = 'template-config.json'


def create_configuration_directory(args: argparse.Namespace, run_configuration: dict):
    """
    When new projects are created, Generoo will add a configuration file to the root directory of the project. This
    configuration file can be used to prepopulate date in subsequent generoo generation tasks. See documentation on
    github for more information.
    """
    print('Creating generoo configuration directory...')
    generoo_directory = f'{args.name}/.generoo'
    try:
        os.makedirs(generoo_directory)
    except FileExistsError:
        print('Generoo configuration directory already exists.')
    run_configuration_file = open(f'{generoo_directory}/run-configuration.json', 'w')
    run_configuration_file.write(json.dumps(run_configuration, indent=4, sort_keys=True))
    run_configuration_file.close()
    print('Successfully created generoo configuration directory.')


def prompt_for_archetype() -> (str, str, str):
    """
    Collects information from the user on the language, framework, and framework version they want to generate sources
    from.

    :return: language, framework, version entered by the user.
    """

    language, _ = pick(get_languages(), "Please choose a language:")
    framework, _ = pick(get_framework(language), "Please choose a framework:")
    version, _ = pick(get_versions(language, framework), "Please choose a version:")
    return language, framework, version


def get_languages() -> list:
    """Traverse templates directory to pull out all of the directories that represent languages. Ignore common.py."""
    root = 'archetypes/'
    return [subdirectory for subdirectory in os.listdir(root) if os.path.isdir(os.path.join(root, subdirectory))
            and subdirectory not in excluded_archetypal_directories]


def get_framework(language: str) -> list:
    """
    Traverse templates/{language} directory to pull out all of the directories that represent archetypes.
    Ignore common.py.
    """
    root = f'archetypes/{language}'
    return [subdirectory for subdirectory in os.listdir(root) if os.path.isdir(os.path.join(root, subdirectory))
            and subdirectory not in excluded_archetypal_directories]


def get_versions(language: str, framework: str) -> list:
    """
    Traverse templates/{language}/{framework} directory to pull out all of the directories that represent versions.
    Ignore common.py.
    """
    root = f'archetypes/{language}/{framework}'
    return [subdirectory for subdirectory in os.listdir(root) if os.path.isdir(os.path.join(root, subdirectory))
            and subdirectory not in excluded_archetypal_directories]


def get_generoo_config(args: argparse.Namespace) -> dict:
    """
    Attempts to load the run configuration from the .generoo file in the root directory. Throws exception if failed to
    find file or failed to load.
    :param args:
    :return:
    """
    configuration = open(f'{args.name}/.generoo/run-configuration.json')
    return json.loads(configuration.read())


def get_template_configuration_metadata(args: argparse.Namespace) -> (str, str):
    """
    :param args:
    :return:
    """
    config = args.template_config
    directory = args.templates

    if directory == archetype_default:
        language, framework, version = prompt_for_archetype()
        directory = f'{directory}/{language}/{framework}/{version}/'
        if not config:
            config = f'{directory}{template_configuration_filename}'
    elif directory is None:
        directory = os.path.dirname(config)
    return directory, config


def resolve_variables(template_configuration: dict) -> dict:
    """
    The first step of the lifecycle is to collect the variables provided in the configuration.

    This will add the name and value of each variable to the run-configuration.

    :param template_configuration:
    :return:
    """
    run_configuration = {}
    variables = template_configuration['variables']
    if variables:
        for variable in variables:
            if variable['name']:
                run_configuration[variable['name']] = variable['value']
            else:
                raise AttributeError
    return run_configuration


def process_follow_ups(prompt_response: str, prompt: dict, run_configuration: dict):
    follow_ups = prompt.get('follow_ups')
    if follow_ups:
        for follow_up in follow_ups:
            conditions = follow_up.get('conditions')
            if conditions:
                if is_valid_input(prompt_response, conditions):
                    process_prompt(follow_up, run_configuration)


def resolve_prompts(run_configuration: dict, template_configuration: dict) -> dict:
    """
    The second step of the lifecycle is to collect the user inputs via the prompts. The values will also be written to
    the run configuration.

    :param run_configuration:
    :param template_configuration:
    :return:
    """
    prompts = template_configuration['prompts']
    if prompts:
        for prompt in prompts:
            if prompt['name'] and prompt['text']:
                process_prompt(prompt, run_configuration)
            else:
                raise AttributeError
    return run_configuration


def process_prompt(prompt, run_configuration):
    value = handle_prompt(prompt)
    run_configuration[prompt['name']] = value
    process_follow_ups(value, prompt, run_configuration)


def resolve_transformations(run_configuration: dict, template_configurations: dict) -> dict:
    """
    The third step of the lifecycle is to make the desired transformations. The values will also be written to the run
    configuration.

    :param run_configuration:
    :param template_configurations:
    :return:
    """
    transformations = template_configurations['transformations']
    if transformations:
        for transformation in transformations:
            name = transformation['name']
            reference = transformation['reference']
            transformation_type = transformation['transformation']
            if name and reference and transformation:
                if transformation_type == 'DASHES':
                    run_configuration[name] = convert_to_hyphen_case(run_configuration[reference])
                if transformation_type == 'SLASHES':
                    run_configuration[name] = package_to_file(run_configuration[reference])
                if transformation_type == 'PERIODS':
                    run_configuration[name] = convert_to_period_case(run_configuration[reference])
                if transformation_type == 'CAPITALIZED':
                    run_configuration[name] = convert_to_caps_no_spaces(run_configuration[reference])
                if transformation_type == 'CAPITALIZED_WITH_SPACES':
                    run_configuration[name] = convert_to_caps_with_spaces(run_configuration[reference])
            else:
                raise AttributeError
    return run_configuration


def fill_in_templates(args: argparse.Namespace, template_directory: str, template_configurations: dict, run_configurations: dict):
    mappings = template_configurations.get('mappings')
    if mappings:
        for mapping in mappings:
            template = mapping['template']
            destination = mapping['destination']
            if template and destination:
                if os.path.isdir(template):
                    if os.path.isdir(destination):
                        recursively_fill_template_in_dir(args, template, destination, run_configurations)
                    else:
                        raise AttributeError(f'{template} is a directory. {destination} must be a directory.')
                else:
                    os.makedirs(template_directory, exist_ok=True)
                    render_template_to_directory(render_destination_path(destination, run_configurations), os.path.join(template_directory, template), run_configurations)
    else:
        if os.path.isdir(template_directory):
            recursively_fill_template_in_dir(args, template_directory, os.curdir, run_configurations)


def recursively_fill_template_in_dir(args: argparse.Namespace, template_dir: str, destination: str, run_configurations: dict):
    """
    Walk the directory structure non-flat template directory and render both the template content as well as the destination
    path.

    :param args:
    :param template_dir:
    :param destination:
    :param run_configurations:
    :return:
    """
    template_dir_len = len(template_dir)
    for root, dirs, files in os.walk(template_dir, topdown=False):
        for name in files:
            print(root)
            file_destination = os.path.join(args.name, destination, root[template_dir_len:], name)
            if len(file_destination) > 0:
                render_template_to_directory(render_destination_path(file_destination, run_configurations), os.path.join(root, name), run_configurations)


def extract_run_configuration(template_configuration: dict) -> dict:
    """
    Runs the lifecycle events for loading the template file. Returns a run configuration.

    :param template_configuration:
    :return:
    """
    run_configuration = resolve_variables(template_configuration)
    run_configuration = resolve_prompts(run_configuration, template_configuration)
    run_configuration = resolve_transformations(run_configuration, template_configuration)
    return run_configuration


def generate_project(args: argparse.Namespace):
    print('No pre-existing generoo run configuration found...')
    template_directory, template_file = get_template_configuration_metadata(args)
    raw_configuration = open(template_file)
    template_configuration = json.loads(raw_configuration.read())
    raw_configuration.close()
    try:
        run_configuration = get_generoo_config(args)
    except IOError:
        run_configuration = extract_run_configuration(template_configuration)
        create_configuration_directory(args, run_configuration)
    fill_in_templates(args, template_directory, template_configuration, run_configuration)


def run(args: argparse.Namespace):
    if args.goal in generate_options:
        if args.scope in project_options:
            generate_project(args)


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Generate code from templates.')

    # Positional Arguments
    parser.add_argument('goal', help='A generator goal. Examples: generate, config')
    parser.add_argument('scope', help='A generator scope. Examples: project, resource')
    parser.add_argument('name', help='The name for the scope. Example: test, pet, inventory')

    # Keyword Arguments
    parser.add_argument('-n', '--no-config',
                        help='Will run generoo without a pre-existing configuration.')
    parser.add_argument('-a', '--auto-config',
                        help='Will run generoo using the pre-existing configuration '
                             'and only prompting for values not present in the configuration.')
    parser.add_argument('-c', '--template-config',
                        help='Points to a location on the system that contains a custom template config.')
    parser.add_argument('-t', '--templates', default=archetype_default,
                        help='Points to a directory on the system that contains templates for a corresponding '
                             'template config')

    arguments = parser.parse_args()
    run(arguments)
