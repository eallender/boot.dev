import os

from utils import is_in_workdir


def get_file_content(working_directory, file_path):
    MAX_CHARS = 10000
    if not os.path.isabs(file_path):
        fp_abs = os.path.abspath(f"{working_directory}/{file_path}")
    else:
        fp_abs = os.path.abspath(file_path)

    if not is_in_workdir(working_directory, fp_abs):
        return f'Error: Cannot list "{file_path}" as it is outside the permitted working directory'

    if not os.path.isfile(fp_abs):
        return f'Error: "{file_path}" is not a directory'

    try:
        with open(fp_abs, "r") as f:
            file_content_string = f.read(MAX_CHARS)
            if len(f.read()) > MAX_CHARS:
                file_content_string += (
                    f'[...File "{file_path}" truncated at 10000 characters]'
                )

        return file_content_string
    except Exception as e:
        f"Error: unable to read file - {file_path}"
