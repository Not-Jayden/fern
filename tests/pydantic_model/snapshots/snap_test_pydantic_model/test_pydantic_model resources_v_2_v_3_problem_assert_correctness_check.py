# This file was auto-generated by Fern from our API Definition.

from __future__ import annotations

import typing

import pydantic
import typing_extensions

from ....commons.list_type import ListType
from ....commons.map_type import MapType
from ....commons.variable_type import VariableType
from .deep_equality_correctness_check import DeepEqualityCorrectnessCheck
from .void_function_definition_that_takes_actual_result import VoidFunctionDefinitionThatTakesActualResult


class AssertCorrectnessCheck_DeepEquality(DeepEqualityCorrectnessCheck):
    type: typing_extensions.Literal["deepEquality"] = "deepEquality"

    class Config:
        frozen = True


class AssertCorrectnessCheck_Custom(VoidFunctionDefinitionThatTakesActualResult):
    type: typing_extensions.Literal["custom"] = "custom"

    class Config:
        frozen = True


AssertCorrectnessCheck = typing_extensions.Annotated[
    typing.Union[AssertCorrectnessCheck_DeepEquality, AssertCorrectnessCheck_Custom],
    pydantic.Field(discriminator="type"),
]
AssertCorrectnessCheck_Custom.update_forward_refs(ListType=ListType, MapType=MapType, VariableType=VariableType)
