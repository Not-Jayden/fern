# This file was auto-generated by Fern from our API Definition.

from __future__ import annotations

import typing

import pydantic
import typing_extensions

from ....core.pydantic_utilities import IS_PYDANTIC_V2, UniversalBaseModel, UniversalRootModel, update_forward_refs
from .foo import Foo as resources_types_types_foo_Foo

T_Result = typing.TypeVar("T_Result")


class _Factory:
    def integer(self, value: int) -> UnionWithBaseProperties:
        if IS_PYDANTIC_V2:
            return UnionWithBaseProperties(root=_UnionWithBaseProperties.Integer(type="integer", value=value))
        else:
            return UnionWithBaseProperties(__root__=_UnionWithBaseProperties.Integer(type="integer", value=value))

    def string(self, value: str) -> UnionWithBaseProperties:
        if IS_PYDANTIC_V2:
            return UnionWithBaseProperties(root=_UnionWithBaseProperties.String(type="string", value=value))
        else:
            return UnionWithBaseProperties(__root__=_UnionWithBaseProperties.String(type="string", value=value))

    def foo(self, value: resources_types_types_foo_Foo) -> UnionWithBaseProperties:
        if IS_PYDANTIC_V2:
            return UnionWithBaseProperties(
                root=_UnionWithBaseProperties.Foo(**value.dict(exclude_unset=True), type="foo")
            )
        else:
            return UnionWithBaseProperties(
                __root__=_UnionWithBaseProperties.Foo(**value.dict(exclude_unset=True), type="foo")
            )


class UnionWithBaseProperties(UniversalRootModel):
    factory: typing.ClassVar[_Factory] = _Factory()

    if IS_PYDANTIC_V2:
        root: typing_extensions.Annotated[
            typing.Union[
                _UnionWithBaseProperties.Integer, _UnionWithBaseProperties.String, _UnionWithBaseProperties.Foo
            ],
            pydantic.Field(discriminator="type"),
        ]

        def get_as_union(
            self,
        ) -> typing.Union[
            _UnionWithBaseProperties.Integer, _UnionWithBaseProperties.String, _UnionWithBaseProperties.Foo
        ]:
            return self.root

    else:
        __root__: typing_extensions.Annotated[
            typing.Union[
                _UnionWithBaseProperties.Integer, _UnionWithBaseProperties.String, _UnionWithBaseProperties.Foo
            ],
            pydantic.Field(discriminator="type"),
        ]

        def get_as_union(
            self,
        ) -> typing.Union[
            _UnionWithBaseProperties.Integer, _UnionWithBaseProperties.String, _UnionWithBaseProperties.Foo
        ]:
            return self.__root__

    def visit(
        self,
        integer: typing.Callable[[int], T_Result],
        string: typing.Callable[[str], T_Result],
        foo: typing.Callable[[resources_types_types_foo_Foo], T_Result],
    ) -> T_Result:
        unioned_value = self.get_as_union()
        if unioned_value.type == "integer":
            return integer(unioned_value.value)
        if unioned_value.type == "string":
            return string(unioned_value.value)
        if unioned_value.type == "foo":
            return foo(resources_types_types_foo_Foo(**unioned_value.dict(exclude_unset=True, exclude={"type"})))


class _UnionWithBaseProperties:
    class Integer(UniversalBaseModel):
        type: typing.Literal["integer"] = "integer"
        value: int

    class String(UniversalBaseModel):
        type: typing.Literal["string"] = "string"
        value: str

    class Foo(resources_types_types_foo_Foo):
        type: typing.Literal["foo"] = "foo"

        class Config:
            allow_population_by_field_name = True


update_forward_refs(UnionWithBaseProperties)
