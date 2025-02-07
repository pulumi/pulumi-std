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
    'AbspathResult',
    'AwaitableAbspathResult',
    'abspath',
    'abspath_output',
]

@pulumi.output_type
class AbspathResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableAbspathResult(AbspathResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return AbspathResult(
            result=self.result)


def abspath(input: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAbspathResult:
    """
    Returns an absolute representation of the specified path.
    If the path is not absolute it will be joined with the current working directory to turn it into an absolute path.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:abspath', __args__, opts=opts, typ=AbspathResult).value

    return AwaitableAbspathResult(
        result=pulumi.get(__ret__, 'result'))
def abspath_output(input: Optional[pulumi.Input[str]] = None,
                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[AbspathResult]:
    """
    Returns an absolute representation of the specified path.
    If the path is not absolute it will be joined with the current working directory to turn it into an absolute path.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('std:index:abspath', __args__, opts=opts, typ=AbspathResult)
    return __ret__.apply(lambda __response__: AbspathResult(
        result=pulumi.get(__response__, 'result')))
