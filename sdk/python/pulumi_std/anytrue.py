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
    'AnytrueResult',
    'AwaitableAnytrueResult',
    'anytrue',
    'anytrue_output',
]

@pulumi.output_type
class AnytrueResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, bool):
            raise TypeError("Expected argument 'result' to be a bool")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> bool:
        return pulumi.get(self, "result")


class AwaitableAnytrueResult(AnytrueResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return AnytrueResult(
            result=self.result)


def anytrue(input: Optional[Sequence[Any]] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAnytrueResult:
    """
    Returns true if any of the elements in a given collection are true or \\"true\\".
    It also returns false if the collection is empty.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:anytrue', __args__, opts=opts, typ=AnytrueResult).value

    return AwaitableAnytrueResult(
        result=pulumi.get(__ret__, 'result'))
def anytrue_output(input: Optional[pulumi.Input[Sequence[Any]]] = None,
                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[AnytrueResult]:
    """
    Returns true if any of the elements in a given collection are true or \\"true\\".
    It also returns false if the collection is empty.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('std:index:anytrue', __args__, opts=opts, typ=AnytrueResult)
    return __ret__.apply(lambda __response__: AnytrueResult(
        result=pulumi.get(__response__, 'result')))
