# This file was auto-generated by Fern from our API Definition.

# flake8: noqa
# fmt: off
# isort: skip_file

from __future__ import annotations

import typing

import pydantic
import typing_extensions

from .parameter import Parameter


class VoidFunctionSignature(pydantic.BaseModel):
    parameters: typing.List[Parameter]

    class Validators:
        """
        Use this class to add validators to the Pydantic model.

            @VoidFunctionSignature.Validators.root
            def validate(values: VoidFunctionSignature.Partial) -> VoidFunctionSignature.Partial:
                ...

            @VoidFunctionSignature.Validators.field("parameters")
            def validate_parameters(v: typing.List[Parameter], values: VoidFunctionSignature.Partial) -> typing.List[Parameter]:
                ...
        """

        _validators: typing.ClassVar[
            typing.List[typing.Callable[[VoidFunctionSignature.Partial], VoidFunctionSignature.Partial]]
        ] = []
        _parameters_validators: typing.ClassVar[typing.List[VoidFunctionSignature.Validators.ParametersValidator]] = []

        @classmethod
        def root(
            cls, validator: typing.Callable[[VoidFunctionSignature.Partial], VoidFunctionSignature.Partial]
        ) -> typing.Callable[[VoidFunctionSignature.Partial], VoidFunctionSignature.Partial]:
            cls._validators.append(validator)
            return validator

        @typing.overload  # type: ignore
        @classmethod
        def field(
            cls, field_name: typing_extensions.Literal["parameters"]
        ) -> typing.Callable[
            [VoidFunctionSignature.Validators.ParametersValidator], VoidFunctionSignature.Validators.ParametersValidator
        ]:
            ...

        @classmethod
        def field(cls, field_name: str) -> typing.Any:
            def decorator(validator: typing.Any) -> typing.Any:
                if field_name == "parameters":
                    cls._parameters_validators.append(validator)
                return validator

            return decorator

        class ParametersValidator(typing_extensions.Protocol):
            def __call__(
                self, v: typing.List[Parameter], *, values: VoidFunctionSignature.Partial
            ) -> typing.List[Parameter]:
                ...

    @pydantic.root_validator
    def _validate(cls, values: typing.Dict[str, typing.Any]) -> typing.Dict[str, typing.Any]:
        for validator in VoidFunctionSignature.Validators._validators:
            values = validator(values)
        return values

    @pydantic.validator("parameters")
    def _validate_parameters(
        cls, v: typing.List[Parameter], values: VoidFunctionSignature.Partial
    ) -> typing.List[Parameter]:
        for validator in VoidFunctionSignature.Validators._parameters_validators:
            v = validator(v, values=values)
        return v

    def json(self, **kwargs: typing.Any) -> str:
        kwargs_with_defaults: typing.Any = {"by_alias": True, **kwargs}
        return super().json(**kwargs_with_defaults)

    def dict(self, **kwargs: typing.Any) -> typing.Dict[str, typing.Any]:
        kwargs_with_defaults: typing.Any = {"by_alias": True, **kwargs}
        return super().dict(**kwargs_with_defaults)

    class Partial(typing_extensions.TypedDict):
        parameters: typing_extensions.NotRequired[typing.List[Parameter]]

    class Config:
        frozen = True
