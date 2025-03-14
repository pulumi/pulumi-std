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
    'SignumResult',
    'AwaitableSignumResult',
    'signum',
    'signum_output',
]

@pulumi.output_type
class SignumResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, float):
            raise TypeError("Expected argument 'result' to be a float")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> float:
        return pulumi.get(self, "result")


class AwaitableSignumResult(SignumResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return SignumResult(
            result=self.result)


def signum(input: Optional[float] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableSignumResult:
    """
    Returns the greatest integer value less than or equal to the argument.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:signum', __args__, opts=opts, typ=SignumResult).value

    return AwaitableSignumResult(
        result=pulumi.get(__ret__, 'result'))
def signum_output(input: Optional[pulumi.Input[float]] = None,
                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[SignumResult]:
    """
    Returns the greatest integer value less than or equal to the argument.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('std:index:signum', __args__, opts=opts, typ=SignumResult)
    return __ret__.apply(lambda __response__: SignumResult(
        result=pulumi.get(__response__, 'result')))
