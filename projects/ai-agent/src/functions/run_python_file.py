import os
import subprocess

from utils import is_in_workdir


def run_python_file(working_directory, file_path):
    if not os.path.isabs(file_path):
        fp_abs = os.path.abspath(f"{working_directory}/{file_path}")
    else:
        fp_abs = os.path.abspath(file_path)

    if not is_in_workdir(working_directory, fp_abs):
        return f'Error: Cannot execute "{file_path}" as it is outside the permitted working directory'

    if not os.path.exists(fp_abs):
        return f'Error: File "{file_path}" not found.'

    try:
        result_str = ""
        result = subprocess.run(
            ["python3", fp_abs],
            timeout=30,
            cwd=os.path.abspath(working_directory),
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
        )
        if not result.stderr and not result.stdout:
            result_str += "No output produced."
        else:
            result_str += f"STDOUT: {result.stdout}\n"
            result_str += f"STDERR: {result.stderr}\n"
        if result.returncode != 0:
            result_str += f"Process exited with code {result.returncode}\n"
        return result_str

    except Exception as e:
        return f"Error: executing Python file: {e}"
