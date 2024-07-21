# This file was auto-generated by Fern from our API Definition.

from __future__ import annotations

import typing

import pydantic
import typing_extensions

from ........core.pydantic_utilities import IS_PYDANTIC_V2, UniversalRootModel, update_forward_refs
from .deep_equality_correctness_check import DeepEqualityCorrectnessCheck
from .void_function_definition_that_takes_actual_result import VoidFunctionDefinitionThatTakesActualResult

T_Result = typing.TypeVar("T_Result")


class _Factory:
    def deep_equality(self, value: DeepEqualityCorrectnessCheck) -> AssertCorrectnessCheck:
        if IS_PYDANTIC_V2:
            return AssertCorrectnessCheck(
                root=_AssertCorrectnessCheck.DeepEquality(**value.dict(exclude_unset=True), type="deepEquality")
            )
        else:
            return AssertCorrectnessCheck(
                __root__=_AssertCorrectnessCheck.DeepEquality(**value.dict(exclude_unset=True), type="deepEquality")
            )

    def custom(self, value: VoidFunctionDefinitionThatTakesActualResult) -> AssertCorrectnessCheck:
        if IS_PYDANTIC_V2:
            return AssertCorrectnessCheck(
                root=_AssertCorrectnessCheck.Custom(**value.dict(exclude_unset=True), type="custom")
            )
        else:
            return AssertCorrectnessCheck(
                __root__=_AssertCorrectnessCheck.Custom(**value.dict(exclude_unset=True), type="custom")
            )


class AssertCorrectnessCheck(UniversalRootModel):
    factory: typing.ClassVar[_Factory] = _Factory()

    if IS_PYDANTIC_V2:
        root: typing_extensions.Annotated[
            typing.Union[_AssertCorrectnessCheck.DeepEquality, _AssertCorrectnessCheck.Custom],
            pydantic.Field(discriminator="type"),
        ]

        def get_as_union(self) -> typing.Union[_AssertCorrectnessCheck.DeepEquality, _AssertCorrectnessCheck.Custom]:
            return self.root

    else:
        __root__: typing_extensions.Annotated[
            typing.Union[_AssertCorrectnessCheck.DeepEquality, _AssertCorrectnessCheck.Custom],
            pydantic.Field(discriminator="type"),
        ]

        def get_as_union(self) -> typing.Union[_AssertCorrectnessCheck.DeepEquality, _AssertCorrectnessCheck.Custom]:
            return self.__root__

    def visit(
        self,
        deep_equality: typing.Callable[[DeepEqualityCorrectnessCheck], T_Result],
        custom: typing.Callable[[VoidFunctionDefinitionThatTakesActualResult], T_Result],
    ) -> T_Result:
        unioned_value = self.get_as_union()
        if unioned_value.type == "deepEquality":
            return deep_equality(
                DeepEqualityCorrectnessCheck(**unioned_value.dict(exclude_unset=True, exclude={"type"}))
            )
        if unioned_value.type == "custom":
            return custom(
                VoidFunctionDefinitionThatTakesActualResult(**unioned_value.dict(exclude_unset=True, exclude={"type"}))
            )


class _AssertCorrectnessCheck:
    class DeepEquality(DeepEqualityCorrectnessCheck):
        type: typing.Literal["deepEquality"] = "deepEquality"

        class Config:
            allow_population_by_field_name = True

    class Custom(VoidFunctionDefinitionThatTakesActualResult):
        type: typing.Literal["custom"] = "custom"

        class Config:
            allow_population_by_field_name = True


update_forward_refs(AssertCorrectnessCheck)
