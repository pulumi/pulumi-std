# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = [
    'TimeaddResult',
    'AwaitableTimeaddResult',
    'timeadd',
    'timeadd_output',
]

@pulumi.output_type
class TimeaddResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableTimeaddResult(TimeaddResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return TimeaddResult(
            result=self.result)


def timeadd(duration: Optional[str] = None,
            timestamp: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableTimeaddResult:
    """
    Adds a duration to a timestamp, returning a new timestamp.
    	Timestamps are represented as strings using RFC 3339 "Date and time format" syntax.
    	'timestamp' must be a string adhering this syntax, i.e. "2017-11-22T00:00:00Z".
    	'duration' is a string representation of a time difference, comprised of sequences of
    	numbers and unit pairs, i.e. "3.5h" or "2h15m".
    	Accepted units are "ns", "us" or "µs", "ms", "s", "m", and "h". The first number may be negative
    	to provide a negative duration, i.e. "-2h15m".
    """
    __args__ = dict()
    __args__['duration'] = duration
    __args__['timestamp'] = timestamp
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:timeadd', __args__, opts=opts, typ=TimeaddResult).value

    return AwaitableTimeaddResult(
        result=pulumi.get(__ret__, 'result'))
def timeadd_output(duration: Optional[pulumi.Input[str]] = None,
                   timestamp: Optional[pulumi.Input[str]] = None,
                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[TimeaddResult]:
    """
    Adds a duration to a timestamp, returning a new timestamp.
    	Timestamps are represented as strings using RFC 3339 "Date and time format" syntax.
    	'timestamp' must be a string adhering this syntax, i.e. "2017-11-22T00:00:00Z".
    	'duration' is a string representation of a time difference, comprised of sequences of
    	numbers and unit pairs, i.e. "3.5h" or "2h15m".
    	Accepted units are "ns", "us" or "µs", "ms", "s", "m", and "h". The first number may be negative
    	to provide a negative duration, i.e. "-2h15m".
    """
    __args__ = dict()
    __args__['duration'] = duration
    __args__['timestamp'] = timestamp
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('std:index:timeadd', __args__, opts=opts, typ=TimeaddResult)
    return __ret__.apply(lambda __response__: TimeaddResult(
        result=pulumi.get(__response__, 'result')))
