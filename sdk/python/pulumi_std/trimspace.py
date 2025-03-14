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
    'TrimspaceResult',
    'AwaitableTrimspaceResult',
    'trimspace',
    'trimspace_output',
]

@pulumi.output_type
class TrimspaceResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableTrimspaceResult(TrimspaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return TrimspaceResult(
            result=self.result)


def trimspace(input: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableTrimspaceResult:
    """
    Removes any space characters from the start and end of the given string,
    	following the Unicode definition of \\"space\\" (i.e. spaces, tabs, newline, etc.).
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:trimspace', __args__, opts=opts, typ=TrimspaceResult).value

    return AwaitableTrimspaceResult(
        result=pulumi.get(__ret__, 'result'))
def trimspace_output(input: Optional[pulumi.Input[str]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[TrimspaceResult]:
    """
    Removes any space characters from the start and end of the given string,
    	following the Unicode definition of \\"space\\" (i.e. spaces, tabs, newline, etc.).
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('std:index:trimspace', __args__, opts=opts, typ=TrimspaceResult)
    return __ret__.apply(lambda __response__: TrimspaceResult(
        result=pulumi.get(__response__, 'result')))
