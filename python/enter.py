import sys
import subprocess
import requests
from os import path

if __name__ == '__main__':
    source_path = path.realpath(__file__)
    source_dir = path.dirname(source_path)

    interpreter = sys.executable
    contest_number = sys.argv[1]
    print("Entering contest:", contest_number)

    process = subprocess.Popen([
        interpreter,
        path.join(source_dir, 'web/main.py')
    ], stdout=subprocess.PIPE, stderr=subprocess.PIPE)

    new_http = requests.request("GET", f"http://127.0.0.1:5000/e/{contest_number}")
    print(new_http.text)
