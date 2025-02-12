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
    'RegexallResult',
    'AwaitableRegexallResult',
    'regexall',
    'regexall_output',
]

@pulumi.output_type
class RegexallResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, list):
            raise TypeError("Expected argument 'result' to be a list")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> Sequence[Any]:
        return pulumi.get(self, "result")


class AwaitableRegexallResult(RegexallResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return RegexallResult(
            result=self.result)


def regexall(pattern: Optional[str] = None,
             string: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableRegexallResult:
    """
    Returns a list of all matches of a regular expression in a string (including named or indexed groups).
    """
    __args__ = dict()
    __args__['pattern'] = pattern
    __args__['string'] = string
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:regexall', __args__, opts=opts, typ=RegexallResult).value

    return AwaitableRegexallResult(
        result=pulumi.get(__ret__, 'result'))
def regexall_output(pattern: Optional[pulumi.Input[str]] = None,
                    string: Optional[pulumi.Input[str]] = None,
                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[RegexallResult]:
    """
    Returns a list of all matches of a regular expression in a string (including named or indexed groups).
    """
    __args__ = dict()
    __args__['pattern'] = pattern
    __args__['string'] = string
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('std:index:regexall', __args__, opts=opts, typ=RegexallResult)
    return __ret__.apply(lambda __response__: RegexallResult(
        result=pulumi.get(__response__, 'result')))
