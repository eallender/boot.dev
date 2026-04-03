import os


def is_in_workdir(work_dir, path_to_check):
    work_dir_abs = os.path.abspath(work_dir)
    path_to_check_abs = os.path.abspath(path_to_check)

    return os.path.commonpath([work_dir_abs, path_to_check_abs]) == work_dir_abs
