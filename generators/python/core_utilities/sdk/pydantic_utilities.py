from functools import wraps
import typing

import pydantic
from .datetime_utils import serialize_datetime

import datetime as dt

IS_PYDANTIC_V2 = pydantic.VERSION.startswith("2.")

if IS_PYDANTIC_V2:
    from pydantic.v1.datetime_parse import (
        parse_date as parse_date,
        parse_datetime as parse_datetime
    )
    from pydantic.v1.typing import (
        get_args as get_args,
        is_union as is_union,
        get_origin as get_origin,
        is_literal_type as is_literal_type,
    )
    encoders_by_type = pydantic.deprecated.json.ENCODERS_BY_TYPE  # type: ignore

else:
    from pydantic.datetime_parse import (
        parse_date as parse_date,
        parse_datetime as parse_datetime
    )
    from pydantic.typing import (
        get_args as get_args,
        is_union as is_union,
        get_origin as get_origin,
        is_literal_type as is_literal_type,
    )
    encoders_by_type = pydantic.json.ENCODERS_BY_TYPE  # type: ignore


T = typing.TypeVar("T")


def deep_union_pydantic_dicts(
    source: typing.Dict[str, typing.Any], destination: typing.Dict[str, typing.Any]
) -> typing.Dict[str, typing.Any]:
    for key, value in source.items():
        if isinstance(value, dict):
            node = destination.setdefault(key, {})
            deep_union_pydantic_dicts(value, node)
        else:
            destination[key] = value

    return destination


def parse_obj_as(type_: typing.Type[T], obj: typing.Any) -> T:
    if IS_PYDANTIC_V2:
        adapter = pydantic.TypeAdapter(type_)
        return adapter.validate_python(obj)
    else:
        return pydantic.parse_obj_as(type_, obj)


def to_jsonable_with_fallback(obj: typing.Any, fallback_serializer: typing.Callable[[typing.Any], typing.Any]) -> typing.Any:
    if IS_PYDANTIC_V2:
        from pydantic_core import to_jsonable_python
        return to_jsonable_python(obj, fallback=fallback_serializer)
    else:
        return fallback_serializer(obj)

class UniversalBaseModel(pydantic.BaseModel):
    if IS_PYDANTIC_V2:
        # Note: `smart_union` is on by defautl in Pydantic v2
        model_config: typing.ClassVar[pydantic.ConfigDict] = pydantic.ConfigDict(
            populate_by_name=True,
            json_encoders={dt.datetime: serialize_datetime}
        )

    class Config:
        smart_union = True
        allow_population_by_field_name = True
        json_encoders = {dt.datetime: serialize_datetime}

    def json(self, **kwargs: typing.Any) -> str:
        kwargs_with_defaults: typing.Any = {"by_alias": True, "exclude_unset": True, **kwargs}
        if IS_PYDANTIC_V2:
            return super().model_dump_json(**kwargs_with_defaults)
        else:
            return super().json(**kwargs_with_defaults)

    def dict(self, **kwargs: typing.Any) -> typing.Dict[str, typing.Any]:
        kwargs_with_defaults_exclude_unset: typing.Any = {"by_alias": True, "exclude_unset": True, **kwargs}
        kwargs_with_defaults_exclude_none: typing.Any = {"by_alias": True, "exclude_none": True, **kwargs}

        if IS_PYDANTIC_V2:
            return deep_union_pydantic_dicts(
                super().model_dump(**kwargs_with_defaults_exclude_unset),
                super().model_dump(**kwargs_with_defaults_exclude_none),
            )
        else:
            return deep_union_pydantic_dicts(
                super().dict(**kwargs_with_defaults_exclude_unset), super().dict(**kwargs_with_defaults_exclude_none)
            )

def universal_root_validator(pre: bool = False):
    def decorator(func):
        @wraps(func)
        def validate(*args, **kwargs):
            if IS_PYDANTIC_V2:
                wrapped_func = pydantic.model_validator("before" if pre else "after")(func)
            else:
                wrapped_func = pydantic.root_validator(pre=pre)(func)
            
            return wrapped_func(*args, **kwargs)
        return validate
    return decorator

def universal_field_validator(field_name: str, pre: bool = False):
    def decorator(func):
        @wraps(func)
        def validate(*args, **kwargs):
            if IS_PYDANTIC_V2:
                wrapped_func = pydantic.field_validator(field_name, mode="before" if pre else "after")(func)
            else:
                wrapped_func = pydantic.validator(field_name, pre=pre)(func)
            
            return wrapped_func(*args, **kwargs)
        return validate
    return decorator
