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
    'CidrnetmaskResult',
    'AwaitableCidrnetmaskResult',
    'cidrnetmask',
    'cidrnetmask_output',
]

@pulumi.output_type
class CidrnetmaskResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableCidrnetmaskResult(CidrnetmaskResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return CidrnetmaskResult(
            result=self.result)


def cidrnetmask(input: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableCidrnetmaskResult:
    """
    Takes an IP address range in CIDR notation and returns the address-formatted subnet mask format
    that some systems expect for IPv4 interfaces.
    For example, cidrnetmask("10.0.0.0/8") returns 255.0.0.0.
    Not applicable to IPv6 networks since CIDR notation is the only valid notation for IPv6.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:cidrnetmask', __args__, opts=opts, typ=CidrnetmaskResult).value

    return AwaitableCidrnetmaskResult(
        result=pulumi.get(__ret__, 'result'))
def cidrnetmask_output(input: Optional[pulumi.Input[str]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[CidrnetmaskResult]:
    """
    Takes an IP address range in CIDR notation and returns the address-formatted subnet mask format
    that some systems expect for IPv4 interfaces.
    For example, cidrnetmask("10.0.0.0/8") returns 255.0.0.0.
    Not applicable to IPv6 networks since CIDR notation is the only valid notation for IPv6.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('std:index:cidrnetmask', __args__, opts=opts, typ=CidrnetmaskResult)
    return __ret__.apply(lambda __response__: CidrnetmaskResult(
        result=pulumi.get(__response__, 'result')))
