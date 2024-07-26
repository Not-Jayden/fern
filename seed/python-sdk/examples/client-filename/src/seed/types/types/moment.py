# This file was auto-generated by Fern from our API Definition.

import datetime as dt
import typing
import uuid

import pydantic

from ...core.pydantic_utilities import IS_PYDANTIC_V2, UniversalBaseModel


class Moment(UniversalBaseModel):
    """
    Examples
    --------
    import datetime
    import uuid

    from seed.types.types import Moment

    Moment(
        id=uuid.UUID(
            "656f12d6-f592-444c-a1d3-a3cfd46d5b39",
        ),
        date=datetime.date.fromisoformat(
            "1994-01-01",
        ),
        datetime=datetime.datetime.fromisoformat(
            "1994-01-01 01:01:01+00:00",
        ),
    )
    """

    id: uuid.UUID
    date: dt.date
    datetime: dt.datetime

    if IS_PYDANTIC_V2:
        model_config: typing.ClassVar[pydantic.ConfigDict] = pydantic.ConfigDict(extra="allow", frozen=True)  # type: ignore # Pydantic v2
    else:

        class Config:
            frozen = True
            smart_union = True
            extra = pydantic.Extra.allow
