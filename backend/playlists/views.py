from rest_framework.generics import ListCreateAPIView, RetrieveUpdateDestroyAPIView
from django_filters.rest_framework import DjangoFilterBackend

from .serializers import PlaylistSerializer
from .filters import PlaylistFilter
from .querysets import base_queryset


class PlaylistListCreateAPIView(ListCreateAPIView):
    """Список плейлистов, создать новый плейлист"""

    queryset = base_queryset()
    serializer_class = PlaylistSerializer
    filter_backends = (DjangoFilterBackend,)
    filterset_class = PlaylistFilter
    pagination_class = None


class PlaylistRUDAPIView(RetrieveUpdateDestroyAPIView):
    """Read / Update / Delete для плейлиста"""

    queryset = base_queryset()
    serializer_class = PlaylistSerializer
