# This file was auto-generated by Fern from our API Definition.

from __future__ import annotations

import typing

import pydantic
import typing_extensions

from ....core.pydantic_utilities import IS_PYDANTIC_V2, UniversalRootModel, update_forward_refs
from .test_submission_status_v_2 import TestSubmissionStatusV2
from .workspace_submission_status_v_2 import WorkspaceSubmissionStatusV2

T_Result = typing.TypeVar("T_Result")


class _Factory:
    def test(self, value: TestSubmissionStatusV2) -> SubmissionStatusV2:
        if IS_PYDANTIC_V2:
            return SubmissionStatusV2(root=_SubmissionStatusV2.Test(**value.dict(exclude_unset=True), type="test"))
        else:
            return SubmissionStatusV2(__root__=_SubmissionStatusV2.Test(**value.dict(exclude_unset=True), type="test"))

    def workspace(self, value: WorkspaceSubmissionStatusV2) -> SubmissionStatusV2:
        if IS_PYDANTIC_V2:
            return SubmissionStatusV2(
                root=_SubmissionStatusV2.Workspace(**value.dict(exclude_unset=True), type="workspace")
            )
        else:
            return SubmissionStatusV2(
                __root__=_SubmissionStatusV2.Workspace(**value.dict(exclude_unset=True), type="workspace")
            )


class SubmissionStatusV2(UniversalRootModel):
    factory: typing.ClassVar[_Factory] = _Factory()

    if IS_PYDANTIC_V2:
        root: typing_extensions.Annotated[
            typing.Union[_SubmissionStatusV2.Test, _SubmissionStatusV2.Workspace], pydantic.Field(discriminator="type")
        ]

        def get_as_union(self) -> typing.Union[_SubmissionStatusV2.Test, _SubmissionStatusV2.Workspace]:
            return self.root

    else:
        __root__: typing_extensions.Annotated[
            typing.Union[_SubmissionStatusV2.Test, _SubmissionStatusV2.Workspace], pydantic.Field(discriminator="type")
        ]

        def get_as_union(self) -> typing.Union[_SubmissionStatusV2.Test, _SubmissionStatusV2.Workspace]:
            return self.__root__

    def visit(
        self,
        test: typing.Callable[[TestSubmissionStatusV2], T_Result],
        workspace: typing.Callable[[WorkspaceSubmissionStatusV2], T_Result],
    ) -> T_Result:
        unioned_value = self.get_as_union()
        if unioned_value.type == "test":
            return test(TestSubmissionStatusV2(**unioned_value.dict(exclude_unset=True, exclude={"type"})))
        if unioned_value.type == "workspace":
            return workspace(WorkspaceSubmissionStatusV2(**unioned_value.dict(exclude_unset=True, exclude={"type"})))


class _SubmissionStatusV2:
    class Test(TestSubmissionStatusV2):
        type: typing.Literal["test"] = "test"

        class Config:
            allow_population_by_field_name = True

    class Workspace(WorkspaceSubmissionStatusV2):
        type: typing.Literal["workspace"] = "workspace"

        class Config:
            allow_population_by_field_name = True


update_forward_refs(SubmissionStatusV2)
