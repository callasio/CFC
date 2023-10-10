import sys
import subprocess
import requests
from os import path
import threading


def start_server():
    source_path = path.realpath(__file__)
    source_dir = path.dirname(source_path)

    interpreter = sys.executable
    contest_number = sys.argv[1]

    _ = subprocess.Popen([
        interpreter,
        path.join(source_dir, 'web/main.py')
    ], stdout=subprocess.PIPE, stderr=subprocess.PIPE)

    new_http = requests.get(f"http://127.0.0.1:5000/e/{contest_number}")
    print(new_http.text)


if __name__ == '__main__':
    print("Entering contest:", sys.argv[1])
    server_thread = threading.Thread(target=start_server)
    server_thread.start()
    server_thread.join()
