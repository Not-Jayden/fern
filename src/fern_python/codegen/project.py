from __future__ import annotations

import os
import shutil
from types import TracebackType
from typing import Generator, Optional, Type

from . import AST
from .dependency_manager import DependencyManager
from .filepath import Filepath
from .module_manager import ModuleManager
from .reference_resolver_impl import ReferenceResolverImpl
from .source_file import SourceFile, SourceFileImpl


class Project:
    """
    with Project("/path/to/project") as project:
        ...
    """

    _filepath: str
    _project_name: str
    _module_manager = ModuleManager()
    _dependency_manager = DependencyManager()

    def __init__(self, filepath: str, project_name: str):
        self._filepath = os.path.join(filepath, project_name)
        self._project_name = project_name

    def source_file(self, filepath: Filepath) -> SourceFile:
        """
        with project.source_file() as source_file:
            ...
        """

        def on_finish(source_file: SourceFileImpl) -> None:
            self._module_manager.register_source_file(
                filepath=filepath,
                source_file=source_file,
            )

        module_path = convert_filepath_to_module_path(filepath)
        source_file = SourceFileImpl(
            filepath=os.path.join(
                self._get_root_module_filepath(),
                *(directory.module_name for directory in filepath.directories),
                f"{filepath.file.module_name}.py",
            ),
            module_path=module_path,
            completion_listener=on_finish,
            reference_resolver=ReferenceResolverImpl(
                project_name=self._project_name,
                module_path_of_source_file=module_path,
            ),
        )
        return source_file

    def add_dependency(self, dependency: AST.Dependency) -> None:
        self._dependency_manager.add_dependency(dependency)

    def start(self) -> None:
        if os.path.exists(self._filepath):
            shutil.rmtree(self._filepath)

    def finish(self) -> None:
        self._module_manager.write_modules(filepath=self._get_root_module_filepath())
        # TODO write dependencies to pyproject.toml

    def _get_root_module_filepath(self) -> str:
        return os.path.join(
            self._filepath,
            "src",
        )

    def __enter__(self) -> Project:
        self.start()
        return self

    def __exit__(
        self,
        exctype: Optional[Type[BaseException]],
        excinst: Optional[BaseException],
        exctb: Optional[TracebackType],
    ) -> None:
        self.finish()


def convert_filepath_to_module_path(filepath: Filepath) -> AST.ModulePath:
    return tuple(get_module_names_from_filepath(filepath))


def get_module_names_from_filepath(filepath: Filepath) -> Generator[str, None, None]:
    for directory in filepath.directories:
        yield directory.module_name
    yield filepath.file.module_name
