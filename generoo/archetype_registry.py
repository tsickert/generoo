import argparse
import os
import requests
import zipfile
import io

registry_types = ['git', 'zip', 'local']

def download_archetype_registry(args: argparse.Namespace):
    url = args.archetype_registry
    destination = '/tmp/generoo/registries/'
    make_temp_dir(destination)
    r = requests.get(url)
    z = zipfile.ZipFile(io.BytesIO(r.content))
    z.extractall()


def make_temp_dir(directory: str):
    os.makedirs(directory, exist_ok=True)

def fetch_registry(args: argparse.Namespace):
    url = args.archetype_registry
    type = args.registry_type

    if url is None or url == '':
        if type is not None:
            print('Registry type provided without a location. Please provide a registry location using the --archetype-registry option.')
    else:

        if type is not None or type == '':
            print('Fetching registry from github')


