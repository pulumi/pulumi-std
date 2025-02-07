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
    'FormatResult',
    'AwaitableFormatResult',
    'format',
    'format_output',
]

@pulumi.output_type
class FormatResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableFormatResult(FormatResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return FormatResult(
            result=self.result)


def format(args: Optional[Sequence[Any]] = None,
           input: Optional[str] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableFormatResult:
    """
    Formats a string according to the given format. The syntax for the format is standard sprintf syntax.
    """
    __args__ = dict()
    __args__['args'] = args
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:format', __args__, opts=opts, typ=FormatResult).value

    return AwaitableFormatResult(
        result=pulumi.get(__ret__, 'result'))
def format_output(args: Optional[pulumi.Input[Sequence[Any]]] = None,
                  input: Optional[pulumi.Input[str]] = None,
                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[FormatResult]:
    """
    Formats a string according to the given format. The syntax for the format is standard sprintf syntax.
    """
    __args__ = dict()
    __args__['args'] = args
    __args__['input'] = input
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('std:index:format', __args__, opts=opts, typ=FormatResult)
    return __ret__.apply(lambda __response__: FormatResult(
        result=pulumi.get(__response__, 'result')))
