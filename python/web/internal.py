from os import path

source_path = path.realpath(__file__)
source_dir = path.dirname(source_path)
INTERNAL_PATH = path.join(source_dir, '../../data')


def read_login_info():
    return read_file('tmp/login-info')


def read_file(rel_path):
    with open(path.join(INTERNAL_PATH, rel_path)) as file:
        content = file.readlines()

    return list(map(str.strip, content))


if __name__ == '__main__':
    print(read_login_info())
