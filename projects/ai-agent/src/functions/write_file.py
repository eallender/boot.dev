import os

from utils import is_in_workdir


def write_file(working_directory, file_path, content):
    if not os.path.isabs(file_path):
        fp_abs = os.path.abspath(f"{working_directory}/{file_path}")
    else:
        fp_abs = os.path.abspath(file_path)

    if not is_in_workdir(working_directory, fp_abs):
        return f'Error: Cannot write to "{file_path}" as it is outside the permitted working directory'

    try:
        with open(fp_abs, "w") as f:
            f.write(content)
        return (
            f'Successfully wrote to "{file_path}" ({len(content)} characters written)'
        )
    except Exception as e:
        return f"Error: {e}"
