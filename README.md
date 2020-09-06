# Generoo

<img src="https://github.com/army-of-one/generoo/blob/master/docs/generoo_icon.jpg" width="200" height="188" />

When we start new projects, we often go through a similar set of steps to bootstrap it. As a developer, we want to
spend time developing. That's where generoo comes in. Generoo allows developers to write a project template once and
then generate new projects from that template in seconds. Project templating without any additional coding gives time 
back to developers so they can focus on writing core business logic.

For an in-depth look at Generoo and it's use, see the [documentation](https://generoo.armyofone.tech).

## Installation

(WIP)

## How Does it Work?

Generoo is simple. Create a template using [Mustache](https://mustache.github.io/)'s syntax for string replacement.

### Template

A template directory or file that has files with a mustache syntax. Typically, templates will be directories
with a nested structure. Generoo looks for a `.generoo` directory in templates. This directory should contain a configuration
file called `generoo.yml`, `generoo.yaml`, or `generoo.json`.

This configuration file contains the instructions that generoo will go through to render the template correctly.

It is possible to provide a custom configuration file, see the [flags]() for more information. Custom configuration files will be added to the `.generoo`
directory in Generoo outputs. 

```json
{
  "prompts": [
    {
      "name": "who",
      "text": "Enter who you want to say hello to"
    }
  ]
}
```

The text is what the user will see when the prompt is shown and the name is the template value that will be replaced.

Running `python3 generoo.py generate project hello-world --template examples/hello-world/hello_world.py --template-config examples/hello-world/template-config.json` will prompt the user:

```
$ Say hello to:
```

When user enters: `World`, then the template is filled out and written to `hello-world/hello_world.py` and looks like:

```python
print('Hello, World')
```

For more information about how the templating system works, see the [Generoo documentation](https://generoo.armyofone.tech).

## Usage

Using generoo is simple. The CLI or python script takes 3 positional arguments:

`generoo <goal> <scope> <name> [options...]`

- `goal` - what you want generoo to do. Example: `generate`
- `scope` - what you want generoo to create. Example: `project`
- `name` - what you want to name what generoo is creating. This will be used as the root directory name. Example: `example`

Positional Arguments (in the order they appear):

### Goals

| Argument | Description | Aliases |
|---|---|---|
|`generate` | Fill in templates for an archetype or custom user project.  | `gen`, `g` |

### Scopes

| Argument | Description | Aliases |
|---|---|---|
|`project` | Generates a new project with the given name.  | `project`, `proj`, `pro`, `p` |

### Options

| Option | Description |
|---|---|
|`-n`, `--no-config` | Will run generoo without a pre-existing configuration.  |
|`-a`, `--auto-config` | Will run generoo using the pre-existing configuration and only prompt for values not present in the configuration.  |
|`-c`, `--template-config` | Points to a location on the system that contains a custom template config.  |
|`-t`, `--template` | Points to a directory on the system that contains templates for a corresponding template config.  |
|`-r`, `--run-configuration` | Points to a file on the system that contains a run configuration for a corresponding template config. |

## Built-In Templates

If no `--template` or `--template-config` arguments are given, then Generoo will generate from its built-in templates. 
Check out the `archetypes` directory to see the templates yourself. Or, better yet, try generating one. 

## Registry

Generoo is dependent on templates to work. Because the output is usually a new directory, it's impractical
to provide a path to a template and/or output directory for each run.

To make life easier, Generoo creates a registry locally that defaults to the home directory (`~/` on linux
based systems.) 

To create a better user experience, Generoo has a registry feature. 

The registry has two commands: 

`generoo register` 

`generoo link`

### Register


The `register` command will copy the target template into the local registry. Then, when
Generoo is run, it will check the local registry for the template.

If changes occur to the template in its original location, the `register` command will need
to be called again to update the Generoo registry.

### Link

The `link` command will generate a link to the template location, meaning that each time the
template is used, it will retrieve the template from its original location, so chages need to
be recommitted.  

### Usage

Once your template is registered, it can be referenced using the name and version fields in that
are provided in the Generoo configuration file (link?). 

Example:

```
generoo generate hello-world@1.0.0
```

## Contributing

Have a template that you'd like to share? Submit a PR with the template and we'll see about getting it
into the built-in templates for this project. 

Want some new functionality? Open an issue or a PR with the changes you'd like to see. 
