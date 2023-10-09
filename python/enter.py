import sys
import subprocess

if __name__ == '__main__':
    interpreter = sys.executable
    contest_number = sys.argv[1]
    print("Entering contest:", contest_number)

    process = subprocess.run([
        interpreter,
        'web/main.py'
    ])
