# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'JoinResult',
    'AwaitableJoinResult',
    'join',
    'join_output',
]

@pulumi.output_type
class JoinResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableJoinResult(JoinResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return JoinResult(
            result=self.result)


def join(input: Optional[Sequence[str]] = None,
         separator: Optional[str] = None,
         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableJoinResult:
    """
    Joins the list with the delimiter for a resultant string.
    """
    __args__ = dict()
    __args__['input'] = input
    __args__['separator'] = separator
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:join', __args__, opts=opts, typ=JoinResult).value

    return AwaitableJoinResult(
        result=pulumi.get(__ret__, 'result'))


@_utilities.lift_output_func(join)
def join_output(input: Optional[pulumi.Input[Sequence[str]]] = None,
                separator: Optional[pulumi.Input[str]] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[JoinResult]:
    """
    Joins the list with the delimiter for a resultant string.
    """
    ...
