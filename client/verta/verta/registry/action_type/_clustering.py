from ..._protos.public.registry import RegistryService_pb2

from . import _ActionType


class Clustering(_ActionType):
    """
    Clustering action of the registered model.

    Examples
    --------
    .. code-block:: python

        from verta.registry import action_type
        reg_model = client.create_registered_model("My Model", workspace="my-org", action_type=action_type.Clustering())

    """

    _ACTION_TYPE = RegistryService_pb2.ActionTypeEnum.ActionType.CLUSTERING
