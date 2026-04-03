import os

from utils import is_in_workdir


def get_files_info(working_directory, directory: str = "."):
    if not os.path.isabs(directory):
        dir_abs = os.path.abspath(f"{working_directory}/{directory}")
    else:
        dir_abs = os.path.abspath(directory)

    if not is_in_workdir(working_directory, dir_abs):
        return f'Error: Cannot list "{directory}" as it is outside the permitted working directory'

    if not os.path.isdir(dir_abs):
        return f'Error: "{directory}" is not a directory'

    items = []
    for item in os.listdir(dir_abs):
        item_path = os.path.join(dir_abs, item)
        items.append(
            f"- {item}: file_size={os.path.getsize(item_path)} bytes, is_dir={os.path.isdir(item_path)}"
        )

    return "\n".join(items)
