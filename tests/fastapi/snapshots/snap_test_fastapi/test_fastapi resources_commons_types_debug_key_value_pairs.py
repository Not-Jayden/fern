# This file was auto-generated by Fern from our API Definition.

# flake8: noqa
# fmt: off
# isort: skip_file

from __future__ import annotations

import typing

import pydantic
import typing_extensions


class DebugKeyValuePairs(pydantic.BaseModel):
    key: DebugVariableValue
    value: DebugVariableValue

    class Validators:
        """
        Use this class to add validators to the Pydantic model.

            @DebugKeyValuePairs.Validators.root
            def validate(values: DebugKeyValuePairs.Partial) -> DebugKeyValuePairs.Partial:
                ...

            @DebugKeyValuePairs.Validators.field("key")
            def validate_key(v: DebugVariableValue, values: DebugKeyValuePairs.Partial) -> DebugVariableValue:
                ...

            @DebugKeyValuePairs.Validators.field("value")
            def validate_value(v: DebugVariableValue, values: DebugKeyValuePairs.Partial) -> DebugVariableValue:
                ...
        """

        _validators: typing.ClassVar[
            typing.List[typing.Callable[[DebugKeyValuePairs.Partial], DebugKeyValuePairs.Partial]]
        ] = []
        _key_validators: typing.ClassVar[typing.List[DebugKeyValuePairs.Validators.KeyValidator]] = []
        _value_validators: typing.ClassVar[typing.List[DebugKeyValuePairs.Validators.ValueValidator]] = []

        @classmethod
        def root(
            cls, validator: typing.Callable[[DebugKeyValuePairs.Partial], DebugKeyValuePairs.Partial]
        ) -> typing.Callable[[DebugKeyValuePairs.Partial], DebugKeyValuePairs.Partial]:
            cls._validators.append(validator)
            return validator

        @typing.overload
        @classmethod
        def field(
            cls, field_name: typing_extensions.Literal["key"]
        ) -> typing.Callable[[DebugKeyValuePairs.Validators.KeyValidator], DebugKeyValuePairs.Validators.KeyValidator]:
            ...

        @typing.overload
        @classmethod
        def field(
            cls, field_name: typing_extensions.Literal["value"]
        ) -> typing.Callable[
            [DebugKeyValuePairs.Validators.ValueValidator], DebugKeyValuePairs.Validators.ValueValidator
        ]:
            ...

        @classmethod
        def field(cls, field_name: str) -> typing.Any:
            def decorator(validator: typing.Any) -> typing.Any:
                if field_name == "key":
                    cls._key_validators.append(validator)
                if field_name == "value":
                    cls._value_validators.append(validator)
                return validator

            return decorator

        class KeyValidator(typing_extensions.Protocol):
            def __call__(self, v: DebugVariableValue, *, values: DebugKeyValuePairs.Partial) -> DebugVariableValue:
                ...

        class ValueValidator(typing_extensions.Protocol):
            def __call__(self, v: DebugVariableValue, *, values: DebugKeyValuePairs.Partial) -> DebugVariableValue:
                ...

    @pydantic.root_validator
    def _validate(cls, values: typing.Dict[str, typing.Any]) -> typing.Dict[str, typing.Any]:
        for validator in DebugKeyValuePairs.Validators._validators:
            values = validator(values)
        return values

    @pydantic.validator("key")
    def _validate_key(cls, v: DebugVariableValue, values: DebugKeyValuePairs.Partial) -> DebugVariableValue:
        for validator in DebugKeyValuePairs.Validators._key_validators:
            v = validator(v, values=values)
        return v

    @pydantic.validator("value")
    def _validate_value(cls, v: DebugVariableValue, values: DebugKeyValuePairs.Partial) -> DebugVariableValue:
        for validator in DebugKeyValuePairs.Validators._value_validators:
            v = validator(v, values=values)
        return v

    def json(self, **kwargs: typing.Any) -> str:
        kwargs_with_defaults: typing.Any = {"by_alias": True, **kwargs}
        return super().json(**kwargs_with_defaults)

    def dict(self, **kwargs: typing.Any) -> typing.Dict[str, typing.Any]:
        kwargs_with_defaults: typing.Any = {"by_alias": True, **kwargs}
        return super().dict(**kwargs_with_defaults)

    class Partial(typing_extensions.TypedDict):
        key: typing_extensions.NotRequired[DebugVariableValue]
        value: typing_extensions.NotRequired[DebugVariableValue]

    class Config:
        frozen = True


from .debug_variable_value import DebugVariableValue  # noqa: E402

DebugKeyValuePairs.update_forward_refs()
